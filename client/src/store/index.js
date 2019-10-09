import Vue from 'vue';
import Vuex from 'vuex';
import VuexPersist from 'vuex-persist';
import createLogger from 'vuex/dist/logger';

Vue.use(Vuex);

import filter from './modules/filter';

import counter from './modules/counter';
import conversation from './modules/conversation';
import users from './modules/users';

const debug = process.env.NODE_ENV !== 'production';

const vuexLocalStorage = new VuexPersist({
  key: 'vuex',
  storage: window.localStorage,
  filter,
});

export default new Vuex.Store({
  modules: {
    counter,
    conversation,
    users,
  },
  strict: debug,
  plugins: [vuexLocalStorage.plugin, createLogger()],
});
