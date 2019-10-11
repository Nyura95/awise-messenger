package socketv2

import "log"

// Action test
type Action struct{}

// NewAction action
func NewAction() *Action {
	return &Action{}
}

func (action *Action) send(message []byte) {
	log.Println("ok")
}
