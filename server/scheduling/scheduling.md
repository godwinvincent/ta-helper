

### Websocket notifications
Gateway can expect to receive a marshaled Struc that contains an array of usernames affected by the event and a `event` string. **Possible events:**
- "question-new"
- "question-deleted" 
- "question-edited" 
- "question-student" 
- "office-new"
- "office-deleted"