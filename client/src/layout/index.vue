<template>
  <div>
    <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
      <router-link :to="{ name: 'home' }" class="navbar-brand">Awise messenger</router-link>

      <button class="navbar-toggler" type="button" @click="toggleMenu">
        <span class="navbar-toggler-icon" />
      </button>

      <div :class="{ show : menuCollapsed}" class="collapse navbar-collapse">
        <ul class="navbar-nav mr-auto">
          <div class="nav-item" v-for="user in users" :key="user.UserID">
            <Avatar :url="user.Avatars" :online="user.Online" />
          </div>
          <router-link :to="{ name: 'home' }" active-class="active" class="nav-item" tag="li">
            <a class="nav-link">Home</a>
          </router-link>
          <router-link :to="{ name: 'multi' }" active-class="active" class="nav-item" tag="li">
            <a class="nav-link">Multi</a>
          </router-link>
        </ul>
        <span class="navbar-text">
          <a class="btn btn-secondary" href="#" @click.prevent="logout">
            <i class="fa fa-sign-out" />
          </a>
        </span>
      </div>
    </nav>

    <div class="container pt-4">
      <slot />
    </div>
  </div>
</template>

<script>
import Avatar from "../component/avatars";
export default {
  name: "Layout",
  components: {
    Avatar
  },
  data() {
    return {
      menuCollapsed: false
    };
  },
  mounted: function() {
    // this.$store.dispatch("users/getAllUsers");
  },
  methods: {
    logout() {
      console.log("logout");
    },
    toggleMenu() {
      this.menuCollapsed = !this.menuCollapsed;
    }
  },
  computed: {
    users() {
      return this.$store.state.users.accounts;
    }
  }
};
</script>