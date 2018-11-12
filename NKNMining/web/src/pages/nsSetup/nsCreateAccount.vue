<template>
    <div class="nkn-setup-page nkn-card-shadow nkn-after-clear">
        <div class="nkn-create-account-panel">
            <label class="nkn-page-title-label">{{$t("nsCreateAccount.titleLabel")}}</label>
            <h1 class="nkn-page-title text-main-blue">{{$t("nsCreateAccount.title")}}</h1>
            <ns-input-item v-for="(inputItem, idx) in inputs" :key="idx" :config="inputItem" />
            <div class="setup-button nkn-after-clear">
                <button class="nkn-normal-btn" type="button" @click="nextStep">{{$t("nsCreateAccount.nextStepbtn")}}</button>
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
  import NsLoading from "../../components/nsLoading"
  import LangMix from "../../js/mixin/lang.js"
  import {loadPage} from "../../js/nsLoading"

  export default {
    components: {
      NsLoading,
      NsInputItem
    },
    mixins:[LangMix],
    name: "ns-setup-create-accounts",
    mounted() {
      loadPage.call(this)
    },
    data: function () {
      return {
        inputs: {
          account: {
            inputId: inputIdPrefix() + "account",
            title: this.$t('nsInput.account.title'),
            placeholder: this.$t('nsInput.account.placeholder'),
            hasAppend: false,
            inputType: 'text',
            maxSize: 20,
            errorInfo: '',
          },

          password: {
            inputId: inputIdPrefix() + "password",
            title: this.$t('nsInput.password.title'),
            placeholder: this.$t('nsInput.password.placeholder'),
            hasAppend: true,
            inputType: 'password',
            maxSize: 20,
            errorInfo: '',
          },

          rePassword: {
            inputId: inputIdPrefix() + "rePassword",
            title: this.$t('nsInput.rePassword.title'),
            placeholder: this.$t('nsInput.rePassword.placeholder'),
            hasAppend: true,
            inputType: 'password',
            maxSize: 20,
            errorInfo: '',
          },

          sn: {
            inputId: inputIdPrefix() + "sn",
            title: this.$t('nsInput.sn.title'),
            placeholder: this.$t('nsInput.sn.placeholder'),
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
          this.inputs.account.errorInfo = this.$t('nsInput.account.errorInfo')
        }

        if(accountInfo.sn.length !== 40) {
          verifyFailed = true
          this.inputs.sn.errorInfo = this.$t('nsInput.sn.errorInfo')
        }

        if(accountInfo.password.length < 8) {
          verifyFailed = true
          this.inputs.password.errorInfo = this.$t('nsInput.password.errorInfo')
        }

        if(accountInfo.password !== accountInfo.rePassword) {
          verifyFailed = true
          this.inputs.password.errorInfo = ''
          this.inputs.rePassword.errorInfo = this.$t('nsInput.rePassword.errorInfo')
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
    },
    watch: {
      lang () {
        this.inputs = {
           account: {
            inputId: inputIdPrefix() + "account",
            title: this.$t('nsInput.account.title'),
            placeholder: this.$t('nsInput.account.placeholder'),
            hasAppend: false,
            inputType: 'text',
            maxSize: 20,
            errorInfo: '',
          },

          password: {
            inputId: inputIdPrefix() + "password",
            title: this.$t('nsInput.password.title'),
            placeholder: this.$t('nsInput.password.placeholder'),
            hasAppend: true,
            inputType: 'password',
            maxSize: 20,
            errorInfo: '',
          },

          rePassword: {
            inputId: inputIdPrefix() + "rePassword",
            title: this.$t('nsInput.rePassword.title'),
            placeholder: this.$t('nsInput.rePassword.placeholder'),
            hasAppend: true,
            inputType: 'password',
            maxSize: 20,
            errorInfo: '',
          },

          sn: {
            inputId: inputIdPrefix() + "sn",
            title: this.$t('nsInput.sn.title'),
            placeholder: this.$t('nsInput.sn.placeholder'),
            hasAppend: false,
            inputType: 'text',
            errorInfo: '',
          },
        }
      }
    }
  }
</script>

<style scoped>

</style>
