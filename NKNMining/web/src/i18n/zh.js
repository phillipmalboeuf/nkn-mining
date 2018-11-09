export default {
  mark: 'NKN',
  nsNav: {
    signOut: '注销',
    resetNodeShell: '重置 @:mark 矿机'
  },
  nsFooter: {
    copyright: 'Copyright {i} 2017-2018 @:mark    |    版权所有',
    attention: '目前软件正处于早期研发阶段，可能存在部分问题。请勿用于正式途径。 '
  },
  nsInput: {
    account: {
      title: '矿机账号',
      placeholder: '请输入您的 @:mark 矿机账号',
      errorInfo: '仅允许输入字母数字！'
    },
    password: {
      title: '矿机密码',
      placeholder: '请输入您的 @:mark 矿机密码',
      errorInfo: '仅允许8-20个字符！'
    },
    rePassword: {
      title: '确认矿机密码',
      placeholder: '请再次输入您的 @:mark 矿机密码',
      errorInfo: '两次输入的矿机密码不一致！'
    },
    sn: {
      title: '序列号',
      placeholder: '请输入您的序列号',
      errorInfo: '您输入的序列号有误！'
    },
    wallet: {
      title: ' ',
      placeholder: '请上传你的钱包文件',
      errorInfo: '请选择并上传你的钱包文件！'
    },
    walletPassword: {
      title: '钱包密码',
      placeholder: '请输入您上传钱包的密码',
      errorInfo: '您输入的密码和钱包不匹配！'
    }
  },
  nsSignIn: {
    titleLabel: '欢迎来到',
    title: '@:mark 节点',
    signInbtn: '登陆',
    loginfail:'登陆失败！'
  },
  nsMain: {
    node: {
      title: '节点连接',
      on: '开',
      off: '关',
      mining:{
        fail: '开启矿机失败。请检查服务器日志。',
        error: "您不能在节点更新时开启矿机。",
        stopfail: '关闭矿机失败。请检查服务器日志。'
      },
      nodeStatus:{
        prefix: '节点状态：',
        init: '@:nsMain.node.nodeStatus.prefix 同步中 ...',
        starting: '正在开启',
        updating: '@:nsMain.node.nodeStatus.prefix 更新中',
        stopping: '正在关闭',
        stoped: '@:nsMain.node.nodeStatus.prefix 已关闭',
        syncing: '@:nsMain.node.nodeStatus.prefix 同步中',
        persisting: '@:nsMain.node.nodeStatus.prefix 同步完成',
        mining: '@:nsMain.node.nodeStatus.prefix 挖矿中'
      },
      version: {
        init: '正在查询版本 ...'
      },
      NKNBlockHeight: '@:mark 网络区块高度：',
      myBlockHeight: '我的区块高度：'
    },
    wallet: {
      title: '我的钱包',
      downloadbtn: '下载',
      transferbtn: '转账',
      addressLabel: '地址',
      balance: {
        init: '正在查询余额 ...'
      },
      balanceLabel: '我的余额',
      balanceUnit: '@:mark'
    },
    neighbors: {
      title: '相邻节点',
      ID: '节点ID',
      IP: '节点IP',
      port: '节点端口'
    },
    rewards: {
      title: '近期挖矿奖励',
      transaction: '交易',
      value: 'Value',
      height: 'Height',
      time: 'Time'
    },
    transfer: {
      title: 'Transfer @:mark',
      addressPlaceholder: 'target address',
      countPlaceholder: 'how much @:mark to transfer',
      passwordPlaceholder: 'wallet password',
      cancel: 'Cancel',
      confirm: 'Transfer',
      alertInfo: {
        default: 'transfer failed!'
      }
    },
    resetdlg: {
      title: 'Reset shell confirm',
      warning: 'Warning: This operation will reset all your @:mark data, include your Wallet!',
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