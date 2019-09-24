import { filter as counter } from './counter';

const filters = {
  counter,
};

export default mutation => {
  const type = mutation.type.split('/');
  if (type.length > 0 && filters[type[0]]) return filters[type[0]]();
  return true;
};
