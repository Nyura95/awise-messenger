import { SET_USERS } from './mutation-types';
import { fetch } from '../../../plugings/request';

export const getAllUsers = ({ commit }, id) => {
  fetch(
    '/api/v1/users',
    'get',
    {},
    {
      Authorization: 'token1',
    }
  ).then(result => {
    commit(SET_USERS, result.Data);
  });
};
export default { getAllUsers };
