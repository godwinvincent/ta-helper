package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/alabama/final-project-alabama/server/scheduling/questions"
)

func (ctx *questions.Context) OfficeHourHandler(w http.ResponseWriter, r *http.Request, user *User) {
	// create office hour, get all office hours
	// v1/officeHour

	if r.Method == "POST" {
		if user.Role != "instructor" {
			http.Error(w, "Only instructor can create office hours", http.StatusForbidden)
			return
		}
		if r.Header.Get("Content-Type") == "application/json" {
			decoder := json.NewDecoder(r.Body)
			var officeHour questions.OfficeHourSession
			err := decoder.Decode(&officeHour)
			if err != nil {
				http.Error(w, "Request Body not in right format", http.StatusBadRequest)
				return
			}
			if err := ctx.OfficeHourCollection.InsertOfficeHour(&officeHour, user.UserName); err != nil {
				http.Error(w, "Error inserting office hours", http.StatusInternalServerError)
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
		officeHours := ctx.OfficeHourCollection.Get()
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

func (ctx *Context) SpecificOfficeHourHandler(w http.ResponseWriter, r *http.Request, user *User) {
	// /v1/officehour/{officeHourID}
	params := r.URL.Query()
	officeHourID := params.Get("officeHourID")
	if r.Method == "GET" {
		// no need to check if user is authorized, all users are able to view all OH and questions.
		if err := ctx.QuestionCollection.GetAll(officeHourID); err != nil {
			http.Error(w, "Error getting all questions", http.StatusInternalServerError)
			return
		}
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
		var question questions.Question
		err := decoder.Decode(&question)
		if err != nil {
			http.Error(w, "Request body in incorrect format", http.StatusBadRequest)
			return
		}
		// the question contains the officeHourID already
		if err := ctx.Insert(&question, user.UserName); err != nil {
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
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Request Body must be in JSON", http.StatusUnsupportedMediaType)
			return
		}
		// make sure the user is an instructor
		if user.Role != "instructor" {
			http.Error(w, "Only an instructor of the office hour can patch the office hour", http.StatusForbidden)
			return
		}
		// db call to get officeHourID and see if instructor of oh
		// channel is part of this oh via db call. If it returns something,
		// assume the instructor is part of the channel.
		validInstructor := ctx.IsValidInstructor(officeHourID, user.UserName)
		if len(validInstructor) == 0 {
			http.Error(w, "Only an instructor of the office hour can patch the office hour", http.StatusForbidden)
			return
		}
		// decode into updates struct when ready
		/*
			decoder := json.NewDecoder(r.Body)
			var updatedQuestion questions.Updates
			err := decoder.Decode(&updatedQuestion)

			// call db to update
			if err := ctx.UpdateOfficeHour(&officeHour, user.UserName); err != nil {
				http.Error(w, "Error inserting office hours", http.StatusInternalServerError)
				return
			}
			// marshal to json for response to user
			jsonStr, err := json.Marshal(officeHour)
			if err != nil {
				http.Error(w, "Error marshalling json response", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonStr)
		*/
	} else if r.Method == "DELETE" {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Request Body must be in JSON", http.StatusUnsupportedMediaType)
			return
		}
		/*
			DB call to check if user.UserName is creator of office hour id
			assuming we get something back if user is the creator of the office hour

			if len(whatwegetback) == 0 {
				http.Error(w, "Only the creator can delete the office hour", http.StatusForbidden)
				return
			}

			DB call to make delete
			if err := ctx.Insert(&question, user.UserName); err != nil {
				http.Error(w, "Error inserting question", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			w.Write("Office Hour Channel Deleted")
		*/
	}
}

func (ctx *questions.Context) SpecificQuestionHandler(w http.ResponseWriter, r *http.Request, user *User) {
	// PATCH questions
	// POST to add student to question
	// GET (?) more info
	// DEL question

}

func (ctx *questions.Context) TAHandler(w http.ResponseWriter, r *http.Request, user *User) {
	//POST answering a question
	//PATCH ?possible editing order and duration
}

func (ctx *questions.Context) FAQHandler(w http.ResponseWriter, r *http.Request, user *User) {

}
