<template>
  <Layout>
    <div class="row">
      <div class="col-5 mx-auto room">
        <Room
          name="Alexis"
          token="2Ds8UdMJQ3i40TOzEOFTe8aYVw4="
          tokenApi="token1"
          :target="2"
          :id="1"
        />
      </div>
      <div class="col-5 mx-auto room">
        <Room
          name="Valentin"
          token="J62tKYjZbp99mEClnHSks4E_XAU="
          tokenApi="token2"
          :target="1"
          :id="2"
        />
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