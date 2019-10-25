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
  </div>
</template>

<script>
// accounts[message.IDAccount]
import { fetch } from "../../plugings/request";
export default {
  name: "Room",
  props: {
    name: String,
    id: Number,
    token: String,
    target: Number,
    tokenApi: String,
    nbCustomers: {
      type: Number,
      default: 100
    }
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
      sockets: [],
      message: "",
      accounts: {},
      open: false,
      error: false
    };
  },
  methods: {
    sendMessage() {
      for (let i = 0; i < this.sockets.length; i++) {
        this.sockets[i].send(this.message);
      }
    },
    close() {
      for (let i = 0; i < this.sockets.length; i++) {
        this.sockets[i].close();
      }
      this.sockets = [];
    },
    start() {
      this.error = false;
      for (let i = 0; i < this.sockets.length; i++) {
        this.sockets[i].close();
      }
      this.sockets = [];

      for (let i = 0; i < this.nbCustomers; i++) {
        this.sockets[i] = new WebSocket(`ws://localhost:3001/${this.token}`);
        this.sockets[i].onopen = () => {
          this.log(`onopen for ${this.name} (${i})`);
          this.open = true;
        };

        this.sockets[i].onclose = evt => {
          this.log(`close for ${this.name} (${i})`);
          this.open = false;
        };

        this.sockets[i].onmessage = evt => {
          const messages = evt.data.split("\n");
          for (let y = 0; y < messages.length; y++) {
            const message = JSON.parse(messages[y]);
            this.log(messages[y]);
            if (message.Action === "Message") {
              this.message = "";
            }
            if (message.Action === "Error") {
              this.error = true;
            }
          }
        };

        this.sockets[i].onerror = err => {
          this.log(err);
          this.error = true;
        };
      }
    },
    log(...messages) {
      console.log(`${this.name}:`, ...messages);
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