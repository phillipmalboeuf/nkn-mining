<template>
    <div class="nkn-setup-page nkn-card-shadow nkn-after-clear">
        <div class="nkn-create-account-panel">
            <label class="nkn-page-title-label">{{$t("nsLoadWallet.titleLabel")}}</label>
            <h1 class="nkn-page-title text-main-blue">{{$t("nsLoadWallet.title")}}</h1>
            <ns-input-item v-for="(inputItem, idx) in inputs" :key="idx" :config="inputItem" />
            <div class="nkn-wallet-select-button">{{$t("nsLoadWallet.upload")}}</div>
            <input class="nkn-wallet-to-load-file-input" type="file" accept=".dat"
                   @change="walletSelected"/>

            <div class="setup-button nkn-after-clear nkn-wallet-setup-button-panel">
                <a class="nkn-link-gen-wallet" @click="toGenWallet">{{$t("nsLoadWallet.link")}}</a>
                <button class="nkn-normal-btn" type="button" @click="nextStep">{{$t("nsLoadWallet.nextStepbtn")}}</button>
            </div>
        </div>
        <div class="nkn-setup-page-wallpaper">
            <img class="nkn-wall-background" :src="'./static/img/wallpaper.png'"/>
            <img class="nkn-wall-pad" :src="'./static/img/wallpad.png'"/>
        </div>

        <ns-loading v-if="!this.$store.state.global.pageLoaded || this.blockUI"/>
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
    name: "ns-load-wallet",
    mounted() {
      loadPage.call(this)

      let $this = $(this.$el)
      let $fileInput = $this.find('.nkn-wallet-to-load-file-input')
      let $fileInputButton = $this.find('.nkn-wallet-select-button')
      let $nameOfFile = $this.find('.' + inputIdPrefix() + "wallet-to-load")
      let $walletInputPanel = $nameOfFile.parent()

      $walletInputPanel.css({
        position: "relative",
      })

      $fileInputButton.appendTo($walletInputPanel).css({
        height: $nameOfFile.outerHeight(),
        lineHeight: $nameOfFile.outerHeight() + "px",
      })

      $fileInput.appendTo($walletInputPanel).css({
        width: $nameOfFile.outerWidth(),
        height: $nameOfFile.outerHeight()
      })
    },
    data: function () {
      return {
        walletJson: '',
        wallet: null,
        blockUI: false,
        inputs: {
          wallet: {
            inputId: inputIdPrefix() + "wallet-to-load",
            title: '&nbsp;',
            placeholder: 'Wallet file to load',
            hasAppend: false,
            inputType: 'text',
            maxSize: 20,
            errorInfo: '',
          },

          password: {
            inputId: inputIdPrefix() + "password",
            title: 'Wallet password',
            placeholder: 'password ot the wallet to load',
            hasAppend: false,
            inputType: 'password',
            maxSize: 20,
            errorInfo: '',
          },
        }
      }
    },

    methods: {
      walletSelected() {
        let $this = $(this.$el)

        this.inputs.password.errorInfo = ''

        let $fileInput = $this.find('.nkn-wallet-to-load-file-input')
        let $nameOfFile = $this.find('.' + inputIdPrefix() + "wallet-to-load")
        let file = $fileInput.get(0).files[0]

        if(!file) {
          return
        }
        $nameOfFile.val(file.name)
      },

      toGenWallet() {
        this.$router.push({name: nsNamespace.SETUP.GEN_WALLET})
      },

      getWalletPassword() {
        let $this = $(this.$el)

        return {
          password: $this.find('.' + inputIdPrefix() + 'password').val(),
        }
      },

      walletDataGen(accountInfo, setWallet) {
        let $this = $(this.$el)

        let walletPassword = this.getWalletPassword()
        this.inputs.password.errorInfo = ''

        let $fileInput = $this.find('.nkn-wallet-to-load-file-input')
        let walletFile = $fileInput.get(0).files[0]

        if(!walletFile) {
          this.inputs.password.errorInfo = 'please select a wallet file'
          return
        }

        let _this = this
        let reader = new FileReader();

        reader.onload = function () {
          _this.walletJson = reader.result
          _this.wallet = nknWallet.loadJsonWallet(_this.walletJson, walletPassword.password)


          console.log(_this.walletJson)
          console.log(_this.wallet)
          if (_this.wallet instanceof nknWallet.nknWalletError) {
            _this.inputs.password.errorInfo = 'Please check if the password matches the wallet.'
            return
          }

          let key = Crypto.AESEnc(walletPassword.password, accountInfo.walletKey, false)

          setWallet({
            reqData: {
              Wallet: _this.wallet.toJSON(),
              Key: key,
            },

            wallet: _this.wallet
          })
        }

        reader.readAsText(walletFile)
      },

      nextStep() {
        let accountInfo = NSLocalStorage.getAccount()
        let _this = this
        this.walletDataGen(accountInfo, function (walletData) {
          _this.blockUI = true
          Http.setWallet(this, walletData.reqData, accountInfo.requestKey, function () {
            NSLocalStorage.setWallet(walletData.reqData.Wallet, walletData.wallet.address)
            _this.$store.commit(nsNamespace.GLOBAL + "/updateWallet", walletData.wallet)
            _this.$router.push({name: nsNamespace.SETUP.SHOW_WALLET})
          }, function (err) {
            _this.blockUI = false
            _this.inputs.password.errorInfo = 'Network error, please try again later.'
            console.error(err)
          })
        })
      }
    }
  }
</script>

<style scoped>
    .nkn-wallet-to-load-file-input {
        opacity: 0;
        position: absolute;
        left: 0;
        top: 0;
        cursor: pointer;
        font-size: 0;
    }

    .nkn-wallet-select-button {
        position: absolute;
        text-align: center;
        color: #ffffff;
        font-size: 14px;
        padding: 0 10px;
        background-color: #243a80;
        right: 0;
        top: 0;
        cursor: pointer;
    }

    .nkn-wallet-setup-button-panel {
        text-align: right;
    }

    .nkn-wallet-setup-button-panel > a.nkn-link-gen-wallet {
        display: inline-block;
        line-height: 110px;
        margin-top: 25px;
        cursor: pointer;
        color: #498ae5;
    }

    .nkn-wallet-setup-button-panel > a.nkn-link-gen-wallet:hover {
        text-decoration: underline;
    }
</style>
