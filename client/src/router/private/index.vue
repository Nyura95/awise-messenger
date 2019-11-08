<template>
  <Layout>
    <div class="row">
      <div class="col-5 mx-auto room">
        <RommPrivate name="Steeven" tokenApi="token3" :target="2" :id="3" />
      </div>
      <div class="col-5 mx-auto room">
        <RommPrivate name="Valentin" tokenApi="token2" :target="3" :id="2" />
      </div>
    </div>
  </Layout>
</template>

<script>
import Layout from "../../layout";
import RommPrivate from "../../component/roomPrivate";
import { fetch } from "../../plugings/request";
export default {
  name: "Home",
  components: {
    Layout,
    RommPrivate
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