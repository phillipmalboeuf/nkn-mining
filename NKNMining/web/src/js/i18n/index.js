import Vue from 'vue'
import VueI18n from 'vue-i18n'
import NSLocalStorage from "../nsLocalStorage"
import Axios from 'axios'

Vue.use(VueI18n)

const i18n = new VueI18n({
  locale: NSLocalStorage.getLanguage() || 'en',
  fallbackLocale: 'en',
  messages: {
    'en': require('./en'),
    'zh': require('./zh')
  }
})

export default i18n
