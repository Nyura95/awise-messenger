# Socket server

## A customers
The customer is represented by a interface. She is a Transactionnal struct for discuss with the client

```go
type Transactional struct {
	Message string
	Token   string
	ID      int
	Action  string
}

type Customer struct {
	Ws   *websocket.Conn
	Send Transactional
}
```
