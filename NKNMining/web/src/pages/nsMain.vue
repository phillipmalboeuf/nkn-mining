<template>
    <div class="nkn-main-page nkn-after-clear">
        <div class="operation-panel nkn-after-clear">
            <div class="operation-connect-status nkn-card-shadow nkn-after-clear">
                <div class="operation-title">{{$t("nsMain.node.title")}}</div>
                <div v-if="this.$store.state.global.nodeRunning" class="switch-button switch-on" @click="stopMining"><img src="../assets/img/icon/switch-on.png"/><span>{{$t("nsMain.node.on")}}</span></div>
                <div v-else class="switch-button" @click="startMining"><img src="../assets/img/icon/switch-off.png"/><span>{{$t("nsMain.node.off")}}</span></div>

                <div class="node-info">{{nodeStatus}}</div>

                <div class="node-info">{{version}}</div>
                <div class="node-info">{{$t("nsMain.node.NKNBlockHeight")}} {{this.nknHeight}}</div>
                <div class="node-info">{{$t("nsMain.node.myBlockHeight")}} {{this.currentHeight}}</div>
            </div>

            <div class="operation-wallet nkn-card-shadow nkn-after-clear">
                <div class="operation-title">{{$t("nsMain.wallet.title")}}</div>
                <div class="wallet-operation-button" @click="downloadWallet"><button>{{$t("nsMain.wallet.downloadbtn")}}</button></div>
                <div class="wallet-operation-button" @click="showTransferDlg"><button>{{$t("nsMain.wallet.transferbtn")}}</button></div>
                <div class="wallet-info-panel">
                    <div class="wallet-address">
                        <label class="info-title">{{$t("nsMain.wallet.addressLabel")}}</label>
                        <div class="info-text text-main-blue">{{walletAddress}}</div>
                    </div>
                    <div class="wallet-balance">
                        <label class="info-title">{{$t("nsMain.wallet.balanceLabel")}}</label>
                        <div class="info-text text-main-blue">
                            <span>{{balance}}</span><span class="balance-unit">{{$t("nsMain.wallet.balanceUnit")}}</span>
                        </div>
                    </div>
                </div>
            </div>

            <div class="info-neighbor-panel nkn-card-shadow nkn-after-clear">
                <div class="operation-title">{{$t("nsMain.neighbors.title")}}</div>
                <div class="neighbor-list-header">
                    <div class="neighbor-list-row">
                        <div class="neighbor-list-cell id-cell">{{$t("nsMain.neighbors.ID")}}</div>
                        <div class="neighbor-list-cell ip-cell">{{$t("nsMain.neighbors.IP")}}</div>
                        <div class="neighbor-list-cell port-cell">{{$t("nsMain.neighbors.port")}}</div>
                    </div>
                </div>

                <div class="neighbor-list-body">
                    <div class="neighbor-list-row"
                         v-for="(v, i) in neighbor" :key="i">
                        <div class="neighbor-list-cell id-cell">{{i}}</div>
                        <div class="neighbor-list-cell ip-cell">{{v.IpAddr}}</div>
                        <div class="neighbor-list-cell port-cell">{{v.Port}}</div>
                    </div>
                </div>
            </div>
        </div>
        <div class="transfer-list nkn-card-shadow">
            <div class="transfer-list-title">{{$t("nsMain.rewards.title")}}</div>
            <div class="transfer-list-header nkn-after-clear">
                <div class="list-cell cell-tx">{{$t("nsMain.rewards.transaction")}}</div>
                <!--<div class="list-cell cell-type">Type</div>-->
                <div class="list-cell cell-value">{{$t("nsMain.rewards.value")}}</div>
                <div class="list-cell cell-height">{{$t("nsMain.rewards.height")}}</div>
                <div class="list-cell cell-time">{{$t("nsMain.rewards.time")}}</div>
            </div>

            <div class="transfer-list-body nkn-after-clear">
                <div v-for="(v, i) in this.miningRewards"
                     class="transfer-row nkn-after-clear"
                     :class="(1 === (i+1) % 2) ? 'transfer-gray-row':''"
                     :key="i"
                     >
                    <div class="list-cell cell-tx" :title="v.Hash">{{v.Hash}}</div>
                    <!--<div class="list-cell cell-type">Mining reward</div>-->
                    <div class="list-cell cell-value">{{v.Value}}&nbsp;{{$t("nsMain.wallet.balanceUnit")}}</div>
                    <div class="list-cell cell-height">{{v.Height}}</div>
                    <div class="list-cell cell-time">{{v.Timestamp}}</div>
                </div>
            </div>
            <!--<div class="transfer-list-footer"></div>-->
        </div>
        <ns-loading v-if="!this.$store.state.global.pageLoaded"/>

        <!-- nkn transfer -->
        <div class="modal fade" id="nkn-transfer-model" tabindex="-1" role="dialog" aria-hidden="true">
            <div class="modal-dialog" role="document">
                <div class="modal-content">
                    <div class="modal-header">
                        <h5 class="modal-title">{{$t("nsMain.transfer.title")}}</h5>
                    </div>
                    <div class="modal-body">
                        <div class="input-group input-group-lg">
                            <input id="transfer-to-address" type="text" class="form-control" :placeholder="$t('nsMain.transfer.addressPlaceholder')">
                        </div>

                        <div class="input-group input-group-lg">
                            <input id="transfer-to-value" type="number" class="form-control" :placeholder="$t('nsMain.transfer.countPlaceholder')">
                        </div>

                        <div class="input-group input-group-lg">
                            <input id="transfer-to-wallet-password" type="password" class="form-control" :placeholder="$t('nsMain.transfer.passwordPlaceholder')">
                        </div>
                    </div>
                    <div class="modal-footer">
                        <button type="button" class="btn btn-secondary" data-dismiss="modal">{{$t("nsMain.transfer.cancel")}}</button>
                        <button type="button" class="btn btn-danger" @click="doTransfer">{{$t("nsMain.transfer.confirm")}}</button>
                    </div>
                </div>
            </div>
        </div>

        <!-- reset shell node dlg-->
        <div class="modal fade" id="reset-confirm-model" tabindex="-1" role="dialog" aria-hidden="true">
            <div class="modal-dialog" role="document">
                <div class="modal-content">
                    <div class="modal-header">
                        <h5 class="modal-title" id="exampleModalLabel">{{$t("nsMain.resetdlg.title")}}</h5>
                    </div>
                    <div class="modal-body">
                        <div class="reset-shell-warning">
                            {{$t("nsMain.resetdlg.warning")}}
                        </div>

                        <div class="input-group input-group-lg">
                            <input id="reset-shell-confirm-password" type="password" class="form-control" :placeholder="$t('nsMain.resetdlg.passwordPlaceholder')">
                        </div>
                    </div>
                    <div class="modal-footer">
                        <button type="button" class="btn btn-secondary" data-dismiss="modal">{{$t("nsMain.resetdlg.cancel")}}</button>
                        <button type="button" class="btn btn-primary" @click="resetShell">{{$t("nsMain.resetdlg.reset")}}</button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
  import NsLoading from "../components/nsLoading"
  import {loadPage} from "../js/nsLoading"
  import {serverStatus} from "../js/nsConst";
  import {nsNamespace} from "../js/nsNamespace"
  import Http from "../js/network/nsHttp"
  import NSLocalStorage from "../js/nsLocalStorage"
  import FileSaver from 'file-saver'
  import Crypto from '../js/crypto/algorithm'

  function loopStatusQuery(scope) {
    Http.getStatus(scope, function (data) {
      scope.nodeInfo = data.Data.nodeInfo
      scope.neighbor = data.Data.neighbor
      scope.currentHeight = data.Data.blockHeight
      scope.nknHeight = data.Data.nknNetworkHeight

      if(!scope.neighbor) {
        scope.neighbor = []
      } else {
        for(let i=0; i<scope.neighbor.length; i++){
          scope.neighbor[i].IpAddr =
            scope.neighbor[i].IpAddr[12] + '.' +
            scope.neighbor[i].IpAddr[13] + '.' +
            scope.neighbor[i].IpAddr[14] + '.' +
            scope.neighbor[i].IpAddr[15]
        }
      }

      if(data.Data.shellStatus !== serverStatus.NS_STATUS_NODE_RUNNING()) {
        if(data.Data.shellStatus === serverStatus.NS_STATUS_UPDATE_BIN()) {
          this.nodeStatus = this.$t('nsMain.node.nodeStatus.updating')
        } else {
          this.nodeStatus = this.$t('nsMain.node.nodeStatus.stoped')
        }

        this.shellStatus = data.Data.shellStatus
        scope.$store.commit(nsNamespace.GLOBAL + "/updateNodeRunning", false)
        setTimeout(function () {
          loopStatusQuery(scope)
        }, 5000)
        return
      }

      scope.$store.commit(nsNamespace.GLOBAL + "/updateNodeRunning", true)
      switch(data.Data.syncStatus) {
        case "SyncStarted":
          scope.nodeStatus = scope.$t('nsMain.node.nodeStatus.syncing')
          break

        case "SyncFinished":
          scope.nodeStatus = scope.$t('nsMain.node.nodeStatus.persisting')
          break

        case "PersistFinished":
          scope.nodeStatus = scope.$t('nsMain.node.nodeStatus.mining')
          break

        default:
          scope.nodeStatus = scope.$t('nsMain.node.nodeStatus.syncing')
          break
      }

      setTimeout(function () {
        loopStatusQuery(scope)
      }, 5000)
    }, function (e) {
      try{
        scope.nodeStatus = scope.$t('nsMain.node.nodeStatus.syncing')
        setTimeout(function () {
        loopStatusQuery(scope)
      }, 5000)
      }catch(e){}
    })
  }

  function loopWalletQuery(scope, tempWallet) {
    if(!scope.walletAddress) {
      setTimeout(function () {
        loopWalletQuery(scope, tempWallet)
      }, 3000)

      return
    }

    tempWallet.queryAssetBalance(scope.walletAddress)
      .then(function (balance) {
        scope.balance = balance.toString()

        setTimeout(function () {
          loopWalletQuery(scope, tempWallet)
        }, 3000)

      })
      .catch( function () {
        setTimeout(function () {
          loopWalletQuery(scope, tempWallet)
        }, 3000)
      })
  }

  function loopMiningRewardsQuery(scope) {
    Http.getRewardsList(scope, scope.walletAddress, 0xffffffff, function (data) {
      if(data.Data) {
        scope.miningRewards = data.Data
      }

      setTimeout(function () {
        loopMiningRewardsQuery(scope)
      }, 3000)
    }, function (err) {
      console.error(err)
      setTimeout(function () {
        loopMiningRewardsQuery(scope)
      }, 3000)
    })
  }

  export default {
    components: {NsLoading},
    name: "ns-main",
    data: function () {
      return {
        walletAddress: "",
        balance: this.$t('nsMain.wallet.balance.init'),
        accountInfo: NSLocalStorage.getAccount(),
        version: this.$t('nsMain.node.version.init'),
        nodeStatus: this.$t('nsMain.node.nodeStatus.init'),
        miningRewards: [],
        nodeInfo: {},
        neighbor: [],
        currentHeight: 0,
        nknHeight: 0,
        shellStatus: serverStatus.NS_STATUS_NODE_EXITED(),
      }
    },
    mounted() {
      loadPage.apply(this)
      this.walletAddress = NSLocalStorage.getWalletAddress()
      loopWalletQuery(this, nknWallet.newWallet("pwd"))
      loopStatusQuery(this)
      loopMiningRewardsQuery(this)

      Http.getVersion(this, function (data) {
        this.version = data.Data.NodeVersion
      }, function () {
        console.log("get version failed")
      })
    },
    methods: {
      showTransferDlg() {
        $("#nkn-transfer-model").modal('show')
      },

      doTransfer() {
        let address = $("#transfer-to-address").val()
        let value = $("#transfer-to-value").val()
        let password = $("#transfer-to-wallet-password").val()

        let walletJson = NSLocalStorage.getWallet()
        let wallet = nknWallet.loadJsonWallet(walletJson, password)
        if(wallet.msg) {
          $("#nkn-transfer-model").modal('hide')
          alert(wallet.msg)
          return
        }

        $("#nkn-transfer-model").modal('hide')
        this.$store.commit(nsNamespace.GLOBAL + "/updatePageLoaded", false)

        let _this = this
        wallet.transferTo(address, value, password, function () {
          setTimeout(function () {
            _this.$store.commit(nsNamespace.GLOBAL + "/updatePageLoaded", true)
          })
        }, function (err) {
          if(err.msg) {
            // alert(err.msg)
            alert(this.$t('nsMain.transfer.alertInfo.default'))
          } else {
            alert(this.$t('nsMain.transfer.alertInfo.default'))
          }

          _this.$store.commit(nsNamespace.GLOBAL + "/updatePageLoaded", true)
        })
      },

      resetShell() {
        let confirmPassword = $("#reset-shell-confirm-password").val()
        if('' === confirmPassword.trim()) {
          return
        }

        confirmPassword = Crypto.HmacSHA256(confirmPassword)
        let account = NSLocalStorage.getAccount()

        if(confirmPassword !== account.accountKey) {
          $("#reset-confirm-model").modal('hide')
          alert(this.$t('nsMain.resetdlg.alertInfo.wrongPass'))
          return
        }

        Http.resetShell(this, account.requestKey, function () {
          window.location.reload()
        }, function () {
          alert(this.$t('nsMain.resetdlg.alertInfo.success'))
        })
      },
      downloadWallet() {
        let walletData = new Blob([NSLocalStorage.getWallet()], {type: "text/plain;charset=utf-8"});
        FileSaver.saveAs(walletData, "wallet.dat");
      },

      startMining() {
        if(serverStatus.NS_STATUS_UPDATE_BIN() === this.shellStatus) {
          alert(this.$t('nsMain.node.mining.error'))
          return
        }
        this.nodeStatus = this.$t('nsMain.node.nodeStatus.starting')
        this.$store.commit(nsNamespace.GLOBAL + "/updatePageLoaded", false)

        let _this = this
        Http.startMining(this, this.accountInfo.requestKey,
          function () {
            setTimeout(function () {
              _this.$store.commit(nsNamespace.GLOBAL + "/updatePageLoaded", true)
              _this.$store.commit(nsNamespace.GLOBAL + "/updateNodeRunning", true)
            }, 800)

          },
          function () {
            alert(this.$t('nsMain.node.mining.fail'))
            this.$store.commit(nsNamespace.GLOBAL + "/updatePageLoaded", true)
          })
      },

      stopMining() {
        this.nodeStatus = this.$t('nsMain.node.nodeStatus.stopping')
        this.$store.commit(nsNamespace.GLOBAL + "/updatePageLoaded", false)

        let _this = this
        Http.stopMining(this, this.accountInfo.requestKey,
          function () {
            setTimeout(function () {
              _this.$store.commit(nsNamespace.GLOBAL + "/updatePageLoaded", true)
              _this.$store.commit(nsNamespace.GLOBAL + "/updateNodeRunning", false)
            }, 800)

          },
          function () {
            alert(this.$t('nsMain.node.mining.stopfail'))
            this.$store.commit(nsNamespace.GLOBAL + "/updatePageLoaded", true)
          })
      },
    }
  }
