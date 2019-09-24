<template>
  <div>
    {{ $t('counter', { counter: counter }) }}
    <button @click.prevent="increment(1)">Increment</button>
  </div>
</template>

<script>
import Socket from "../../plugings/socket";

export default {
  name: "Home",
  methods: {
    increment(add) {
      this.$store.dispatch("counter/increment", add);
    }
  },
  mounted: function() {
    this.init();
  },
  data: function() {
    return { socket: null };
  },
  methods: {
    init: function() {
      this.socket = new Socket("ws://localhost:3001");
      this.socket.init(
        function() {
          this.socket.sendMessage(
            "onload",
            JSON.stringify({ token: "dfsdfsf" })
          ); // init user with token
        }.bind(this)
      );
    }
  },
  computed: {
    counter() {
      return this.$store.state.counter.counter;
    }
  }
};
</script>
