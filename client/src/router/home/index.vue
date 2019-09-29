<template>
  <div class="container">
    <div class="row" v-for="item in messages" :key="item.IDMessage">
      <div class="col-12">{{ item.Message }}</div>
    </div>
    <div class="row">
      <div class="col-12">
        <div class="row">
          <div class="col-auto">
            <input type="text" class="form-control" v-model="message" />
          </div>
          <div class="col-auto">
            <div class="btn btn-primary" v-on:click="sendMessage(message)">Envoyer</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Socket from "../../plugings/socket";
import { fetch } from "../../plugings/request";
export default {
  name: "Home",
  mounted: function() {
    this.init();
  },
  unmount: function() {
    this.close();
  },
  data: function() {
    return { socket: null, message: "" };
  },
  methods: {
    init() {
      this.socket = new Socket("ws://localhost:3001");
      this.socket.onmessage = this.newMessage;
      this.socket.init(
        function() {
          this.$store.dispatch("conversation/getConversation", 18);
        }.bind(this)
      );

      this.$store.watch(
        (state, getters) => getters["conversation/getTokenConversation"],
        newValue => {
          this.socket.toConversation(newValue);
        }
      );
    },
    sendMessage(msg) {
      this.socket.sendMessage(msg);
    },
    newMessage(message) {
      if (message.Action === "newMessage") {
        this.message = "";
        this.$store.dispatch("conversation/addMessage", message.Data);
      }
    },
    increment(add) {
      this.$store.dispatch("counter/increment", add);
    },
    close() {
      this.socket.close();
    }
  },
  computed: {
    messages() {
      return this.$store.state.conversation.messages;
    }
  }
};
</script>
