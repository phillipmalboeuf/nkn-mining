<template>
  <div id="app">
    <ns-nav/>
    <div id="nkn-shell-page">
      <router-view/>
    </div>
    <ns-footer/>
  </div>
</template>

<script>
import NsNav from "./components/nsNav"
import NsFooter from "./components/nsFooter"
import Http from "./js/network/nsHttp"
import {serverStatus} from "./js/nsConst"
import {nsNamespace} from "./js/nsNamespace";
import NSLocalStorage from "./js/nsLocalStorage"
// import nknWallet from "nkn-wallet"

export default {
  components: {
    NsFooter,
    NsNav
  },
  name: 'App',
  created() {
    nknWallet.configure({
      assetId: '4945ca009174097e6614d306b66e1f9cb1fce586cb857729be9e1c5cc04c9c02',
      rpcAddr: nknShellWebConfig.nknNode,
    })
  }
}
</script>

<style>
  html {
    min-width: 1440px;
  }

  body {
    width: 100%;
    margin: 0 auto;
    background: #f7f7f7;
  }

  #nkn-shell-page {
    width: 1440px;
    padding: 0 60px 20px;
    margin: 0 auto;
  }

  #nkn-shell-page > .nkn-page-cointainer {
    width: 100%;
  }

  .nkn-card-shadow {
    box-shadow: rgba(0, 0, 0, 0.12) 0 2px 6px, rgba(0, 0, 0, 0.24) 0 1px 2px;
  }

  .nkn-after-clear::after {
    width: 0;
    height: 0;
    content: ' ';
    visibility: hidden;
    display: block;
    clear: both;
  }
</style>

<style>
    .text-main-blue {
        color: #243a80;
    }

    .background-main-blue {
        background-color: #243a80;
    }

    .text-label-gray {
        color: #808c9b;
    }
</style>

<style>
  .nkn-page-title-label {
    font-size: 20px;
    line-height: 20px;
    color: #8d96b4;
  }

  .nkn-page-title {
    font-size: 48px;
    margin-bottom: 105px;
  }
</style>

<style>
  .nkn-setup-page {
    padding: 60px 50px;
    background-color: #ffffff;
  }

  .nkn-setup-page > div {
    float: left;
    width: 50%;
    height: 875px;
  }

  div.nkn-setup-page-wallpaper {
    position: relative;
  }

  div.nkn-setup-page-wallpaper > img.nkn-wall-background {
    width: 100%;
  }

  div.nkn-setup-page-wallpaper > img.nkn-wall-pad {
    width: 560px;
    position: absolute;
    left: 30px;
    top: 40px;

    animation:pad-ani 3s infinite;

  }

  @keyframes pad-ani{
    0% {
      -webkit-transform: translate(0,0);
      -moz-transform: translate(0,0);
      -ms-transform: translate(0,0);
      -o-transform: translate(0,0);
      transform: translate(0,0);
    }
    50% {
      -webkit-transform: translate(0,-4%);
      -moz-transform: translate(0,-4%);
      -ms-transform: translate(0,-4%);
      -o-transform:translate(0,-4%);
      transform: translate(0,-4%);
    }
    100% {
      -webkit-transform: translate(0, 0);
      -moz-transform: translate(0, 0);
      -ms-transform: translate(0, 0);
      -o-transform: translate(0, 0);
      transform: translate(0, 0);
    }
  }


  .nkn-setup-page div.setup-button {
    width: 100%;
    padding-right: 94px;
    margin-top: 40px;
  }

  .nkn-setup-page .setup-button > button {
    float: right;
    margin-top: 50px;
    margin-left: 40px;
  }
</style>

<style>
    .nkn-normal-btn {
        width: 160px;
        height: 60px;
        border-radius: 4px;
        color: #ffffff;
        cursor: pointer;
        background: rgba(48, 139, 236, 1);
        transition: background-color 0.2s;
    }

    .nkn-normal-btn,
    .nkn-normal-btn:active {
        border: 0 !important;
        outline: 0 none !important;
    }

    .nkn-normal-btn:hover {
        background: rgba(48, 139, 236, 0.8);
        transition: background-color 0.2s;
    }

    .nkn-normal-btn:active {
        background: rgba(48, 139, 236, 1);
        transition: background-color 0.2s;
    }
</style>