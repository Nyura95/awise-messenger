import { SET_CONVERSATION, ADD_MESSAGE } from './mutation-types';

export default {
  [SET_CONVERSATION](state, payload) {
    state.conversation = payload.Conversation;
    state.messages = payload.Messages;
    state.target = payload.Target;
    state.user = payload.User;
  },
  [ADD_MESSAGE](state, payload) {
    state.messages = [...state.messages, payload];
  },
};
