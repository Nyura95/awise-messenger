import Vue from 'vue';
import Vuex from 'vuex';
import VuexPersist from 'vuex-persist';
import createLogger from 'vuex/dist/logger';

Vue.use(Vuex);

import filter from './modules/filter';

import counter from './modules/counter';
import conversation from './modules/conversation';

const debug = process.env.NODE_ENV !== 'production';

const vuexLocalStorage = new VuexPersist({
  key: 'vuex', // The key to store the state on in the storage provider.
  storage: window.localStorage, // or window.sessionStorage or localForage
  // Function that passes the state and returns the state with only the objects you want to store.
  // reducer: state => state,
  // Function that passes a mutation and lets you decide if it should update the state in localStorage.
  filter,
});

export default new Vuex.Store({
  /**
   * Assign the modules to the store.
   */
  modules: {
    counter,
    conversation,
  },

  /**
   * If strict mode should be enabled.
   */
  strict: debug,

  /**
   * Plugins used in the store.
   */
  plugins: [vuexLocalStorage.plugin, createLogger()],
});
