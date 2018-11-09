<template>
    <div class="ns-create-account nkn-card-shadow nkn-after-clear">
        <div class="nkn-sign-in-panel">
            <label class="nkn-page-title-label">{{$t("nsSignIn.titleLabel")}}</label>
            <h1 class="nkn-page-title text-main-blue">{{$t("nsSignIn.title")}}</h1>
            <ns-input-item v-for="(inputItem, idx) in inputs" :key="idx" :config="inputItem" />
            <div class="nkn-sign-in">
                <button class="nkn-normal-btn sign-in-button" type="button" @click="signIn">{{$t("nsSignIn.signInbtn")}}</button>
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
  import NsInputItem from "../components/nsInputItem"
  import NsLoading from "../components/nsLoading"
  import NSLocalStorage from "../js/nsLocalStorage"
  import Crypto from "../js/crypto/algorithm"
  import Is from "is_js"
  import Mathjs from "mathjs"
  import Http from "../js/network/nsHttp"
  import {inputIdPrefix} from "../js/nsConst"
  import {nsNamespace} from "../js/nsNamespace"
  import {loadPage} from "../js/nsLoading"

  export default {
    components: {
      NsLoading,
      NsInputItem
    },
    name: "ns-sign-in",
    mounted() {
      if(NSLocalStorage.checkLogin()) {
        this.$nextTick(function () {
          this.$router.push({name: nsNamespace.MAIN})
        })
      } else {
        loadPage.call(this)
      }
    },
    methods: {
      array2HexString(bytes) {
        return Array.from(bytes, function(byte) {
          return ('0' + (byte & 0xFF).toString(16)).slice(-2)
        }).join('')
      },

      getAccountData() {
        let $this = $(this.$el)

        return {
          account: $this.find('.' + inputIdPrefix() + 'account').val(),
          password: $this.find('.' + inputIdPrefix() + 'password').val(),
        }
      },

      accountDataCollect() {
        let verifyFailed = false
        let accountInfo = this.getAccountData()
        this.inputs.account.errorInfo = ''
        this.inputs.password.errorInfo = ''

        if(!Is.alphaNumeric(accountInfo.account)) {
          verifyFailed = true
          this.inputs.account.errorInfo = this.$t('nsInput.account.errorInfo')
        }

        if(accountInfo.password.length < 8) {
          verifyFailed = true
          this.inputs.password.errorInfo = this.$t('nsInput.password.errorInfo')
        }

        return verifyFailed ? null : accountInfo
      },

      signIn() {
        let accountInfo = this.accountDataCollect()
        if(null === accountInfo) {
          return
        }

        let loginData = {
          nonce : this.array2HexString(Mathjs.random([16], 0, 255)),
          randomKey : this.array2HexString(Mathjs.random([32], 0, 255)),
        }

        let reqKey = accountInfo.account + Crypto.HmacSHA256(accountInfo.password)

        Http.login(
          this,
          loginData,
          reqKey,
          function (data) {
            let key = loginData.nonce + loginData.randomKey
            let responseData = Crypto.AESDec(data.Data, key)

            responseData = JSON.parse(responseData)

            if(responseData.Nonce !== loginData.nonce) {
              return
            }

            let walletJsonObj = JSON.parse(responseData.Wallet)

            let accountKey = Crypto.HmacSHA256(accountInfo.password)
            let requestKey = Crypto.HmacSHA256(responseData.SN + accountKey)
            let walletKey = Crypto.HmacSHA256(responseData.SN + accountKey + accountInfo.account)

            NSLocalStorage.setLogin(accountInfo.account,
              responseData.Wallet,
              walletJsonObj.Address,
              {
                account: accountKey,
                req: requestKey,
                wallet: walletKey,
              })

            this.$router.push({name: nsNamespace.MAIN})
          },
          function () {
            alert(this.$t('nsSignIn.nsInput.loginfail'))
          }
        )
      }
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
        }
      }
    }
  }
</script>

<style scoped>
    .ns-create-account {
        padding: 60px 50px;
        background-color: #ffffff;
    }

    .ns-create-account > div {
        float: left;
        width: 50%;
        height: 875px;
    }

    .nkn-sign-in {
        width: 100%;
        padding-right: 94px;
        margin-top: 40px;
    }

    .sign-in-button {
        margin-top: 50px;
        width: 100%;
    }
</style>
