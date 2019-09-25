<template>
  <div>
    {{ $t('counter', { counter: counter }) }}
    <button @click.prevent="close()">increment</button>
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
    return { socket: null };
  },
  methods: {
    init() {
      this.socket = new Socket("ws://localhost:3001");
      this.socket.onmessage = this.newMessage;
      this.socket.init(
        function() {
          // this.$store.dispatch("conversation/getConversation", 18);
        }.bind(this)
      );
      console.log(this.$store);
      // this.$store.watch(
      //   (state, getters) => {
      //     console.log(getters);
      //     return getters.getConversation;
      //   },
      //   (newValue, oldValue) => {
      //     console.log(`Updating from ${oldValue} to ${newValue}`);
      //   }
      // );
    },
    sendMessage() {
      this.socket.sendMessage("send");
    },
    newMessage(message) {
      console.log(message);
    },
    increment(add) {
      this.$store.dispatch("counter/increment", add);
    },
    close() {
      this.socket.close();
    }
  },
  computed: {
    counter() {
      return this.$store.state.counter.counter;
    }
  }
};
</script>
