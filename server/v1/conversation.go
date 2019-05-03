package v1

import (
	"awise-messenger/helpers"
	"awise-messenger/models"
	"awise-messenger/server/response"
	"encoding/json"
	"log"
	"net/http"
	"sort"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

type listConversations struct {
	*models.Conversation
	Message models.Message
	Target  models.User
	NotRead int
}

type allConversations []listConversations

func (p allConversations) Len() int {
	return len(p)
}

func (p allConversations) Less(i, j int) bool {
	return p[i].Message.UpdatedAt.After(p[j].Message.UpdatedAt)
}

func (p allConversations) Swap(i, j int) {

	p[i], p[j] = p[j], p[i]
}

type oneConversation struct {
	Conversation models.Conversation
	Messages     []*models.Message
	User         models.User
	Target       models.User
	NotRead      int
}

// GetAllConvo for the user
func GetAllConvo(w http.ResponseWriter, r *http.Request) {
	log.Println("Get all conversation !")
	// Get ID token
	idUser := context.Get(r, "user_id").(int)

	// Get all conversation in relationship with the user
	conversations, _ := models.FindAllConversationByIDUser(idUser)
	log.Printf("Conversations found : %d", len(conversations))

	// Assign the conversation struct with the extended struct listConversations
	allConversation := make(allConversations, 0)
	userTarget := make(chan models.User)
	for _, conv := range conversations {
		// Start research message of the current conversation

		var IDuserTarget int
		if conv.IDCreator != idUser {
			IDuserTarget = conv.IDCreator
		} else {
			IDuserTarget = conv.IDReceiver
		}

		go helpers.GetUser(IDuserTarget, userTarget)

		message := models.Message{IDMessage: conv.IDLastMessage}
		message.FindOne()

		messagesNotRead, _ := models.FindAllMessageNotRead(conv.IDConversation, IDuserTarget)

		log.Printf("For the conversation %d, message : (%s)", conv.IDConversation, message.Message)
		// Show the last massage
		allConversation = append(allConversation, listConversations{Message: message, Conversation: conv, Target: <-userTarget, NotRead: len(messagesNotRead)})
	}

	log.Println("Sort")
	// Sort the array on the update_at in the last message
	sort.Sort(allConversation)

	log.Println("Return conversatio")
	json.NewEncoder(w).Encode(response.BasicResponse(allConversation, "ok", 1))
}

// GetConvo for get or create a conversation with uniq_id
func GetConvo(w http.ResponseWriter, r *http.Request) {
	log.Println("Get a conversation !")
	// Get ID token
	idUser := context.Get(r, "user_id").(int)

	// Get id target from the params
	idTarget, err := helpers.StringToInt(mux.Vars(r)["id"])
	if idTarget == 0 || err != nil {
		log.Printf("The params is wrong ! %s", mux.Vars(r))
		json.NewEncoder(w).Encode(response.BasicResponse(new(interface{}), "Body is wrong (params)", -1))
		return
	}

	if idUser == idTarget {
		log.Printf("The id's are similar !")
		json.NewEncoder(w).Encode(response.BasicResponse(new(interface{}), "The id's are similar !", -2))
		return
	}

	// Seek users in goroutine
	user := make(chan models.User)
	target := make(chan models.User)
	go helpers.GetUser(idUser, user)
	go helpers.GetUser(idTarget, target)

	// create the uniq hash
	uniqHash := helpers.Uniqhash(idUser, idTarget)
	log.Printf("uniq hash : %s", uniqHash)

	// find conv with the hash_uniq
	conversation := models.Conversation{UniqHash: uniqHash}
	if err = conversation.FindOneByHash(); err != nil {
		// If conversation is not find
		log.Println("Conversation is not find, creation in progress ...")
		// Default value
		conversation.Title = "Title"
		conversation.Token = helpers.Token(uniqHash)
		conversation.IDCreator = idUser
		conversation.IDReceiver = idTarget
		conversation.IDStatus = 1

		// Create a new conversation
		if err = conversation.Create(); err != nil {
			log.Printf("Error when creating the conversation ! (%v)", conversation)
			json.NewEncoder(w).Encode(response.BasicResponse(new(interface{}), "Error creating a conversation", -3))
			return
		}
		log.Printf("Conversation create (%v)", conversation)

		log.Println("Return conversation")
		// return response
		json.NewEncoder(w).Encode(response.BasicResponse(oneConversation{Conversation: conversation, Messages: make([]*models.Message, 0), User: <-user, Target: <-target}, "ok", 1))
		return
	}
	log.Printf("Conversation find (%v)", conversation.IDConversation)
	// Seek all messages
	messages, _ := models.FindAllMessageByIDConversation(conversation.IDConversation)

	// Count message not read
	messagesNotRead, _ := models.FindAllMessageNotRead(conversation.IDConversation, idTarget)

	log.Println("Return conversation")
	// return response
	json.NewEncoder(w).Encode(response.BasicResponse(oneConversation{Conversation: conversation, Messages: messages, User: <-user, Target: <-target, NotRead: len(messagesNotRead)}, "ok", 1))
}
