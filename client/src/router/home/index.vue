<template>
  <Layout>
    <div class="row">
      <div class="col-12">Bonjour</div>
    </div>
  </Layout>
</template>

<script>
import Layout from "../../layout";
import Socket from "../../plugings/socket";
import { fetch } from "../../plugings/request";
export default {
  name: "Home",
  components: {
    Layout
  },
  mounted: function() {
    this.init();
  },
  unmount: function() {
    // this.close();
  },
  data: function() {
    return { socket: null, message: "" };
  },
  methods: {
    init() {
      document.cookie = "X-Authorization=" + "dqsdqsdDQSDQS" + "; path=/";
      var conn = new WebSocket("ws://localhost:3001");

      conn.onopen = () => {
        conn.send("Bonjour");
      };

      conn.onclose = function(evt) {
        // var item = document.createElement("div");
        // item.innerHTML = "<b>Connection closed.</b>";
        console.log("close");
      };
      conn.onmessage = function(evt) {
        var messages = evt.data.split("\n");
        for (var i = 0; i < messages.length; i++) {
          //  var item = document.createElement("div");
          // item.innerText = messages[i];
          console.log(messages[i]);
        }
      };

      // this.socket = new Socket("ws://localhost:3001");
      // this.socket.onmessage = this.newMessage;
      // this.socket.init(
      //   function() {
      //     // this.$store.dispatch("conversation/getConversation", 18);
      //   }.bind(this)
      // );
      // this.$store.watch(
      //   (state, getters) => getters["conversation/getTokenConversation"],
      //   newValue => {
      //     this.socket.toConversation(newValue);
      //   }
      // );
    },
    sendMessage(msg) {
      // this.socket.sendMessage(msg);
    },
    newMessage(message) {
      console.log(message);
      // if (message.Action === "newMessage") {
      //   this.message = "";
      //   this.$store.dispatch("conversation/addMessage", message.Data);
      // }
    },
    increment(add) {
      this.$store.dispatch("counter/increment", add);
    },
    close() {
      // this.socket.close();
    }
  },
  computed: {
    messages() {
      return this.$store.state.conversation.messages;
    }
  }
};
</script>