</script>

<style scoped>
    .operation-panel {
        float: left;
        width: 500px;
    }

    .operation-panel > div {
        background-color: #ffffff;
        padding: 40px 30px;
        border-radius: 6px;
        margin-bottom: 20px;
    }

    .operation-title {
        float: left;
        font-size: 20px;
    }

    .operation-connect-status > div {
        height: 24px;
        line-height: 24px;
    }

    .operation-connect-status .switch-button {
        float: right;
        cursor: pointer;
        color: #aaa;
    }

    .operation-connect-status .switch-button.switch-on {
        color: #318bEc;
    }

    .operation-connect-status .switch-button > img {
        display: inline-block;
        width: 70px;
        margin-top: -4px;
    }

    .operation-connect-status .switch-button > span {
        display: inline-block;
        padding-left: 10px;
        font-size: 24px;
    }

    .operation-wallet .wallet-operation-button {
        float: right;
        margin-left: 20px;
    }

    .operation-wallet .wallet-operation-button button {
        display: block;
        font-size: 16px;
        width: 90px;
        height: 30px;
        line-height: 24px;
        padding: 0;
        background-color: #ffffff;
        color: rgba(48, 139, 236, 1);
        border-radius: 15px;
        cursor: pointer;
        transition: background-color,color 0.2s;
    }

    .operation-wallet .wallet-operation-button button:hover {
        background-color: rgba(48, 139, 236, 1);
        color: #ffffff;
        transition: background-color,color 0.2s;
    }

    .operation-wallet .wallet-operation-button button:active {
        background-color: rgba(48, 139, 236, 0.8);
        color: #ffffff;
    }

    .operation-wallet .wallet-operation-button button,
    .operation-wallet .wallet-operation-button button:active {
        border: 1px solid rgba(48, 139, 236, 1);
        outline: 0 none;
    }

    .wallet-info-panel {
        clear: both;
    }

    .wallet-info-panel > div {
        padding: 30px 0;
    }

    .wallet-info-panel .info-title {
        font-size: 16px;
        line-height: 16px;
        color: #8d96b4;
        margin-bottom: 6px;
    }

    .wallet-info-panel .wallet-address {
        border-bottom: 1px solid #ddd;
    }

    .wallet-info-panel .wallet-address .info-text {
        font-size: 20px;
        line-height: 20px;
    }

    .wallet-info-panel .wallet-balance .info-text {
        font-size: 30px;
        line-height: 30px;
    }

    .wallet-info-panel .wallet-balance .info-text .balance-unit {
        font-size: 16px;
        padding-left: 10px;
    }

    .not-available-announcement {
        position: absolute;
        left: 0;
        top: 0;
        width: 100%;
        height: 100%;
        font-size: 24px;
        color: #aaaaaa;
        padding-left: 30px;
        padding-top: 100px;
    }
