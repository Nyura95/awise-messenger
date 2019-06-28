package socket

import (
	"errors"

	"golang.org/x/net/websocket"
)

// Info fddsfsdf
type Info struct {
	Token          string
	UserID         int
	TargetID       int
	ConversationID int
}

// Customer general
type Customer struct {
	Ws   *websocket.Conn
	Info Info
}

// Customers stock all customer online
var Customers []*Customer

func deleteCustomer(customer *Customer) {
	customer.Ws.Close()
	q := make([]*Customer, 0)
	if len(Customers) == 1 {
		Customers = q
		return
	}

	for _, n := range Customers {
		if n.Info.UserID != customer.Info.UserID {
			q = append(q, n)
		}
	}
	Customers = q
}

func getCustomerByID(id int) *Customer {
	for _, n := range Customers {
		if id == n.Info.UserID {
			return n
		}
	}
	return nil
}

func getCustomerByToken(token string) *Customer {
	for _, n := range Customers {
		if token == n.Info.Token {
			return n
		}
	}
	return nil
}

func (customer Customer) sendMessage(transaction Transactionnal) error {
	message, err := encryptMessage(transaction)
	if err != nil {
		return err
	}
	if err = websocket.Message.Send(customer.Ws, message); err != nil {
		return err
	}
	return nil
}

func (customer Customer) sendMessageToCustomer(id int, transaction Transactionnal) error {

	target := getCustomerByID(id)

	if target == nil {
		return errors.New("This customer is not connected")
	}

	message, err := encryptMessage(transaction)
	if err != nil {
		return err
	}
	if err = websocket.Message.Send(target.Ws, message); err != nil {
		return err
	}
	return nil
}
