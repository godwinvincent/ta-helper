package questions

/**
 * This file handles the communication between
 * API gateway and the Database for all questions
 * in the TA Queue.
 *
 */

import (
	"fmt"
)

// ------------- Question Functions -------------

// question length may not be more than 700 characters (that's plenty long)

// Insert a question into the DB.
// Must pass is username of the person who created the question.
func (c *QuestionCollection) Insert(q *Question, username string) error {
	// make sure question is clean
	if err := questIsClean(q); err != nil {
		return err
	}
	// add question creator
	q.Students = append(q.Students, username)
	// insert into DB

	return nil
}

// Add a student to question

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
	// position
	return nil
}
