import { config } from '../../config';
import fetchival from 'fetchival';

export const fetch = (url, method = 'get', payload = {}, headers = {}) =>
  new Promise((resolve, reject) => {
    fetchival(config.api.basepath + url, {
      headers: {
        ...headers,
      },
    })
      [method](payload)
      .then(res => resolve(res))
      .catch(err => reject(err));
  });
