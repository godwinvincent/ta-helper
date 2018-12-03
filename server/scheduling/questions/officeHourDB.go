package questions

import (
	"gopkg.in/mgo.v2/bson"
)

func (ctx *Context) OfficeHoursInsert(oh *OfficeHourSession, username string) error {
	if err := officeHoursIsClean(oh); err != nil {
		return err
	}
	oh.TAs = append(oh.TAs, username)
	ctx.OfficeHourCollection.collection.Insert(oh)
	return nil
}

func (ctx *Context) GetOfficeHours() ([]OfficeHourSession, error) {
	var results []OfficeHourSession
	if err := ctx.OfficeHourCollection.collection.Find(bson.M{}).All(&results); err != nil {
		return nil, err
	}
	return results, nil
}

func officeHoursIsClean(oh *OfficeHourSession) error {
	return nil
}
