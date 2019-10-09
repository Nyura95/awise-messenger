<template>
  <div class="avatar">
    <img :src="url" />
    <span class="online" v-if="online"></span>
  </div>
</template>

<script>
export default {
  name: "Avatar",
  props: {
    url: String,
    online: {
      type: Boolean,
      default: false
    }
  }
};
</script>

<style lang="scss" scoped>
.avatar {
  position: relative;
  img {
    height: 40px;
    width: 40px;
    border-radius: 50%;
    margin-right: 10px;
    margin-left: 10px;
  }
  .online {
    position: absolute;
    height: 5px;
    width: 5px;
    border-radius: 50%;
    background-color: lightgreen;
    bottom: 0px;
    right: 10px;
  }
}
</style>