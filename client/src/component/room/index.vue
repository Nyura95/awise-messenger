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
      message :
      <input v-model="message" type="text" @keyup.enter="sendMessage" />
      <input type="button" @click="sendMessage" value="submit" />
    </div>
    <div class="col-12 mt-4">
      <div class="row" v-for="(message, key) in messages" :key="key">
        <div class="col-2">{{accounts[message.IDAccount]}}:</div>
        <div class="col-10">{{message.Message}}</div>
      </div>
    </div>
  </div>
</template>

<script>
import { fetch } from "../../plugings/request";
export default {
  name: "Room",
  props: {
    name: String,
    token: String,
    target: Number,
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
      error: false
    };
  },
  methods: {
    sendMessage() {
      this.socket.send(this.message);
    },
    close() {
      this.socket.close();
      this.socket = null;
    },
    chargeMessage() {
      fetch(
        "/api/v2/conversations/target/" + this.target,
        "get",
        {},
        {
          Authorization: this.tokenApi
        }
      ).then(result => {
        this.messages = result.Data.Messages;
        for (let i = 0; i < result.Data.Accounts.length; i++) {
          this.accounts[result.Data.Accounts[i].ID] =
            result.Data.Accounts[i].Firstname;
        }
        console.log(result);
      });
    },
    start() {
      this.error = false;
      if (this.socket) {
        this.socket.close();
      }
      this.socket = new WebSocket(`ws://localhost:3001/${this.token}`);

      this.socket.onopen = () => {
        this.log("onopen");
        this.chargeMessage();
        this.open = true;
      };

      this.socket.onclose = evt => {
        this.log("close");
        this.messages = [];
        this.open = false;
      };

      this.socket.onmessage = evt => {
        const messages = evt.data.split("\n");
        for (let i = 0; i < messages.length; i++) {
          const message = JSON.parse(messages[i]);
          this.log(messages[i]);
          if (message.Action === "Message") {
            this.message = "";
            this.messages.push(message.Message);
          }
          if (message.Action === "Connection") {
            new Notification("Nouvelle connexion", {
              body: `l'utilisateur ${message.User} vient de se connecter`,
              lang: "FR",
              tag: new Date()
            });
          }
          if (message.Action === "Disconnection") {
            new Notification("Déconnection", {
              body: `l'utilisateur ${message.User} vient de se déconnecter`,
              lang: "FR",
              tag: new Date()
            });
          }
          if (message.Action === "Error") {
            this.error = true;
          }
        }
      };
      this.socket.onerror = err => {
        this.log(err);
        this.messages = [];
        this.error = true;
      };
    },
    log(...messages) {
      console.log(`${this.name}:`, ...messages);
    }
  }
};
</script>

<style lang="scss" scoped>
</style>