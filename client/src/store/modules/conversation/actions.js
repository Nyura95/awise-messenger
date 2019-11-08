import { SET_HASH } from './mutation-types';

export const setHash = ({ commit }, payload) => {
  commit(SET_HASH, payload);
};

export default {
  setHash,
};
