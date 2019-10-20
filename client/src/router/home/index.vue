<template>
  <Layout>
    <div class="row">
      <div class="col-5 mx-auto room">
        <Room name="Alexis" token="j_MuTU05PMBfh_A5t6077B4R5tw=" tokenApi="token1" :target="2" />
      </div>
      <div class="col-5 mx-auto room">
        <Room name="Valentin" token="1ouZBXoXyqX7iyW3m2BH1O7Xakk=" tokenApi="token2" :target="1" />
      </div>
    </div>
  </Layout>
</template>

<script>
import Layout from "../../layout";
import Room from "../../component/room";
import { fetch } from "../../plugings/request";
export default {
  name: "Home",
  components: {
    Layout,
    Room
  },
  data: function() {
    return {};
  },
  methods: {
    increment(add) {
      this.$store.dispatch("counter/increment", add);
    }
  },
  computed: {
    messages() {
      return this.$store.state.conversation.messages;
    }
  }
};
</script>

<style scoped>
.room {
  border: 1px solid black;
}
</style>