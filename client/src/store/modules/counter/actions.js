import { INCREMENT, DECREMENT } from './mutation-types';

export const decrement = ({ commit }, desc) => {
  commit(DECREMENT, { counter: desc });
};

export const increment = ({ commit }, add) => {
  commit(INCREMENT, { counter: add });
};

export const asyncIncrement = ({ dispatch }, add) => {
  setTimeout(() => {
    dispatch('increment', add);
  }, 500);
};

export default {
  decrement,
  increment,
  asyncIncrement,
};
