import { SET_CONVERSATION, ADD_MESSAGE } from './mutation-types';
import { fetch } from '../../../plugings/request';

export const getConversation = ({ commit }, id) => {
  fetch(
    '/api/v1/conversation/target/' + id,
    'get',
    {},
    {
      Authorization:
        'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjU4OCwiZm5hbWUiOiJ4eEB4eC54eCIsImxuYW1lIjoiR09VUkVUVEUiLCJlbWFpbCI6Inh4QHh4Lnh4IiwiaWF0IjoxNTY2MjI1MDMzLCJleHAiOjE1NjYyMjUwOTN9.P7y02ThoG0z-Ytpj1wZKiJ3-7M4XhJilSJ_bnk50YZI',
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