</style>

<style scoped>
    .node-info {
        clear: both;
        padding-top: 10px;
        font-size: 14px;
        font-weight: 300;
        color:#8d96b4
    }
</style>

<style scoped>
    .transfer-list {
        position: relative;
        float: right;
        width: 800px;
        min-height: 600px;
        padding-top: 40px;
        padding-bottom: 10px;
        border-radius: 6px;
        background-color: #ffffff;
    }

    .transfer-list > .transfer-list-title {
        font-size: 20px;
        line-height: 20px;
        padding-top: 2px;
        padding-bottom: 30px;
        padding-left: 30px;
    }

    .transfer-list .list-cell  {
        float: left;
        word-break: break-all;
        text-overflow: ellipsis;
        white-space:nowrap;
        overflow: hidden;
    }

    .transfer-list > .transfer-list-header {
        font-size: 16px;
        padding: 0 30px;
    }

    .transfer-list > .transfer-list-body > .transfer-row {
        padding: 12px 30px;
        border-bottom: 1px solid #d0d5e7;
    }

    .transfer-list > .transfer-list-body > .transfer-row:last-child {
        border: 0;
    }

    .cell-tx {
        width: 320px;
    }

    .cell-type {
        width: 140px;
    }

    .cell-value {
        padding-left: 20px;
        width: 100px;
    }

    .cell-height {
        width: 120px;
    }

    .cell-time {
        width: 200px;
    }

    .transfer-gray-row {
        /*background-color: #eee;*/
    }

    .transfer-list-header .list-cell {
        color: #8992b0;
        font-size: 16px;
    }

    .transfer-list-body .list-cell {
        color: #253a7e;
        font-size: 14px;
        cursor: default;
    }
