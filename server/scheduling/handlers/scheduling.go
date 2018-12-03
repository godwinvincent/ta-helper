package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/alabama/final-project-alabama/server/scheduling/models"
)

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
			var officeHour models.OfficeHourSession
			err := decoder.Decode(&officeHour)
			if err != nil {
				http.Error(w, "Request Body not in right format", http.StatusBadRequest)
				return
			}
			if err := ctx.OfficeHoursInsert(&officeHour, user.UserName); err != nil {
				http.Error(w, "Error inserting office hours", http.StatusInternalServerError)
				return
			}
			jsonStr, err := json.Marshal(officeHour)
			if err != nil {
				http.Error(w, "Error marshalling json response", http.StatusInternalServerError)
				return
			}
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
		var question models.Question
		err := decoder.Decode(&question)
		if err != nil {
			http.Error(w, "Request body in incorrect format", http.StatusBadRequest)
			return
		}
		// the question contains the officeHourID already
		if err := ctx.QuestionInsert(&question, user.UserName); err != nil {
			http.Error(w, "Error inserting question", http.StatusInternalServerError)
			return
		}
		jsonStr, err := json.Marshal(question)
		if err != nil {
			http.Error(w, "Error marshalling json response", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonStr)
	} else if r.Method == "PATCH" {
		// get office hour id and see if instructor of oh channel is part of this oh via db call
		if user.Role != "instructor" {
			http.Error(w, "Only an instructor of the office hour can patch the office hour", http.StatusForbidden)
			return
		}
	}
}

func (ctx *Context) SpecificQuestionHandler(w http.ResponseWriter, r *http.Request, user *User) {
	// PATCH questions
	// POST to add student to question
	// GET (?) more info
	// DEL question

}

func (ctx *Context) TAHandler(w http.ResponseWriter, r *http.Request, user *User) {
	//POST answering a question
	//PATCH ?possible editing order and duration
}

func (ctx *Context) FAQHandler(w http.ResponseWriter, r *http.Request, user *User) {

}
