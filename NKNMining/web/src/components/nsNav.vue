<template>
    <nav class="nk-nav nkn-card-shadow">
        <div class="nkn-nav-bar">
            <div class="nkn-logo" @click="goHome">
                <img src="../assets/img/icon/logo.png"/>
                <span> - V0.0.7</span>
            </div>
            <div class="nav-bar-dropdown">
                <div v-if="showUserIcon" class="nkn-user"
                    data-toggle="dropdown"
                    aria-haspopup="true"
                    aria-expanded="false">
                    <img src="../assets/img/icon/userhead.png"/>
                </div>
                <div class="nkn-nav-menu dropdown-menu">
                    <button class="dropdown-item" type="button"
                            @click="signOut()">{{$t("nsNav.signOut")}}</button>
                    <button class="dropdown-item" type="button"
                            @click="resetNodeShell()">{{$t("nsNav.resetNodeMining")}}</button>
                </div>
            </div>
            <div class="nav-bar-dropdown">
                <div class="nkn-lang-sel"
                data-toggle="dropdown"
                aria-haspopup="true"
                aria-expanded="false">
                    <img src="../assets/img/icon/language.png"/><span class="text-main-blue ">{{$t('nav.langSel')}}</span>
                </div>
                <div class="nkn-nav-lang-sel-menu dropdown-menu">
                    <button class="dropdown-item" type="button"
                            @click="changeLanguage('en')">{{$t('nav.langList.en')}}</button>
                    <button class="dropdown-item" type="button"
                            @click="changeLanguage('zh')">{{$t('nav.langList.zh')}}</button>
                </div>
            </div>
        </div>
    </nav>
</template>

<script>
  import NSLocalStorage from '../js/nsLocalStorage'
  import Http from '../js/network/nsHttp'
  import {nsNamespace} from "../js/nsNamespace";

  export default {
    name: "ns-nav",
    computed: {
      showUserIcon() {
        return nsNamespace.MAIN === this.$route.name
      }
    },
    methods: {
      signOut() {
        NSLocalStorage.logout()
        this.$router.push({name: nsNamespace.SIGN_IN})
      },

      resetNodeShell() {
        let $dlg = $("#reset-confirm-model")
        if(0 === $dlg.length) {
          return
        }

        $dlg.modal('show')
      },

      goHome() {
        this.$router.push("/")
      },

      changeLanguage(lang) {
        if(this.$i18n.locale === lang) return 
        this.$i18n.locale = lang
        this.$store.commit(nsNamespace.GLOBAL+'/updateLanguage',lang)
        NSLocalStorage.setLanguage(lang)
      }
    },
  }
</script>

<style scoped>
    .nk-nav {
        height: 60px;
        width: 100%;
        margin-bottom: 30px;
        background: #ffffff;
        line-height: 60px;
        font-size: 16px;
    }

    .nkn-nav-bar {
        width: 1440px;
        height: 60px;
        padding: 0 60px;
        margin: 0 auto;
    }

    .nkn-logo {
        float: left;
        color: #243a80;
    }

    .nkn-logo > img {
        height: 30px;
        margin-top: -4px;
    }

    .nkn-user, .nkn-lang-sel {
        float: right;
        font-size: 20px;
        cursor: pointer;
    }

    .nkn-user > img, .nkn-lang-sel > img {
        height: 20px;
        padding-right: 10px;
        margin-top: -4px;
    }

    .text-main-blue {
        padding-right: 10px;
    }

    .nkn-nav-lang-sel-menu {
        text-align: center;
    }
</style>

<style scoped>
    .dropdown-menu {
        margin-left: -80px;
    }

    .dropdown-item {
        line-height: 40px;
    }
</style>