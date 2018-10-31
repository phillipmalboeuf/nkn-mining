<template>
    <div class="nkn-setup-page nkn-card-shadow nkn-after-clear">
        <div class="nkn-create-account-panel">
            <label class="nkn-page-title-label">setup - step 1</label>
            <h1 class="nkn-page-title text-main-blue">Create account</h1>
            <ns-input-item v-for="(inputItem, idx) in inputs" :key="idx" :config="inputItem" />
            <div class="setup-button nkn-after-clear">
                <button class="nkn-normal-btn" type="button" @click="nextStep">NEXT</button>
            </div>
        </div>
        <div class="nkn-setup-page-wallpaper">
            <img class="nkn-wall-background" :src="'./static/img/wallpaper.png'"/>
            <img class="nkn-wall-pad" :src="'./static/img/wallpad.png'"/>
        </div>
        <ns-loading v-if="!this.$store.state.global.pageLoaded"/>
    </div>
</template>

<script>
  import NsInputItem from "../../components/nsInputItem"
  import Is from "is_js"
  import Http from "../../js/network/nsHttp"
  import Crypto from "../../js/crypto/algorithm"
  import {inputIdPrefix} from "../../js/nsConst"
  import NSLocalStroage from "../../js/nsLocalStorage"
  import NsLoading from "../../components/nsLoading";
  import {loadPage} from "../../js/nsLoading";

  export default {
    components: {
      NsLoading,
      NsInputItem
    },
    name: "ns-setup-create-accounts",
    mounted() {
      loadPage.call(this)
    },
    data: function () {
      return {
        inputs: {
          account: {
            inputId: inputIdPrefix() + "account",
            title: 'Account',
            placeholder: 'Alphanumeric only',
            hasAppend: false,
            inputType: 'text',
            maxSize: 20,
            errorInfo: '',
          },

          password: {
            inputId: inputIdPrefix() + "password",
            title: 'Account password',
            placeholder: '8-20 characters',
            hasAppend: true,
            inputType: 'password',
            maxSize: 20,
            errorInfo: '',
          },

          rePassword: {
            inputId: inputIdPrefix() + "rePassword",
            title: 'Confirm account password',
            placeholder: '8-20 characters',
            hasAppend: true,
            inputType: 'password',
            maxSize: 20,
            errorInfo: '',
          },

          sn: {
            inputId: inputIdPrefix() + "sn",
            title: 'System initialization serial number',
            placeholder: '40 characters',
            hasAppend: false,
            inputType: 'text',
            errorInfo: '',
          },
        }
      }
    },

    methods: {
      getAccountData() {
        let $this = $(this.$el)

        return {
          account: $this.find('.' + inputIdPrefix() + 'account').val(),
          password: $this.find('.' + inputIdPrefix() + 'password').val(),
          rePassword: $this.find('.' + inputIdPrefix() + 'rePassword').val(),
          sn: $this.find('.' + inputIdPrefix() + 'sn').val(),
        }
      },
      accountDataCollect() {
        let verifyFailed = false
        let accountInfo = this.getAccountData()
        this.inputs.account.errorInfo = ''
        this.inputs.password.errorInfo = ''
        this.inputs.rePassword.errorInfo = ''
        this.inputs.sn.errorInfo = ''

        if(!Is.alphaNumeric(accountInfo.account)) {
          verifyFailed = true
          this.inputs.account.errorInfo = 'Alphanumeric only!'
        }

        if(accountInfo.sn.length !== 40) {
          verifyFailed = true
          this.inputs.sn.errorInfo = 'Invalid serial number'
        }

        if(accountInfo.password.length < 8) {
          verifyFailed = true
          this.inputs.password.errorInfo = 'Please input 8-20 characters'
        }

        if(accountInfo.password !== accountInfo.rePassword) {
          verifyFailed = true
          this.inputs.password.errorInfo = ''
          this.inputs.rePassword.errorInfo = 'The password for the two input is inconsistent!'
        }

        return verifyFailed ? null : accountInfo
      },
      nextStep() {
        let accountInfo = this.accountDataCollect()
        if(null === accountInfo) {
            return
        }

        let account = accountInfo.account
        let key = Crypto.HmacSHA256(accountInfo.password)
        Http.setAccount(this, {
          Account: account,
          Key: key,
        }, accountInfo.sn, function () {
          let requestKey = Crypto.HmacSHA256(accountInfo.sn + key)
          let walletKey = Crypto.HmacSHA256(accountInfo.sn + key + account)

          NSLocalStroage.setAccount(account, key, requestKey, walletKey)
          this.$router.push("/gen_wallet")
        }, function (err) {
          alert(err.response.data)
        })
      }
    }
  }
</script>

<style scoped>

</style>
