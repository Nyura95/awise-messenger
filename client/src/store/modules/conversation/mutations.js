import { SET_CONVERSATION } from './mutation-types';

export default {
  [SET_CONVERSATION](state, payload) {
    state.conversation = payload.Conversation;
    state.messages = payload.Messages;
    state.target = payload.Target;
    state.user = payload.User;
  },
};
