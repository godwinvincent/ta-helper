package questions

/**
 * This file handles the communication between
 * API gateway and the Database for all questions
 * in the TA Queue.
 *
 */

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

// ------------- Question Functions -------------

// question length may not be more than 700 characters (that's plenty long)

//Get all question in office hour id
func (ctx *Context) GetAllQuestions(officeHourID string) ([]Question, error) {
	var results []Question
	if err := ctx.QuestionCollection.collection.Find(bson.M{"_id": bson.ObjectIdHex(officeHourID)}).All(&results); err != nil {
		return nil, err
	}
	return results, nil
}

// Insert a question into the DB.
// Must pass is username of the person who created the question.
func (ctx *Context) QuestionInsert(q *Question, creatorUsername string) error {

	qColl := ctx.QuestionCollection
	oColl := ctx.OfficeHourCollection

	// make sure question is clean
	if err := questIsClean(q); err != nil {
		return err
	}

	// add question creator
	q.Students = append(q.Students, creatorUsername)

	// find how many questions are already in the Office Hour Session
	office := OfficeHourSession{}
	if err := oColl.collection.Find(bson.M{"_id": bson.ObjectIdHex(q.OfficeHourID)}).One(&office); err != nil {
		return err
	}

	// modify the position of the question
	q.QuestionPosition = office.NumQuestions + 1

	// insert into DB
	if err := qColl.collection.Insert(q); err != nil {
		return err
	}
	return nil
}

func (c *QuestionCollection) GetAll(officeHourID string) error {
	// db call to get all questions in given office hour

	return nil
}

// Add a student to question
func (ctx *Context) QuestionAddStudent(q *Question, studentUsername string) error {
	return nil
}

// Remove Student from question
// if question has no students delete questions

//

// Change a questions' order
// Only a TA can do that

// ------------- Helper Functions -------------

func questIsClean(q *Question) error {
	// message body may not be too long
	if len(q.QuestionBody) > MaxQuestLength {
		return fmt.Errorf("question may not be longer than %d, it currently is %d", MaxQuestLength, len(q.QuestionBody))
	}
	// make sure that the question is part of an Office Hour session
	if len(q.OfficeHourID) == 0 {
		return fmt.Errorf("this question must be associated to an office hour, office hour id is of length: %d", len(q.OfficeHourID))
	}
	// make sure that the question is part of an Office Hour session
	if len(q.QuestionType) == 0 {
		return fmt.Errorf("this question must have a question type")
	}
	// position

	return nil
}
