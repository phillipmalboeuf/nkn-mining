import {nsNamespace} from './nsNamespace'

import nsCreateAccount from '../pages/nsSetup/nsCreateAccount'
import nsGenWallet from '../pages/nsSetup/nsGenWallet'
import nsLoadWallet from '../pages/nsSetup/nsLoadWallet'
import nsShowWallet from '../pages/nsSetup/nsShowWallet'
import nsMain from '../pages/nsMain'
import nsSignIn from '../pages/nsSignIn'

const nsRouters = [
  {path: "/", name: nsNamespace.LOADING, component: nsCreateAccount},

  {path: "/create_account", name: nsNamespace.SETUP.CREATE_ACCOUNT, component: nsCreateAccount},
  {path: "/gen_wallet", name: nsNamespace.SETUP.GEN_WALLET, component: nsGenWallet},
  {path: "/load_wallet", name: nsNamespace.SETUP.LOAD_WALLET, component: nsLoadWallet},
  {path: "/show_wallet", name: nsNamespace.SETUP.SHOW_WALLET, component: nsShowWallet},

  {path: "/main", name: nsNamespace.MAIN, component: nsMain},
  {path: "/sign_in", name: nsNamespace.SIGN_IN, component: nsSignIn},
]

export {
  nsRouters
}