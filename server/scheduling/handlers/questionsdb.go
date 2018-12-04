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

//QuestionGetOne get a specific question
func (ctx *Context) QuestionGetOne(questionID string) (models.Question, error) {
	var result models.Question

	if err := ctx.QuestionCollection.Collection.Find(bson.M{"_id": bson.ObjectIdHex(questionID)}).One(&result); err != nil {
		return result, err
	}
	return result, nil
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

	if err := oColl.Collection.Update(bson.M{"_id": bson.ObjectIdHex(q.OfficeHourID)}, bson.M{"$set": bson.M{"numQuestions": q.QuestionPosition}}); err != nil {
		return err
	}
	return nil
}

// QuestionAddStudent adds a student to a question
// Takes in a the ID of the question and the username
// of the student who should be added.
func (ctx *Context) QuestionAddStudent(questionID string, studentUsername string) error {

	err2 := ctx.QuestionCollection.Collection.Update(bson.M{"_id": bson.ObjectIdHex(questionID)}, bson.M{"$addToSet": bson.M{"students": studentUsername}})

	if err2 != nil {
		return err2
	}
	return nil
}

// QuestionRemStudent removes a student from a question.
// Takes in a the ID of the question and the username
// of the student that should be removed.
func (ctx *Context) QuestionRemStudent(questionID string, studentUsername string) error {
	err2 := ctx.QuestionCollection.Collection.Update(bson.M{"_id": bson.ObjectIdHex(questionID)}, bson.M{"$pull": bson.M{"students": studentUsername}})

	if err2 != nil {
		return err2
	}

	// call delete on the question: it checks if no students are in it.
	// if there are non then it deletes the question
	ctx.QuestionRemStudent(questionID, "student")

	return nil
}

// GetAll returns all questions from a specific Offie Hour Session
func (ctx *Context) GetAll(officeHourID string) ([]models.Question, error) {
	// db call to get all questions in given office hour
	var results []models.Question
	if err := ctx.QuestionCollection.Collection.Find(bson.M{}).All(&results); err != nil {
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

// QuestionDelete deletes a question.
// Requirements: question must either have no students in it,
// or the user must be an instructor
func (ctx *Context) QuestionDelete(questionID string, userRole string) error {
	if userRole == "instructor" {
		// delete question
		if err := ctx.QuestionCollection.Collection.Remove(bson.M{"_id": bson.ObjectIdHex(questionID)}); err != nil {
			return err
		}

	} else {
		// get the question and check how many students are in it
		q, err := ctx.QuestionGetOne(questionID)
		if err != nil {
			return err
		}
		if len(q.Students) == 0 {
			if err := ctx.QuestionCollection.Collection.Remove(bson.M{"_id": bson.ObjectIdHex(questionID)}); err != nil {
				return err
			}
		} else {
			return errors.New("question still has students associated to it")
		}

	}
	return nil
}

// ------------------- Not DONE ---------------- // FIXME:

// Change a questions' order
// Only an instructor can do that

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
