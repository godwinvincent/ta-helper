package handlers

import (
	"github.com/alabama/final-project-alabama/server/scheduling/models"
	"gopkg.in/mgo.v2/bson"
)

func (ctx *Context) OfficeHoursInsert(oh *models.OfficeHourSession, username string) error {
	if err := officeHoursIsClean(oh); err != nil {
		return err
	}
	oh.TAs = append(oh.TAs, username)
	ctx.OfficeHourCollection.Collection.Insert(oh)
	return nil
}

func (ctx *Context) GetOfficeHours() ([]models.OfficeHourSession, error) {
	var results []models.OfficeHourSession
	if err := ctx.OfficeHourCollection.Collection.Find(bson.M{}).All(&results); err != nil {
		return nil, err
	}
	return results, nil
}

func officeHoursIsClean(oh *models.OfficeHourSession) error {
	return nil
}
