package socket

import (
	"awise-messenger/models"
	"encoding/json"
	"errors"
	"log"
)

// onRead event
func onRead(customer *Customer) error {
	log.Println("Action onRead")

	// Update all messages by the target from the conversation in 'read'
	message := models.Message{IDConversation: customer.Info.ConversationID, IDUser: customer.Info.TargetID}
	message.UpdateMessageRead()

	// Research conversation by token
	conversation := models.Conversation{IDConversation: customer.Info.ConversationID}
	if err := conversation.FindOne(); err != nil {
		return errors.New("Conversation not find")
	}

	// Update conversation status
	if conversation.IDStatus != 2 {
		conversation.IDStatus = 2
		conversation.Update()
	}

	return nil
}

// OnSend Action
type OnSend struct {
	Message string
}

// onSend event
func onSend(transactional Transactionnal, customer *Customer) (models.Message, error) {
	log.Println("Action onSend")

	// parsing action
	action := OnSend{}
	if err := json.Unmarshal([]byte(transactional.Data), &action); err != nil {
		return models.Message{}, errors.New("Error parsing action")
	}

	log.Printf("Message: %s", action.Message)

	// start find conversation
	conversation := models.Conversation{IDConversation: customer.Info.ConversationID}
	if err := conversation.FindOne(); err != nil {
		return models.Message{}, errors.New("Conversation not find")
	}

	log.Printf("Conversation find %d", conversation.IDConversation)

	// Create message
	message := models.Message{IDUser: customer.Info.UserID, IDStatus: 1, IDConversation: conversation.IDConversation, Message: action.Message}
	if err := message.Create(); err != nil {
		return models.Message{}, errors.New("Error create message")
	}

	log.Printf("Message created")

	// Update status conversation and IDFirst/IDLast
	conversation.IDStatus = 1
	conversation.IDLastMessage = message.IDMessage
	if conversation.IDFirstMessage == 0 {
		conversation.IDFirstMessage = message.IDMessage
	}

	if err := conversation.Update(); err != nil {
		return models.Message{}, errors.New("Error update conversation")
	}

	log.Printf("Conversation updated")

	return message, nil
}

// onClose action
func onClose(customer *Customer) error {
	log.Println("Action onClose")
	deleteCustomer(customer)
	log.Println("Customer deleted")
	return nil
}

// OnLoad Action
type OnLoad struct {
	Token string
}

// onLoad action
func onLoad(transactional Transactionnal, customer *Customer) error {
	log.Println("Action onLoad")

	// parsing action
	action := OnLoad{}
	if err := json.Unmarshal([]byte(transactional.Data), &action); err != nil {
		return errors.New("Error parsing")
	}

	log.Printf("Token: %s", action.Token)

	// check if user already exist
	if exists := getCustomerByToken(action.Token); exists != nil {
		// delete if he exist
		log.Println("User already exist, delete in progress ...")
		deleteCustomer(exists)
	}

	// start find conversation
	conversation := models.Conversation{TokenCreator: action.Token, TokenReceiver: action.Token}
	if err := conversation.FindOneByTokenCreator(); err != nil {
		if err := conversation.FindOneByTokenReceiver(); err != nil {
			return errors.New("Conversation does not find")
		}
		customer.Info.UserID = conversation.IDReceiver
		customer.Info.TargetID = conversation.IDCreator
	}

	// assign the customer at the id and token
	if customer.Info.UserID == 0 {
		customer.Info.UserID = conversation.IDCreator
		customer.Info.TargetID = conversation.IDReceiver
	}

	customer.Info.Token = action.Token
	customer.Info.ConversationID = conversation.IDConversation

	log.Printf("Conversation find customer update (UserID: %d, TargetID: %d, Token: %s, ConversationID: %d", customer.Info.UserID, customer.Info.TargetID, customer.Info.Token, customer.Info.ConversationID)

	// add customer
	Customers = append(Customers, customer)

	return nil
}
