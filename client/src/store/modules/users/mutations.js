import { SET_USERS } from './mutation-types';

export default {
  [SET_USERS](state, payload) {
    state.accounts = [...state.accounts, ...payload];
  },
};
