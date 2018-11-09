<template>
    <div class="nkn-setup-page nkn-card-shadow nkn-after-clear">
        <div class="nkn-create-account-panel">
            <label class="nkn-page-title-label">{{$t("nsGenWallet.titleLabel")}}</label>
            <h1 class="nkn-page-title text-main-blue">{{$t("nsGenWallet.title")}}</h1>
            <ns-input-item v-for="(inputItem, idx) in inputs" :key="idx" :config="inputItem" />
            <div class="setup-button nkn-after-clear nkn-wallet-setup-button-panel">
                <a class="nkn-link-load-wallet" @click="toLoadWallet">{{$t("nsGenWallet.link")}}</a>
                <button class="nkn-normal-btn" type="button" @click="nextStep">{{$t("nsGenWallet.nextStepbtn")}}</button>
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
  import Http from "../../js/network/nsHttp"
  import Crypto from "../../js/crypto/algorithm"
  import NSLocalStorage from "../../js/nsLocalStorage"
  import {nsNamespace} from "../../js/nsNamespace";
  import NsLoading from "../../components/nsLoading";
  import {inputIdPrefix} from "../../js/nsConst"
  import {loadPage} from "../../js/nsLoading";

  export default {
    components: {
      NsLoading,
      NsInputItem
    },
    name: "ns-gen-wallet",
    mounted() {
      loadPage.call(this)
    },
    data: function () {
      return {
        inputs: {
          password: {
            inputId: inputIdPrefix() + "password",
            title: 'Wallet password',
            placeholder: '8-20 characters',
            hasAppend: true,
            inputType: 'password',
            maxSize: 20,
            errorInfo: '',
          },

          rePassword: {
            inputId: inputIdPrefix() + "rePassword",
            title: 'Confirm wallet password',
            placeholder: '8-20 characters',
            hasAppend: true,
            inputType: 'password',
            maxSize: 20,
            errorInfo: '',
          },
        }
      }
    },

    methods: {
      getWalletPassword() {
        let $this = $(this.$el)

        return {
          password: $this.find('.' + inputIdPrefix() + 'password').val(),
          rePassword: $this.find('.' + inputIdPrefix() + 'rePassword').val(),
        }
      },

      walletDataGen(accountInfo) {
        let verifyFailed = false
        let walletPassword = this.getWalletPassword()
        this.inputs.password.errorInfo = ''
        this.inputs.rePassword.errorInfo = ''

        if(walletPassword.password.length < 8) {
          verifyFailed = true
          this.inputs.password.errorInfo = 'Please input 8-20 characters'
        }

        if(walletPassword.password !== walletPassword.rePassword) {
          verifyFailed = true
          this.inputs.password.errorInfo = ''
          this.inputs.rePassword.errorInfo = 'The password for the two input is inconsistent!'
        }

        if(verifyFailed) {
          return null
        }

        let wallet = nknWallet.newWallet(walletPassword.password)
        if(!wallet) {
          return null
        }

        let key = Crypto.AESEnc(walletPassword.password, accountInfo.walletKey, false)

        return {
          reqData: {
            Wallet: wallet.toJSON(),
            Key: key,
          },

          wallet: wallet
        }
      },

      toLoadWallet() {
        this.$router.push({name: nsNamespace.SETUP.LOAD_WALLET})
      },

      nextStep() {
        let accountInfo = NSLocalStorage.getAccount()
        let walletData = this.walletDataGen(accountInfo)

        Http.setWallet(this, walletData.reqData, accountInfo.requestKey, function () {
          NSLocalStorage.setWallet(walletData.reqData.Wallet, walletData.wallet.address)
          this.$store.commit(nsNamespace.GLOBAL + "/updateWallet", walletData.wallet)
          this.$router.push({name: nsNamespace.SETUP.SHOW_WALLET})
        }, function (err) {
          console.error(err)
        })
      }
    }
  }
</script>

<style scoped>
    .nkn-wallet-setup-button-panel {
        text-align: right;
    }

    .nkn-wallet-setup-button-panel > a.nkn-link-load-wallet {
        display: inline-block;
        line-height: 110px;
        margin-top: 25px;
        cursor: pointer;
        color: #498ae5;
    }

    .nkn-wallet-setup-button-panel > a.nkn-link-load-wallet:hover {
        text-decoration: underline;
    }
</style>