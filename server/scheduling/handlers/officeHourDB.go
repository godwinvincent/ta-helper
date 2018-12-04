package handlers

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/alabama/final-project-alabama/server/scheduling/models"
	"gopkg.in/mgo.v2/bson"
)

func (ctx *Context) OfficeHoursInsert(oh *models.NewOfficeHourSession, username string) error {
	if err := officeHoursIsClean(oh); err != nil {
		return err
	}
	oh.TAs = append(oh.TAs, username)
	if err := ctx.OfficeHourCollection.Collection.Insert(oh); err != nil {
		return err
	}
	return nil
}

func (ctx *Context) GetOfficeHours() ([]models.OfficeHourSession, error) {
	var results []models.OfficeHourSession
	if err := ctx.OfficeHourCollection.Collection.Find(bson.M{}).All(&results); err != nil {
		return nil, err
	}
	for _, oh := range results {
		log.Println(oh.ID.Hex())
		decodedID, err := hex.DecodeString(oh.ID.Hex())
		if err != nil {
			log.Println(err)
			return nil, err
		}
		oh.ID = bson.ObjectId(decodedID)
	}
	return results, nil
}

func (ctx *Context) UpdateOfficeHours(officeHourID string, updatedOfficeHour *models.UpdateOfficeHourSession) error {
	if err := cleanUpdate(updatedOfficeHour); err != nil {
		return err
	}
	if err := ctx.OfficeHourCollection.Collection.Update(bson.M{"_id": bson.ObjectIdHex(officeHourID)}, bson.M{"$set": bson.M{"name": updatedOfficeHour.Name}}); err != nil {
		return err
	}
	return nil
}

func (ctx *Context) CheckOwnershipOfOfficeHours(officeHourID string, username string) error {
	var result models.OfficeHourSession
	if err := ctx.OfficeHourCollection.Collection.Find(bson.M{}).One(&result); err != nil {
		return err
	}
	if !sliceContains(result.TAs, username) {
		return fmt.Errorf("You do not own this office hours")
	}
	return nil

}

func (ctx *Context) RemoveOfficeHour(officeHourID string) error {
	if err := ctx.OfficeHourCollection.Collection.Remove(bson.M{"_id": bson.ObjectIdHex(officeHourID)}); err != nil {
		return err
	}
}

func officeHoursIsClean(oh *models.NewOfficeHourSession) error {
	return nil
}

func cleanUpdate(oh *models.UpdateOfficeHourSession) error {
	return nil
}

func sliceContains(sl []string, v string) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}
