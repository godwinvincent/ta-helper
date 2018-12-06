package handlers

/**
 * This file handles the communication between
 * API gateway and the Database for all questions
 * in the TA Queue.
 *
 */

import (
	"encoding/hex"
	"errors"
	"fmt"
	"log"

	"github.com/alabama/final-project-alabama/server/scheduling/models"
	"gopkg.in/mgo.v2/bson"
)

// ------------- Question Functions -------------

// GetAllQuestions gets all questions from an office hour session.
// Takes in office hour id.
// func (ctx *Context) GetAllQuestions(officeHourID string) ([]models.Question, error) {
// 	var results []models.Question
// 	if err := ctx.QuestionCollection.Collection.Find(bson.M{"offHourID": officeHourID}).All(&results); err != nil {
// 		return nil, err
// 	}
// 	return results, nil
// }

// --------------- Get Questions ---------------

//QuestionGetOne get a specific question
func (ctx *Context) QuestionGetOne(questionID string) (models.Question, error) {
	var result models.Question

	if err := ctx.QuestionCollection.Collection.Find(bson.M{"_id": bson.ObjectIdHex(questionID)}).One(&result); err != nil {
		return result, err
	}
	return result, nil
}

// GetAll returns all questions from a specific Offie Hour Session
func (ctx *Context) GetAllQuestions(officeHourID string) ([]models.Question, error) {
	// db call to get all questions in given office hour
	var results []models.Question
	if err := ctx.QuestionCollection.Collection.Find(bson.M{"offHourID": officeHourID}).All(&results); err != nil {
		return nil, err
	}
	// convert each ID to a readable format
	for _, oh := range results {
		decodedID, err := hex.DecodeString(oh.ID.Hex())
		if err != nil {
			return nil, err
		}
		oh.ID = bson.ObjectId(decodedID)
	}
	return results, nil
}

// GetStudentsQuestion gets the students from a single question
func (ctx *Context) GetStudentsQuestion(questionID string) ([]string, error) {
	question, err := ctx.QuestionGetOne(questionID)
	if err != nil {
		return nil, err
	}

	return question.Students, nil

}

// --------------- Modify Questions ---------------

// QuestionNotify notifies all the students in an office hour session
// that one of the questions was either deleted or updated.
func (ctx *Context) QuestionNotify(officeHourID string, updateType string) error {
	log.Println("notifying for question")
	if updateType != "question-new" && updateType != "question-deleted" && updateType != "question-modified" {
		return fmt.Errorf("error: updateType not supported in QuestionNotify(): %s", updateType)
	}
	usernames, err := ctx.OfficeHourGetAllStudents(officeHourID)
	if err != nil {
		return err
	}
	tas, err := ctx.OfficeHourGetAllTAs(officeHourID)
	if err != nil {
		return err
	}
	usernames = append(usernames, tas...)
	log.Println(usernames)
	msg := models.WebsocketMsg{usernames, updateType}
	if err := ctx.WebSocketStore.SendNotifToRabbit(&msg); err != nil {
		return fmt.Errorf("failed to notify students it's their questions turn %s", err)
	}

	return nil

}

// QuestionInsert inserts a question into the DB.
// Must pass in username of the person who created the question.
func (ctx *Context) QuestionInsert(q *models.NewQuestion, creatorUsername string) error {

	qColl := ctx.QuestionCollection
	oColl := ctx.OfficeHourCollection

	// make sure question is clean
	if err := questIsClean(q); err != nil {
		return err
	}

	// add question creator
	q.Students = append(q.Students, creatorUsername)

	// find how many questions are already in the Office Hour Session
	office := models.OfficeHourSession{}
	if err := oColl.Collection.Find(bson.M{"_id": bson.ObjectIdHex(q.OfficeHourID)}).One(&office); err != nil {
		return err
	}

	// modify the position of the question
	q.QuestionPosition = office.NumQuestions + 1

	// insert into DB
	if err := qColl.Collection.Insert(q); err != nil {
		return err
	}

	if err := ctx.IncrementOfficeHours(q.OfficeHourID, 1); err != nil {
		return err
	}

	if err := ctx.QuestionNotify(q.OfficeHourID, "question-new"); err != nil {
		return err
	}
	return nil
}

// QuestionDelete deletes a question.
// Requirements: question must either have no students in it,
// or the user must be an instructor
func (ctx *Context) QuestionDelete(questionID string, userRole string) error {
	q, err := ctx.QuestionGetOne(questionID)
	if err != nil {
		return err
	}
	if userRole == "instructor" || len(q.Students) == 0 {
		// delete question
		if err := ctx.MoveQuestionsUpBelow(questionID); err != nil {
			return err
		}
		if err := ctx.QuestionCollection.Collection.Remove(bson.M{"_id": bson.ObjectIdHex(questionID)}); err != nil {
			return err
		}
		if err := ctx.IncrementOfficeHours(q.OfficeHourID, -1); err != nil {
			return err
		}

	} else {
		return fmt.Errorf("can't delete due to role or students in question")
	}

	if err := ctx.QuestionNotify(q.OfficeHourID, "question-deleted"); err != nil {
		return err
	}

	return nil
}

