import { SET_CONVERSATION, ADD_MESSAGE } from './mutation-types';
import { fetch } from '../../../plugings/request';

export const getConversation = ({ commit }, id) => {
  fetch(
    '/api/v1/conversation/target/' + id,
    'get',
    {},
    {
      Authorization: 'token1',
    }
  ).then(result => {
    commit(SET_CONVERSATION, result.Data);
  });
};

export const addMessage = ({ commit }, message) => {
  console.log(message);
  commit(ADD_MESSAGE, message);
};

export default {
  getConversation,
  addMessage,
};
