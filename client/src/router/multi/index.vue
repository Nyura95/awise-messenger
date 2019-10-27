<template>
  <Layout>
    <div class="row">
      <div class="col-5 mx-auto room">
        <Multirooms
          name="Alexis"
          token="j_MuTU05PMBfh_A5t6077B4R5tw="
          tokenApi="token1"
          :target="2"
          :id="1"
          :nbCustomers="110"
        />
      </div>
      <div class="col-5 mx-auto room">
        <Multirooms
          name="Valentin"
          token="1ouZBXoXyqX7iyW3m2BH1O7Xakk="
          tokenApi="token2"
          :target="1"
          :id="2"
          :nbCustomers="115"
        />
      </div>
    </div>
  </Layout>
</template>

<script>
import Layout from "../../layout";
import Multirooms from "../../component/multirooms";
import { fetch } from "../../plugings/request";
export default {
  name: "Home",
  components: {
    Layout,
    Multirooms
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