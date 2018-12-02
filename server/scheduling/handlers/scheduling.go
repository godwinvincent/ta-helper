package handlers

import (
	"net/http"
)

func (ctx *Context) OfficeHourHandler(w http.ResponseWriter, r *http.Request, user *User) {
	// create office hour, get all office hours
	// v1/officeHour
}

func (ctx *Context) SpecificOfficeHourHandler(w http.ResponseWriter, r *http.Request, user *User) {
	// /v1/officehour/{officeHourID}
	// GET all questions for office hours
	// POST new question
	// PATH Office hour name
	// DELETE Office hours
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
