<template>
  <div class="row">
    <div class="col-12">
      {{ name }}
      <input
        type="button"
        @click="start"
        value="connect"
        v-bind:class="connectButtonClass"
        class="btn"
      />
      <input
        type="button"
        @click="close"
        value="disconnect"
        v-bind:class="disconnectButtonClass"
        class="btn"
      />
    </div>
    <div class="text-danger" v-if="error">Une erreur est survenu !</div>

    <div class="col-12 mt-4">
      <div class="scroll" v-chat-scroll="{always: false, smooth: true}">
        <div class="row mt-2 mb-2 container_message" v-for="(message, key) in messages" :key="key">
          <div class="col-12" v-if="message.idAccount === id">
            <div class="row">
              <div class="col-auto message">
                <span @click="updateMessage(message.id)">{{message.message}}</span>
              </div>
            </div>
          </div>
          <div class="col-12" v-if="message.idAccount !== id">
            <div class="row">
              <div class="col-auto left message target">
                <span @click="updateMessage(message.id)">{{message.message}}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="col-12 mt-4">
      message :
      <input v-model="message" type="text" @keyup.enter="sendMessage" />
      <input type="button" @click="sendMessage" value="submit" />
    </div>
  </div>
</template>

<script>
// accounts[message.idAccount]
import { fetch } from "../../plugings/request";
import AwiseSocket from "../../plugings/socket/index";
export default {
  name: "Room",
  props: {
    id: Number,
    name: String,
    token: String,
    idconversation: Number,
    tokenApi: String
  },
  mounted: function() {
    if ("Notification" in window) {
      Notification.requestPermission(() => {
        if (this.open) {
          this.start();
        }
      });
    }
  },
  destroyed: function() {
    this.close();
  },
  computed: {
    connectButtonClass: function() {
      return {
        "btn-secondary": !this.open,
        "btn-success": this.open
      };
    },
    disconnectButtonClass: function() {
      return {
        "btn-danger": !this.open,
        "btn-secondary": this.open
      };
    }
  },
  data: function() {
    return {
      socket: null,
      message: "",
      messages: [],
      accounts: {},
      open: false,
      error: false,
      conversation: 0
    };
  },
  methods: {
    updateMessage(id) {
      if (this.message !== "") {
        this.update(id, this.message);
        return;
      }
      this.del(id);
    },
    update(id, message) {
      fetch(
        "/api/v2/conversations/" + this.conversation + "/messages/" + id,
        "put",
        {
          message
        },
        {
          Authorization: this.tokenApi
        }
      ).then(result => {
        console.log(result);
      });
    },
    del(id) {
      fetch(
        "/api/v2/conversations/" + this.conversation + "/messages/" + id,
        "delete",
        {},
        {
          Authorization: this.tokenApi
        }
      ).then(result => {
        console.log(result);
      });
    },
    sendMessage() {
      if (this.message !== "") {
        this.socket.send(this.message);
      }
    },
    close() {
      if (this.socket) {
        this.socket.close();
        this.socket = null;
      }
    },
    chargeMessage() {
      fetch(
        "/api/v2/conversations/" + this.idconversation,
        "get",
        {},
        {
          Authorization: this.tokenApi
        }
      ).then(result => {
        this.messages = result.data.messages.reverse();
        this.conversation = result.data.id;
        console.log(result);
      });
    },
    start() {
      this.error = false;
      if (this.socket) {
        this.socket.close();
      }

      this.socket = new AwiseSocket("ws://localhost:3001");

      this.socket.onclose = () => {
        this.messages = [];
        this.open = false;
      };

      this.socket.onerror = err => {
        console.log(err);
        this.messages = [];
        this.error = true;
      };

      this.socket.private = token => {
        console.log(token);
      };

      this.socket.message = message => {
        this.message = "";
        this.messages.push(message);
      };

      this.socket.update = message => {
        for (let i = 0; i < this.messages.length; i++) {
          if (message.id === this.messages[i].id) {
            this.messages[i] = message;
          }
        }
        this.messages = [...this.messages];
      };

      this.socket.delete = message => {
        for (let i = 0; i < this.messages.length; i++) {
          if (message.id === this.messages[i].id) {
            this.messages.splice(i, 1);
          }
        }
        this.messages = [...this.messages];
      };

      this.socket.connection = user => {
        new Notification("Nouvelle connexion", {
          body: `l'utilisateur ${user} vient de se connecter`,
          lang: "FR",
          tag: new Date()
        });
      };

      this.socket.disconnection = user => {
        new Notification("Déconnection", {
          body: `l'utilisateur ${user} vient de se déconnecter`,
          lang: "FR",
          tag: new Date()
        });
      };

      this.socket.error = (lockey, message) => {
        this.error = true;
      };

      this.socket.initConversation(this.token, () => {
        this.chargeMessage();
        this.open = true;
      });
    }
  }
};
</script>

<style lang="scss" scoped>
.left {
  margin-left: auto;
}

.scroll {
  overflow-y: auto;
  overflow-x: hidden;
  height: 150;
}
::-webkit-scrollbar {
  display: none;
}

.container_message {
  padding-left: 30px;
  padding-right: 30px;
  .message {
    padding: 5px;
    border-bottom-left-radius: 48px;
    border-top-left-radius: 48px;
    border-top-right-radius: 48px;
    border-bottom-right-radius: 48px;
    background-color: #3578e5;
  }

  .message span {
    margin-left: 5px;
    margin-right: 5px;
    color: white;
    font-weight: 500;
  }

  .target {
    background-color: lightgray;
  }
  .target span {
    color: black;
  }
}
</style>