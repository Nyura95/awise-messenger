import { INCREMENT, DECREMENT } from './mutation-types';

export default {
  [INCREMENT](state, payload) {
    state.counter += payload.counter;
  },
  [DECREMENT](state, payload) {
    state.counter -= payload.counter;
  },
};
