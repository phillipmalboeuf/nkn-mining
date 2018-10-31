import Crypto from '../crypto/algorithm'
import Is from 'is_js'

function encrypt(plaintext, key, rebuildKey) {
  let cipherData =  Crypto.AESEnc(plaintext, key, rebuildKey)

  return {
    Data: cipherData
  }
}

let RequestData = function () {
  let dataString = ""
  let plainData = ""

  this.setData = function (data) {
    if(!Is.string(data)) {
      dataString = JSON.stringify(data)
    } else {
      dataString = data
    }
  }

  this.setPlainData = function (data) {
    if(!Is.string(data)) {
      plainData = JSON.stringify(data)
    } else {
      plainData = data
    }
  }

  this.encryptedData = function(key, rebuildKey = true) {
    return encrypt(dataString, key, rebuildKey)
  }

  this.fullData = function (key, rebuildKey = true) {
    let encryptedData = this.encryptedData(key, rebuildKey)
    encryptedData.PlainData = plainData
    return encryptedData
  }
}

export  {
  RequestData,
}