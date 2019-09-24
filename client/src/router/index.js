import Vue from 'vue';
import VueRouter from 'vue-router';

import Home from './home';

Vue.use(VueRouter);

export const router = new VueRouter({
  mode: 'history',
  routes: [{ path: '/', component: Home }],
});
