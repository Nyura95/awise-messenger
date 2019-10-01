<template>
  <div class="container container_tchat">
    <div class="row">
      <div class="col-12" v-for="item in messages" :key="item.IDMessage">{{item.Messages}}</div>
    </div>
    <div class="row">
      <div class="col-12">
        <input type="button" value="Envoyer" class="btn btn-primary" />
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Tchat",
  props: {
    messages: Array
  }
};
</script>