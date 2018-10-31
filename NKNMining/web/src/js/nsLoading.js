import {nsNamespace} from "./nsNamespace"
import {serverStatus} from "./nsConst"
import NSLocalStorage from "./nsLocalStorage"
import Http from "./network/nsHttp"

let goToPage = function(currentPage, targetPage, $router) {
  if(currentPage !== targetPage) {
    $router.push({name: targetPage})
  }
}

let loadPage = function () {
  let currentPage = this.$route.name
  let $router = this.$router

  Http.getStatus(this, function (data) {
    this.$store.commit(nsNamespace.GLOBAL + "/updatePageLoaded", true)

    switch(data.Data.shellStatus) {
      case serverStatus.NS_STATUS_CTEATE_ACCOUNT():
        NSLocalStorage.clear()
        goToPage(currentPage, nsNamespace.SETUP.CREATE_ACCOUNT, $router)
        break

      case serverStatus.NS_STATUS_GEN_WALLET():
        if(!NSLocalStorage.verifyGenWalletState()) {
          goToPage(currentPage, nsNamespace.SIGN_IN, $router)
        } else {
          switch (currentPage) {
            case nsNamespace.SETUP.GEN_WALLET:
              goToPage(currentPage, nsNamespace.SETUP.GEN_WALLET, $router)
              break

            case nsNamespace.SETUP.LOAD_WALLET:
              goToPage(currentPage, nsNamespace.SETUP.LOAD_WALLET, $router)
              break

            default:
              goToPage(currentPage, nsNamespace.SETUP.GEN_WALLET, $router)
              break
          }
        }
        break

      case serverStatus.NS_STATUS_NODE_RUNNING():
        if(!NSLocalStorage.checkLogin() || !NSLocalStorage.verifyData()) {
          goToPage(currentPage, nsNamespace.SIGN_IN, $router)
        } else {
          this.$store.commit(nsNamespace.GLOBAL + "/updateNodeRunning", true)
          goToPage(currentPage, nsNamespace.MAIN, $router)
        }
        break

      case serverStatus.NS_STATUS_NODE_EXITED():
        if(!NSLocalStorage.checkLogin() || !NSLocalStorage.verifyData()) {
          goToPage(currentPage, nsNamespace.SIGN_IN, $router)
        } else {
          this.$store.commit(nsNamespace.GLOBAL + "/updateNodeRunning", false)
          goToPage(currentPage, nsNamespace.MAIN, $router)
        }
        break

      case serverStatus.NS_STATUS_UPDATE_BIN():
        if(!NSLocalStorage.checkLogin() || !NSLocalStorage.verifyData()) {
          goToPage(currentPage, nsNamespace.SIGN_IN, $router)
        } else {
          this.$store.commit(nsNamespace.GLOBAL + "/updateNodeRunning", false)
          goToPage(currentPage, nsNamespace.MAIN, $router)
        }
        break

      default:
        alert("invalid server!")
        break
    }
  }, function (e) {
    console.log(e)
    alert("connect to server failed!")
  })
}

export {
  loadPage
}