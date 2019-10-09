import { SET_USERS } from './mutation-types';
import { fetch } from '../../../plugings/request';

export const getAllUsers = ({ commit }, id) => {
  fetch(
    '/api/v1/users',
    'get',
    {},
    {
      Authorization:
        'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjU4OCwiZm5hbWUiOiJ4eEB4eC54eCIsImxuYW1lIjoiR09VUkVUVEUiLCJlbWFpbCI6Inh4QHh4Lnh4IiwiaWF0IjoxNTY2MjI1MDMzLCJleHAiOjE1NjYyMjUwOTN9.P7y02ThoG0z-Ytpj1wZKiJ3-7M4XhJilSJ_bnk50YZI',
    }
  ).then(result => {
    commit(SET_USERS, result.Data);
  });
};
export default { getAllUsers };
