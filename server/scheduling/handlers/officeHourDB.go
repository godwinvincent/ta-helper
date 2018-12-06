package handlers

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/alabama/final-project-alabama/server/scheduling/models"
	"gopkg.in/mgo.v2/bson"
)

// OfficeHourNotify sends out a notification to every single user
// that there is a new or a deleted office hour session
func (ctx *Context) OfficeHourNotify(updateType string) error {

	// check that it is an accepted update type
	if updateType != "office-new" && updateType != "office-deleted" {
		return fmt.Errorf("error: updateType not supported in OfficeHourNotify(): %s", updateType)
	}

	// get every user in Mongo
	var allUsers []User
	if err := ctx.UsersCollection.Collection.Find(bson.M{}).All(&allUsers); err != nil {
		return err
	}

	// make a slice of all usernames
	allUsernames := make([]string, len(allUsers))
	log.Printf("username count: %d", len(allUsernames))
	for i := 0; i < len(allUsernames)-1; i++ {
		log.Printf("index: %d", i)
		log.Printf(allUsers[i].UserName)
		allUsernames[i] = allUsers[i].UserName
	}

	msg := models.WebsocketMsg{allUsernames, updateType}
	if err := ctx.WebSocketStore.SendNotifToRabbit(&msg); err != nil {
		return fmt.Errorf("failed to notify students that OH have changed: %s", err)
	}

	return nil
}

// OfficeHourGetAllStudents gets all student usernames from an office hour session
// where each student must be currently associated to at least one live question
func (ctx *Context) OfficeHourGetAllStudents(officeHourID string) ([]string, error) {
	//map of usernames, helps us not have duplicates
	m := make(map[string]bool, 0)
	var result []string

	// TODO: ideally we'd not get each question, we'd only get the student field:
	// https://stackoverflow.com/questions/31116528/select-column-from-mongodb-in-golang-using-mgo
	// Get all questions from the office hour
	allQuestions, allErr := ctx.GetAllQuestions(officeHourID)
	if allErr != nil {
		return nil, allErr
	}

	// get student usernames from each question
	// make sure each username appears in result only once
	for _, quest := range allQuestions {
		for _, user := range quest.Students {
			m[user] = true
		}
	}

	for username := range m {
		result = append(result, username)
	}

	return result, nil
}

func (ctx *Context) OfficeHoursGetOne(officeHourID string) (models.OfficeHourSession, error) {
	var result models.OfficeHourSession

	if err := ctx.OfficeHourCollection.Collection.Find(bson.M{"_id": bson.ObjectIdHex(officeHourID)}).One(&result); err != nil {
		return result, err
	}
	return result, nil
}

func (ctx *Context) OfficeHoursInsert(oh *models.NewOfficeHourSession, username string) error {
	if err := officeHoursIsClean(oh); err != nil {
		return err
	}
	oh.TAs = append(oh.TAs, username)
	if err := ctx.OfficeHourCollection.Collection.Insert(oh); err != nil {
		return err
	}

	if err := ctx.OfficeHourNotify("office-new"); err != nil {
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

	if err := ctx.OfficeHourNotify("office-deleted"); err != nil {
		return err
	}

	return nil
}

func (ctx *Context) IncrementOfficeHours(officeHourID string, by int) error {
	if err := ctx.OfficeHourCollection.Collection.Update(bson.M{"_id": bson.ObjectIdHex(officeHourID)}, bson.M{"$inc": bson.M{"numQuestions": by}}); err != nil {
		return err
	}
	return nil
}

// TODO:
func officeHoursIsClean(oh *models.NewOfficeHourSession) error {
	return nil
}

// TODO:
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
