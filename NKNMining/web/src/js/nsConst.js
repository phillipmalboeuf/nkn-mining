let inputIdPrefix = function () {
  return 'ns-input-'
}

let serverStatus = new function () {
  this.NS_STATUS_GEN_WALLET = ()=>{return 1}
  this.NS_STATUS_NODE_RUNNING = ()=>{return 2}
  this.NS_STATUS_NODE_EXITED = ()=>{return 3}
  this.NS_STATUS_UPDATE_BIN = ()=>{return 4}
  this.NS_STATUS_INITIALIZATION = ()=>{return 5}
}

export {
  inputIdPrefix,
  serverStatus,
}