// QuestionUpdate changes a field in a question.
// Accepted fields:
//	- body
//	- type
func (ctx *Context) QuestionUpdate(questionID string, field string, newFieldData string) error {
	switch field {
	case "body":
		if err := ctx.QuestionCollection.Collection.Update(bson.M{"_id": questionID}, bson.M{"$set": bson.M{"questBody": newFieldData}}); err != nil {
			return err
		}
	case "type":
		if err := ctx.QuestionCollection.Collection.Update(bson.M{"_id": questionID}, bson.M{"$set": bson.M{"questType": newFieldData}}); err != nil {
			return err
		}
	default:
		return errors.New("you may only update the <body> and type <fields>")
	}
	return nil
}

// --------------- Add & Remove Students ---------------

// QuestionAddStudent adds a student to a question
// Takes in a the ID of the question and the username
// of the student who should be added.
func (ctx *Context) QuestionAddStudent(questionID string, studentUsername string) error {

	err2 := ctx.QuestionCollection.Collection.Update(bson.M{"_id": bson.ObjectIdHex(questionID)}, bson.M{"$addToSet": bson.M{"students": studentUsername}})

	if err2 != nil {
		return err2
	}
	q, err := ctx.QuestionGetOne(questionID)
	if err != nil {
		return err
	}
	if err := ctx.QuestionNotify(q.OfficeHourID, "question-modified"); err != nil {
		return err
	}
	return nil
}

// QuestionRemStudent removes a student from a question.
// Takes in a the ID of the question and the username
// of the student that should be removed.
func (ctx *Context) QuestionRemStudent(questionID string, studentUsername string) error {
	if err := ctx.QuestionCollection.Collection.Update(bson.M{"_id": bson.ObjectIdHex(questionID)}, bson.M{"$pull": bson.M{"students": studentUsername}}); err != nil {
		return err
	}

	// call delete on the question: it checks if no students are in it.
	// if there are non then it deletes the question
	ctx.QuestionDelete(questionID, "student")
	q, err := ctx.QuestionGetOne(questionID)
	if err != nil {
		return err
	}
	if err := ctx.QuestionNotify(q.OfficeHourID, "question-modified"); err != nil {
		return err
	}

	return nil
}

// --------------- Move Question ---------------
func (ctx *Context) MoveQuestionsUpBelow(questionID string) error {
	q, err := ctx.QuestionGetOne(questionID)
	if err != nil {
		return err
	}
	if _, err := ctx.QuestionCollection.Collection.UpdateAll(bson.M{"offHourID": q.OfficeHourID, "questPos": bson.M{"$gt": q.QuestionPosition}}, bson.M{"$inc": bson.M{"questPos": -1}}); err != nil {
		return err
	}
	return nil

}

func (ctx *Context) MoveQuestionUp(questionID string) error {
	q, err := ctx.QuestionGetOne(questionID)
	if q.QuestionPosition == 1 {
		return fmt.Errorf("cannot move first question up")
	}
	if err != nil {
		log.Println("failed in get")
		return err
	}
	if err := ctx.QuestionCollection.Collection.Update(bson.M{"offHourID": q.OfficeHourID, "questPos": q.QuestionPosition - 1}, bson.M{"$inc": bson.M{"questPos": 1}}); err != nil {
		log.Println("failed in update 1")
		return err
	}
	// if err := ctx.QuestionCollection.Collection.Update(bson.M{"$and": []bson.M{bson.M{"officeHourID": q.OfficeHourID}, bson.M{"questPos": q.QuestionPosition - 1}}}, bson.M{"$inc": bson.M{"questPos": 1}}); err != nil {
	// 	log.Println("failed in update 1")
	// 	return err
	// }
	if err := ctx.QuestionCollection.Collection.Update(bson.M{"_id": bson.ObjectIdHex(questionID)}, bson.M{"$inc": bson.M{"questPos": -1}}); err != nil {
		log.Println("failed in update 2")
		return err
	}
	if err := ctx.QuestionNotify(q.OfficeHourID, "question-modified"); err != nil {
		return err
	}
	return nil
}

func (ctx *Context) MoveQuestionDown(questionID string) error {
	q, err := ctx.QuestionGetOne(questionID)
	oh, err := ctx.OfficeHoursGetOne(q.OfficeHourID)
	if q.QuestionPosition == oh.NumQuestions {
		return fmt.Errorf("cannot move last question down")
	}
	if err != nil {
		return err
	}
	if err := ctx.QuestionCollection.Collection.Update(bson.M{"offHourID": q.OfficeHourID, "questPos": q.QuestionPosition + 1}, bson.M{"$inc": bson.M{"questPos": -1}}); err != nil {
		return err
	}
	if err := ctx.QuestionCollection.Collection.Update(bson.M{"_id": bson.ObjectIdHex(questionID)}, bson.M{"$inc": bson.M{"questPos": 1}}); err != nil {
		return err
	}
	if err := ctx.QuestionNotify(q.OfficeHourID, "question-modified"); err != nil {
		return err
	}
	return nil
}

// ------------- Helper Functions -------------

func questIsClean(q *models.NewQuestion) error {
	// message body may not be too long
	if len(q.QuestionBody) > models.MaxQuestLength {
		return fmt.Errorf("question may not be longer than %d, it currently is %d", models.MaxQuestLength, len(q.QuestionBody))
	}
	// make sure that the question is part of an Office Hour session
	if len(q.OfficeHourID) == 0 {
		return fmt.Errorf("this question must be associated to an office hour, office hour id is of length: %d", len(q.OfficeHourID))
	}
	// make sure that the question is part of an Office Hour session
	if len(q.QuestionType) == 0 {
		return fmt.Errorf("this question must have a question type")
	}

	return nil
}
