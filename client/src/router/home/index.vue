<template>
  <div>
    {{ $t('counter', { counter: counter }) }}
    <button @click.prevent="close()">Increment</button>
  </div>
</template>

<script>
import Socket from "../../plugings/socket";

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
          this.socket.toConversation("X8JzNaGnELklxc2qjO0_4VYznfw=");
        }.bind(this)
      );
    },
    sendMessage() {
      this.socket.sendMessage("send");
    },
    newMessage(message) {
      console.log(message);
      if (message.Action === "close") {
        this.init();
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
    counter() {
      return this.$store.state.counter.counter;
    }
  }
};
</script>
