<template>
  <layout>
    <router-view></router-view>
  </layout>
</template>

<script>
import Layout from "./layout/index";
export default {
  name: "app",
  components: { Layout }
};
</script>