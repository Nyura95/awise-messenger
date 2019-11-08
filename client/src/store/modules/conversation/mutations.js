import { SET_HASH } from './mutation-types';

export default {
  [SET_HASH](state, payload) {
    state.hashs[payload.idconversation] = payload.hash;
  },
};
