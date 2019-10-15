<template>
  <Layout>
    <div class="row">
      <div class="col-6">
        <div class="row">
          <div class="col-12">
            token :
            <input v-model="token" type="text" />
          </div>
          <div class="col-12">
            target :
            <input v-model="target" type="text" />
          </div>
          <div class="col-12">
            <input type="button" @click="start" value="submit" />
          </div>
          <div class="col-12">
            message :
            <input v-model="message" type="text" />
          </div>
          <div class="col-12">
            <input type="button" @click="sendMessage" value="submit" />
          </div>
        </div>
      </div>
      <div class="col-6">
        <div class="row">
          <div class="col-12">
            token2 :
            <input v-model="token2" type="text" />
          </div>
          <div class="col-12">
            target2 :
            <input v-model="target2" type="text" />
          </div>
          <div class="col-12">
            <input type="button" @click="start2" value="submit" />
          </div>
          <div class="col-12">
            message :
            <input v-model="message2" type="text" />
          </div>
          <div class="col-12">
            <input type="button" @click="sendMessage2" value="submit" />
          </div>
        </div>
      </div>
    </div>
  </Layout>
</template>

<script>
import Layout from "../../layout";
import Socket from "../../plugings/socket";
import { fetch } from "../../plugings/request";
let conn = null;
let conn2 = null;
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
    return {
      socket: null,
      token: "token1",
      target: "2",
      message: "",
      token2: "token2",
      target2: "1",
      message2: ""
    };
  },
  methods: {
    sendMessage() {
      conn.send(this.message);
    },
    sendMessage2() {
      conn2.send(this.message2);
    },
    start() {
      if (conn) {
        conn.close();
      }
      conn = new WebSocket(`ws://localhost:3001/${this.token}/${this.target}`);

      conn.onopen = () => {
        console.log("onopen");
      };

      conn.onclose = function(evt) {
        console.log("close");
      };

      conn.onmessage = function(evt) {
        var messages = evt.data.split("\n");
        for (var i = 0; i < messages.length; i++) {
          console.log("1 :", messages[i]);
        }
      };

      conn.onerror = function(err) {
        console.log(err);
      };
    },
    start2() {
      if (conn2) {
        conn2.close();
      }
      conn2 = new WebSocket(
        `ws://localhost:3001/${this.token2}/${this.target2}`
      );

      conn2.onopen = () => {
        console.log("onopen2");
      };

      conn2.onclose = function(evt) {
        console.log("close2");
      };

      conn2.onmessage = function(evt) {
        var messages = evt.data.split("\n");
        for (var i = 0; i < messages.length; i++) {
          console.log("2 :", messages[i]);
        }
      };

      conn2.onerror = function(err) {
        console.log(err);
      };
    },
    init() {
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
