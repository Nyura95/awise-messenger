import Vue from 'vue';
import App from './App';

import { router } from './router';
import store from './store';
import i18n from './i18n';

import 'bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';

Vue.config.productionTip = false;

/* eslint-disable no-new */
new Vue({
  el: '#main',
  template: '<App/>',
  components: { App },
  router,
  store,
  i18n,
});
