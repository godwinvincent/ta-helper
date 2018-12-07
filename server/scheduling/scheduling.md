### About
**scheduling** is responsible for managing `questions`, `office hours`, and `users`.

### Websocket notifications
Gateway can expect to receive a marshaled Struc that contains an array of usernames affected by the event and a `event` string. 

```
type WebsocketMsg struct {
	Usernames []string `json:"usernames"`
	Event string `json:"event"`
}
```

**Possible events:**
- "question-yourTurn"
- "question-new"
- "question-deleted" 
- "office-new"
- "office-deleted"

**TA has begun answering a question**
Hit the endpoint below and pass in the question being answered. User must be an instructor. It notifies all students in that question that it is their turn.
Example: `POST /v1/ws/?qid=<QuestionID>`
