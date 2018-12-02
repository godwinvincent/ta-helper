package questions

// ------------- Strucs & Constants -------------

// MaxQuestLength : max num of characters
// a question can have
const MaxQuestLength = 7000

// Question represents a question in the TA queue.
// A question can have multiple students associated with it.
type Question struct {
	QuestionPosition int      `json:"questPos" bson:"questPos"`
	OfficeHourID     string   `json:"offHourID" bson:"offHourID"`
	QuestionBody     string   `json:"questBody" bson:"questBody"`
	Students         []string `json:"students" bson:"students"`
	QuestionType     string   `json:"questType" bson:"questType"`
}

// QuestionCollection represents a connection to the
// question collection in our DB
type QuestionCollection MongoCollection
