export default {
  namespaced: true,
  state: {
    pageLoaded: false,
    wallet: null,
    nodeRunning: false,
    nodeSyncState: null,
  },

  mutations: {
    updateSyncState(state, syncState) {
      state.nodeSyncState = syncState
    },
    updateNodeRunning(state, running) {
      state.nodeRunning = running
    },

    updatePageLoaded(state, loaded) {
      state.pageLoaded = loaded
    },

    updateWallet(state, wallet) {
      state.wallet = wallet
    },
  }
}