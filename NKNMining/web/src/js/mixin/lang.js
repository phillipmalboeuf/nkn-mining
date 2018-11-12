// import {inputIdPrefix} from "../nsConst"
// import i18n from "../i18n/index"
// class NsInputItem{
//   constructor (inputId, title, placeholder, hasAppend, inputType, maxSize, errorInfo) {
//     this.inputId = inputIdPrefix() + inputId
//     this.title = title
//     this.placeholder = placeholder
//     this.hasAppend = hasAppend
//     this.inputType = inputType
//     this.maxSize = maxSize
//     this.errorInfo = errorInfo
//   }
// }

// let account = () => {
//   return new NsInputItem(
//     "account",
//     this.$t('nsInput.account.title'),
//     this.$t('nsInput.account.placeholder'),
//     false,
//     'text',
//     20,
//     '')
// }

// let password = () =>{
//   return new NsInputItem(
//     "password",
//     this.$t('nsInput.password.title'),
//     this.$t('nsInput.password.placeholder'),
//     true,
//     'password',
//     20,
//     ''
//   )
// }

// let rePassword = () =>{
//   return new NsInputItem(
//     "rePassword",
//     this.$t('nsInput.rePassword.title'),
//     this.$t('nsInput.rePassword.placeholder'),
//     true,
//     'password',
//     20,
//     ''
//   )
// }

// let serialNum = () => {
//   return new NsInputItem(
//     "sn",
//     this.$t('nsInput.sn.title'),
//     this.$t('nsInput.sn.placeholder'),
//     false,
//     'text',
//     ''
//   )
// }

// let wallet = () => {
//   return new NsInputItem(
//     "wallet-to-load",
//     this.$t('nsInput.wallet.title'),
//     this.$t('nsInput.wallet.placeholder'),
//     false,
//     'text',
//     20,
//     ''
//   )
// }

// let walletPassword = () => {
//   return new NsInputItem(
//     "password",
//     this.$t('nsInput.walletPassword.title'),
//     this.$t('nsInput.walletPassword.placeholder'),
//     false,
//     'password',
//     20,
//     '',
//   )
// }

export default {
  computed: {
    lang() {
      return this.$store.state.global.language
    }
  }
}