</style>

<!--<div class="info-neighbor-panel nkn-card-shadow nkn-after-clear">-->
    <!--<div class="operation-title">Neighbor</div>-->
    <!--<div class="neighbor-list-header">-->
        <!--<div class="neighbor-list-row">-->
            <!--<div class="neighbor-list-cell id-cell">ID</div>-->
            <!--<div class="neighbor-list-cell ip-cell">IP</div>-->
            <!--<div class="neighbor-list-cell port-cell">Port</div>-->
        <!--</div>-->
    <!--</div>-->

    <!--<div class="neighbor-list-body">-->
        <!--<div class="neighbor-list-row"-->
             <!--v-for="(v, i) in neighbor" :key="i">-->
            <!--<div class="neighbor-list-cell id-cell">{{i}}</div>-->
            <!--<div class="neighbor-list-cell ip-cell">{{v.IpAddr}}</div>-->
            <!--<div class="neighbor-list-cell port-cell">{{v.Port}}</div>-->
        <!--</div>-->
    <!--</div>-->
<!--</div>-->

<style scoped>

    .info-neighbor-panel .neighbor-list-row,
    .info-neighbor-panel > div {
        float: left;
        clear: both;
        width: 100%;
    }

    .info-neighbor-panel .neighbor-list-header .neighbor-list-cell {
        color: #8992b0;
        font-size: 16px;
        height: 35px;
        line-height: 35px;
    }

    .info-neighbor-panel .neighbor-list-cell {
        float: left;
        width: 33%;
    }

    .info-neighbor-panel .id-cell {

    }

    .neighbor-list-body .neighbor-list-row {
        height: 35px;
        line-height: 35px;
        color: #253a7e;
        font-size: 14px;
        cursor: default;
        border-bottom: 1px solid #d0d5e7;
    }

    .neighbor-list-body .neighbor-list-row:last-child {
        border: 0;
    }
</style>

<style scoped>
    .reset-shell-warning {
        color: #c22;
        margin-bottom: 10px;
        font-size: 14px;
    }

    #nkn-transfer-model input {
        font-size: 14px;
        font-weight: 300;
        margin-bottom: 10px;
    }
</style>
