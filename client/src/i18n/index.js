import Vue from 'vue';
import VueI18n from 'vue-i18n';

import messages from './locale';
import { config } from '../config';

Vue.use(VueI18n);

// Create VueI18n instance with options
export default new VueI18n({
  locale: config.i18n.defaultLang,
  messages,
});
