import actions from './actions';
import getters from './getters';
import mutations from './mutations';
import state from './state';

export const filter = () => false;

export default {
  namespaced: true,
  actions,
  getters,
  mutations,
  state,
};
