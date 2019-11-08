<template>
  <Layout>
    <div class="row">
      <div class="col-4 mx-auto room">
        <Room name="Alexis" tokenApi="token1" :idconversation="16" :id="1" />
      </div>
      <div class="col-4 mx-auto room">
        <Room name="Valentin" tokenApi="token2" :idconversation="16" :id="2" />
      </div>
      <div class="col-4 mx-auto room">
        <Room name="Steeven" tokenApi="token3" :idconversation="16" :id="3" />
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