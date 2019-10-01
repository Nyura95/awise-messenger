<template>
  <div class="container">
    <Tchat :messages="$store.state.conversation.messages" />
  </div>
</template>

<script>
import Tchat from "../../component/tchat/container";
import Socket from "../../plugings/socket";
import { fetch } from "../../plugings/request";
export default {
  name: "Home",
  components: {
    Tchat
  },
  mounted: function() {
    this.init();
  },
  unmount: function() {
    this.close();
  },
  data: function() {
    return { socket: null, message: "", color: "black" };
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
          this.color = "white";
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
