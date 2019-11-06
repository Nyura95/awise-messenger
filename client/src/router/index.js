import Vue from 'vue';
import VueRouter from 'vue-router';

import Home from './home';
import Charge from './charge';
import Multi from './multi';
import VueChatScroll from 'vue-chat-scroll';
Vue.use(VueChatScroll);

Vue.use(VueRouter);

export const router = new VueRouter({
  mode: 'history',
  routes: [
    { path: '/', name: 'home', component: Home, meta: { requiresAuth: true } },
    { path: '/charge', name: 'charge', component: Charge, meta: { requiresAuth: true } },
    { path: '/multi', name: 'multi', component: Multi, meta: { requiresAuth: true } },
  ],
});

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    // this route requires auth, check if logged in
    // if not, redirect to login page.
    // if (!auth.loggedIn()) {
    // next({
    //   path: '/login',
    //   query: { redirect: to.fullPath },
    // });
    // } else {
    //   next();
    // }

    next();
  } else {
    next(); // make sure to always call next()!
  }
});
