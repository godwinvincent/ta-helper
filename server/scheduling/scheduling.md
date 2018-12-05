

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
