<template>
  <Layout>
    <div class="row">
      <div class="col-5 mx-auto room">
        <RoomTarget name="Alexis" tokenApi="token1" :target="2" :id="1" />
      </div>
      <div class="col-5 mx-auto room">
        <RoomTarget name="Valentin" tokenApi="token2" :target="1" :id="2" />
      </div>
    </div>
  </Layout>
</template>

<script>
import Layout from "../../layout";
import RoomTarget from "../../component/roomTarget";
import { fetch } from "../../plugings/request";
export default {
  name: "Home",
  components: {
    Layout,
    RoomTarget
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