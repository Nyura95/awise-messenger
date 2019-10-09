import { filter as counter } from './counter';
import { filter as conversation } from './conversation';
import { filter as users } from './users';

const filters = {
  counter,
  conversation,
  users,
};

export default mutation => {
  const type = mutation.type.split('/');
  if (type.length > 0 && filters[type[0]]) return filters[type[0]]();
  return true;
};
