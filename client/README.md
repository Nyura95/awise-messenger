# Messenger GoLang

## Récupération / Création d'une conversation  
Pour pouvoir faire ça tu dois avoir un token valide, créer par l'api nodeJS. Il me servira pour te valider en tant qu'utilisateur connecté. 

```
URL : https://messenger-api.appointrip.com/api/v1/conversation
HEADER : {
  Authorization: "tokenValide"
}
BODY : {
  idReceiver: "idUtilisateur"(en int)
}

Réponse type :
{
    "StatusCode": 200,
    "Reason": 1,
    "Comment": "ok",
    "Success": true,
    "Data": {
        "Conversation": {
            "IDConversation": 1,
            "UniqHash": "19",
            "Token": "s_DH9rt2OvG-kdnnTqv-sZncHx8=",
            "Title": "Title",
            "IDCreator": 1,
            "IDReceiver": 18,
            "IDLastMessage": 0,
            "IDFirstMessage": 0,
            "IDStatus": 3,
            "CreatedAt": "2019-03-08T16:52:09Z",
            "UpdatedAt": "2019-03-08T16:52:09Z"
        },
        "Messages": [
            {
                "IDMessage": 1,
                "IDUser": 1,
                "IDConversation": 1,
                "Message": "Hello, world!",
                "IDStatus": 1,
                "CreatedAt": "2019-03-10T00:12:27Z",
                "UpdatedAt": "2019-03-10T00:12:27Z"
            },
            {
                "IDMessage": 2,
                "IDUser": 1,
                "IDConversation": 1,
                "Message": "Hello, world!",
                "IDStatus": 1,
                "CreatedAt": "2019-03-10T00:14:10Z",
                "UpdatedAt": "2019-03-10T00:14:10Z"
            },
            ...
        ],
        "Users": [
            {
                "UserID": 1,
                "Email": "aa@qq.qq",
                "Password": "06dbf3a1a766116a351c45a41fc73de5",
                "Phone": "",
                "Avatars": ""
            },
            {
                "UserID": 18,
                "Email": "guide@qq.qq",
                "Password": "a8174f84cfa339b4a79290625b510b8e",
                "Phone": "",
                "Avatars": "https://image.appointrip.com/7tvpjlap3ti1z3mefk6asf.jpg"
            }
        ]
    }
}
```

## Connexion au socket
Pour pouvoir se connecter au socket, le token de la conversation est obligatoire (cela me permet de reconnaitre la discussion) sinon tu seras deconnecté du socket.

### Exemple d'un 'send' client

```js
window.onload = function () {
  sock = new WebSocket("ws://messenger.appointrip.com"); // connexion au socket
  sock.onopen = function () { // fonction à l'ouverture du socket
    sock.send(JSON.stringify({
      token: "s_DH9rt2OvG-kdnnTqv-sZncHx8=", // envoi du token
      id: 1, // envoi de l'id client actelle (plus tard ce sera le token api je pense, je verrais)
      action: "onload" // l'action que tu appels dans le back, ici onload
    }));
  };

  sock.onclose = function (e) {
    console.log('connection closed (' + e.code + ')'); // Les erreurs éventuelles
  };

  sock.onmessage = function (e) {
    console.log('message received: ' + e.data); // Récéption des messages sous la même forme que l'envoi { token: "s_DH9rt2OvG-kdnnTqv-sZncHx8=", id: 1, action: "send", message: "...." }
  };

   sock.send(JSON.stringify({ // Envoi d'un message
      action: "send", // l'action est 'send'
      message: "Bonjour !" // le message
    }));
};
```
L'achitecture est faite pour pouvoir créer plusieurs actions qui viendront plus tard.

### Interface Transactionelle entre le client et le back
Ceci est le SEUL payload accepté par le back. S'il n'est pas bon, pour le moment tu vas simplement avoir une erreur.
```typescript
interface Transactionnal {
  token: string;
  id: number;
  action: 'send' | 'onload'; // pour le moment
  message: string;
}
```

## Les statuts
Il y a simplement deux status possibles pour les conversations et les messages
```
1: En attente (Message non lu)
2: Lu (Aucun message non lu)
```