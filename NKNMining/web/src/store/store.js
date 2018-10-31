import Vue from 'vue'
import vuex from 'vuex'
import {nsNamespace} from "../js/nsNamespace"

Vue.use(vuex)

import global from './global'

let modules = {}
modules[nsNamespace.GLOBAL] = global

export default new vuex.Store({
  modules: modules
})