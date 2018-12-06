package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alabama/final-project-alabama/server/scheduling/models"
)

// OfficeHourHandler :
//	- creates office hours
//	- gets all office hours
func (ctx *Context) OfficeHourHandler(w http.ResponseWriter, r *http.Request, user *User) {
	// create office hour, get all office hours
	// v1/officeHour
	if r.Method == "POST" {
		if user.Role != "instructor" {
			http.Error(w, "Only instructor can create office hours", http.StatusForbidden)
			return
		}
		if r.Header.Get("Content-Type") == "application/json" {
			decoder := json.NewDecoder(r.Body)
			var officeHour models.NewOfficeHourSession
			err := decoder.Decode(&officeHour)
			if err != nil {
				http.Error(w, "Request Body not in right format", http.StatusBadRequest)
				return
			}
			if err := ctx.OfficeHoursInsert(&officeHour, user.UserName); err != nil {
				http.Error(w, "Error inserting office hours: "+err.Error(), http.StatusInternalServerError)
				return
			}
			jsonStr, err := json.Marshal(officeHour)
			if err != nil {
				http.Error(w, "Error marshalling json response", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonStr)
		} else {
			http.Error(w, "Request Body must be in JSON", http.StatusUnsupportedMediaType)
			return
		}

	} else if r.Method == "GET" {
		officeHours, err := ctx.GetOfficeHours()
		if err != nil {
			http.Error(w, "Error getting office hours", http.StatusInternalServerError)
			return
		}
		jsonStr, err := json.Marshal(officeHours)
		if err != nil {
			http.Error(w, "Error marshalling json response", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonStr)

	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
}

// SpecificOfficeHourHandler :
//	- gets all questions from an office hour session
//	- make a new question in an office hour session
//	- update office hour session name
//	- remove an office hour session
func (ctx *Context) SpecificOfficeHourHandler(w http.ResponseWriter, r *http.Request, user *User) {
	// /v1/officehour/{officeHourID}

	params := r.URL.Query()
	officeHourID := params.Get("oh")
	if r.Method == "GET" {
		questions, err := ctx.GetAllQuestions(officeHourID)
		if err != nil {
			http.Error(w, "Error getting office hours", http.StatusInternalServerError)
			return
		}
		jsonStr, err := json.Marshal(questions)
		if err != nil {
			http.Error(w, "Error marshalling json response", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonStr)

	} else if r.Method == "POST" {
		if user.Role != "student" {
			http.Error(w, "Only a student can create a question", http.StatusForbidden)
			return
		}
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Request Body must be in JSON", http.StatusUnsupportedMediaType)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var question models.NewQuestion
		err := decoder.Decode(&question)
		if err != nil {
			http.Error(w, "Request body in incorrect format", http.StatusBadRequest)
			return
		}

		question.OfficeHourID = officeHourID
		if err := ctx.QuestionInsert(&question, user.UserName); err != nil {
			http.Error(w, "Error inserting question", http.StatusInternalServerError)
			return
		}

		jsonStr, err := json.Marshal(question)
		if err != nil {
			http.Error(w, "Error marshalling json response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonStr)
	} else if r.Method == "PATCH" {
		// check content type
		// if r.Header.Get("Content-Type") != "application/json" {
		// 	http.Error(w, "Request Body must be in JSON", http.StatusUnsupportedMediaType)
		// 	return
		// }
		// make sure the user is an instructor
		if user.Role != "instructor" {
			http.Error(w, "Only an instructor of the office hour can patch the office hour", http.StatusForbidden)
			return
		}
		if err := ctx.CheckOwnershipOfOfficeHours(officeHourID, user.UserName); err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var updatedOfficeHour models.UpdateOfficeHourSession
		err := decoder.Decode(&updatedOfficeHour)

		// call db to update
		if err := ctx.UpdateOfficeHours(officeHourID, &updatedOfficeHour); err != nil {
			http.Error(w, "Error inserting office hours", http.StatusInternalServerError)
			return
		}
		// marshal to json for response to user
		jsonStr, err := json.Marshal(updatedOfficeHour)
		if err != nil {
			http.Error(w, "Error marshalling json response", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonStr)

	} else if r.Method == "DELETE" {
		if user.Role != "instructor" {
			http.Error(w, "Only an instructor of the office hour can patch the office hour", http.StatusForbidden)
			return
		}
		if err := ctx.CheckOwnershipOfOfficeHours(officeHourID, user.UserName); err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		if err := ctx.RemoveOfficeHour(officeHourID); err != nil {
			http.Error(w, "Error deleting office hours", http.StatusForbidden)
			return
		}
		w.Write([]byte("Deleted Office Hours"))
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
}

// SpecificQuestionHandler :
//	- Gets a single question
//	- Patch: allows instructor to update a question
//	- Del: remove student from a question
func (ctx *Context) SpecificQuestionHandler(w http.ResponseWriter, r *http.Request, user *User) {
	// get question ID
	params := r.URL.Query()
	questionID := params.Get("qid")
	if questionID == "" {
		http.Error(w, "empty qid", http.StatusBadRequest)
		return
	}

	// begin routing
	if r.Method == "GET" {
		// Get a specific question
		q, err := ctx.QuestionGetOne(questionID)
		if err != nil {
			http.Error(w, "failed to get question", 500)
			return
		}

		jsonStr, err := json.Marshal(q)
		if err != nil {
			http.Error(w, "Error marshalling json response", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonStr)

	} else if r.Method == "POST" {
		if err := ctx.QuestionAddStudent(questionID, user.UserName); err != nil {
			http.Error(w, "Error adding student to question", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("added student to question"))

	} else if r.Method == "PATCH" {
		if user.Role != "instructor" {
			http.Error(w, "Only instructors can edit questions", http.StatusForbidden)
			return
		}
		decoder := json.NewDecoder(r.Body)
		var updates models.UpdateQuestion
		if err := decoder.Decode(&updates); err != nil {
			http.Error(w, "Error decoding json body", http.StatusBadRequest)
			return
		}
		if len(updates.Mode) == 0 || len(updates.Update) == 0 {
			http.Error(w, "JSON must have <mode> and <body> with len longer than zero", http.StatusBadRequest)
			return
		}

		// update question
		switch updates.Mode {
		case "body":
			if err := ctx.QuestionUpdate(questionID, "body", updates.Update); err != nil {
				http.Error(w, "failed to update question body: "+err.Error(), 500)
				return
			}
		case "type":
			if err := ctx.QuestionUpdate(questionID, "type", updates.Update); err != nil {
				http.Error(w, "failed to update question type: "+err.Error(), 500)
				return
			}
		case "order":
			if updates.Update == "up" {
				if err := ctx.MoveQuestionUp(questionID); err != nil {
					http.Error(w, "Error moving question up: "+err.Error(), http.StatusInternalServerError)
					return
				}
			} else if updates.Update == "down" {
				if err := ctx.MoveQuestionDown(questionID); err != nil {
					http.Error(w, "Error moving question down: "+err.Error(), http.StatusInternalServerError)
					return
				}
			}
		default:
			http.Error(w, "Update mode not supported", http.StatusBadRequest)
			return
		}
		w.Write([]byte("updated"))

	} else if r.Method == "DELETE" {
		if user.Role == "student" {
			if err := ctx.QuestionRemStudent(questionID, user.UserName); err != nil {
				http.Error(w, "Error removing student from questions", http.StatusInternalServerError)
				return
			}
			w.Write([]byte("removed student from question"))
		} else if user.Role == "instructor" {
			if err := ctx.QuestionDelete(questionID, "instructor"); err != nil {
				http.Error(w, "error deleting question as instructor", http.StatusInternalServerError)
				return
			}
		}

	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	// PATCH questions
	// GET more info

}

// WebSocketConnectionHandler :
//	- Post: if instructor: notifies all students in a question that it is their turn
//		Example: /v1/question/answer?qid=<QuestionID>
func (ctx *Context) WebSocketConnectionHandler(w http.ResponseWriter, r *http.Request, user *User) {
	if r.Method == "POST" {

		// get question ID
		params := r.URL.Query()
		questionID := params.Get("qid")
		if questionID == "" {
			http.Error(w, "empty qid", http.StatusBadRequest)
			return
		}

		// check that they're an instructor
		if user.Role != "instructor" {
			http.Error(w, "only instructors/TAs may notify students", http.StatusForbidden)
			return
		}

		// get all usernames for a question
		usernames, getErr := ctx.GetStudentsQuestion(questionID)
		if getErr != nil {
			http.Error(w, "error on notify students it's their questions turn: failed to get all usernames", 500)
			return
		}
		// send to rabbit
		msg := models.WebsocketMsg{usernames, "question-yourTurn"}
		if err := ctx.WebSocketStore.SendNotifToRabbit(&msg); err != nil {
			http.Error(w, "failed to notify students it's their questions turn", 500)
			return
		}

		// respond
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("students notified"))

	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

}

// ---------------- Stretch Functions ----------------
// TODO:
//	-
//	-
func (ctx *Context) FAQHandler(w http.ResponseWriter, r *http.Request, user *User) {

}

// ---------------- Helper Functions ----------------

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
