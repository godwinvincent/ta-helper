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
			http.Error(w, "Only instructor can create office hours", http.StatusUnauthorized)
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
			if err := ctx.OfficeHourCollection.Insert(&officeHour, user.UserName); err != nil {
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
		officeHours := ctx.OfficeHourCollection.Get()
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

func (ctx *questions.Context) SpecificOfficeHourHandler(w http.ResponseWriter, r *http.Request, user *User) {
	// /v1/officehour/{officeHourID}
	// GET all questions for office hours
	// POST new question
	// PATH Office hour name
	// DELETE Office hours
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
