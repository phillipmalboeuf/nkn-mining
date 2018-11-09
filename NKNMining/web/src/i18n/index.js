import Vue from 'vue'
import VueI18n from 'vue-i18n'
import en from './en'
import zh from './zh'
import NSLocalStorage from "../js/nsLocalStorage"

Vue.use(VueI18n)

export default new VueI18n({
  locale: NSLocalStorage.getLanguage() || 'en',
  fallbackLocale: 'en',
  messages: {
    en, zh
  }
})

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