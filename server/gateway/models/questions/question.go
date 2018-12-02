package questions

// ------------- Strucs & Constants -------------
const MaxQuestLength = 7000

// Question represents a question in the TA queue.
// A question can have multiple students associated with it.
type Question struct {
	QuestionPosition int      `json:"questPos" bson:"questPos"`
	QuestionBody     string   `json:"questBody" bson:"questBody"`
	Students         []string `json:"students" bson:"students"`
	QuestionType     string   `json:"questType" bson:"questType"`
}
