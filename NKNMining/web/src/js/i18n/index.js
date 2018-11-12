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

// const loadedLanguages = ['en']

// function setI18nLanguage (lang) {
//   i18n.locale = lang
//   Axios.defaults.headers.common['Accept-Language'] = lang
//   document.querySelector('html').setAttribute('lang', lang)
//   return lang
// }

// export function loadLanguageAsync (lang) {
//   if (i18n.locale !== lang) {
//     if (!loadedLanguages.includes(lang)) {
//       return import(/* webpackChunkName: "lang-[request]" */ `@/lang/${lang}`).then(msgs => {
//         i18n.setLocaleMessage(lang, msgs.default)
//         loadedLanguages.push(lang)
//         return setI18nLanguage(lang)
//       })
//     } 
//     return Promise.resolve(setI18nLanguage(lang))
//   }
//   return Promise.resolve(lang)
// }