<template>
  <div>
    <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
      <router-link :to="{ name: 'home' }" class="navbar-brand">Vue 2 Boilerplate</router-link>

      <button class="navbar-toggler" type="button" @click="toggleMenu">
        <span class="navbar-toggler-icon" />
      </button>

      <div :class="{ show : menuCollapsed}" class="collapse navbar-collapse">
        <ul class="navbar-nav mr-auto">
          <router-link :to="{ name: 'home' }" active-class="active" class="nav-item" tag="li">
            <a class="nav-link">Home</a>
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
      <div class="row">
        <div class="col col-12">
          <slot />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Layout",
  data() {
    return {
      menuCollapsed: false
    };
  },
  methods: {
    logout() {
      console.log("logout");
    },
    toggleMenu() {
      this.menuCollapsed = !this.menuCollapsed;
    }
  }
};
</script>
