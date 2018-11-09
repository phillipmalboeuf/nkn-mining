export default {
  nsNav: {
    signOut: 'Sign out',
    resetNodeShell: 'Reset Node shell'
  },
  nsFooter: {
    copyright: 'Copyright {i} 2017-2018 NKN    |    All rights reserved',
    attention: 'This software is in the early development stage and may not have all functions working properly. It should be used only for testing now. '
  },
  nsInput: {
    account: {
      title: 'Account',
      placeholder: 'Your NKN shell account',
      errorInfo: 'Alphanumeric only!'
    },
    password: {
      title: 'Your account password',
      placeholder: '8-20 characters',
      errorInfo: 'Please input 8-20 characters'
    },
    rePassword: {
      title: 'Confirm account password',
      placeholder: '8-20 characters',
      errorInfo: 'The password for the two input is inconsistent!'
    },
    sn: {
      title: 'System initialization serial number',
      placeholder: '40 characters',
      errorInfo: 'Invalid serial number'
    },
    wallet: {
      title: ' ',
      placeholder: 'Wallet file to load',
      errorInfo: 'please select a wallet file'
    },
    walletPassword: {
      title: 'Wallet password',
      placeholder: 'password ot the wallet to load',
      errorInfo: 'Please check if the password matches the wallet.'
    }
  },
  nsSignIn: {
    titleLabel: 'Welcome to',
    title: 'NKN NODE',
    signInbtn: 'Sign in',
    loginfail:'login failed!'
  },
  nsMain: {
    node: {
      title: 'Node connect',
      on: 'on',
      off: 'off',
      mining:{
        fail: 'start mining failed. please check log on the server.',
        error: "you can't mining when node updating.",
        stopfail: 'stop mining failed. please check log on the server.'
      },
      nodeStatus:{
        prefix: 'Node status: ',
        init: '@:nsMain.node.nodeStatus.prefix sync ...',
        starting: 'starting',
        updating: '@:nsMain.node.nodeStatus.prefix updating',
        stopping: 'stopping',
        stoped: '@:nsMain.node.nodeStatus.prefix stopped',
        syncing: '@:nsMain.node.nodeStatus.prefix syncing',
        persisting: '@:nsMain.node.nodeStatus.prefix persisting',
        mining: '@:nsMain.node.nodeStatus.prefix mining'
      },
      version: {
        init: 'Querying version ...'
      },
      NKNBlockHeight: 'NKN network block height:',
      myBlockHeight: 'my block height:'
    },
    wallet: {
      title: 'Wallet',
      downloadbtn: 'download',
      transferbtn: 'transfer',
      addressLabel: 'Address',
      balance: {
        init: 'Querying balance ...'
      },
      balanceLabel: 'Balance',
      balanceUnit: 'NKN'
    },
    neighbors: {
      title: 'Neighbors',
      ID: 'ID',
      IP: 'IP',
      port: 'Port'
    },
    rewards: {
      title: 'Your latest mining rewards',
      transaction: 'Transaction',
      value: 'Value',
      height: 'Height',
      time: 'Time'
    },
    transfer: {
      title: 'Transfer NKN',
      addressPlaceholder: 'target address',
      countPlaceholder: 'how much NKN to transfer',
      passwordPlaceholder: 'wallet password',
      cancel: 'Cancel',
      confirm: 'Transfer',
      alertInfo: {
        default: 'transfer failed!'
      }
    },
    resetdlg: {
      title: 'Reset shell confirm',
      warning: 'Warning: This operation will reset all your NKN data, include your Wallet!',
      passwordPlaceholder: 'Account password',
      cancel: 'Cancel',
      reset: 'Reset',
      alertInfo: {
        wrongPass: 'wrong password!',
        success: 'reset node shell success'
      }
    }
  },
  nsCreateAccount: {
    titleLabel: 'setup - step 1',
    title: 'Create account',
    nextStepbtn: 'NEXT'
  },
  nsGenWallet: {
    titleLabel: 'setup - step 2',
    title: 'Generate wallet',
    link: 'Load wallet',
    nextStepbtn: 'NEXT'
  },
  nsLoadWallet: {
    titleLabel: 'setup - step 2',
    title: 'Load wallet',
    upload: 'Select file',
    link: 'Generate wallet',
    nextStepbtn: 'NEXT',
    netWorkError: 'Network error, please try again later.'
  },
  nsShowWallet: {
    titleLabel: 'setup - step 3',
    title: 'Download wallet',
    privateKey: 'Private key',
    address: 'Wallet address',
    signInbtn: 'Sign in',
    downloadbtn: 'Download wallet'
  }
}