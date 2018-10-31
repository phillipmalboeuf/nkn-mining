import CryptoJS from 'crypto-js'

function HmacSHA256(message, secret) {
  secret = secret || message
  return CryptoJS.HmacSHA256(message, secret).toString()
}

function AESEnc(plaintext, pwd, rebuildPwd = true) {
  let key = null
  if(rebuildPwd) {
    key = HmacSHA256(pwd, pwd)
  } else {
    key = pwd
  }

  let ret = null

  let iv = CryptoJS.enc.Hex.parse(key.slice(0, 32))
  key = CryptoJS.enc.Hex.parse(key)
  plaintext = CryptoJS.enc.Utf8.parse(plaintext)

  let encCfg = {
    iv:iv,
    mode: CryptoJS.mode.CFB ,
    padding: CryptoJS.pad.NoPadding
  }

  try {
    ret = CryptoJS.AES.encrypt(plaintext, key, encCfg)
      .ciphertext.toString()
  } catch (e) {
    console.error(e)
  }

  return ret
}

function AESDec(ciphertext, pwd, rebuildPwd = true) {
  let key = null
  if(rebuildPwd) {
    key = HmacSHA256(pwd, pwd)
  } else {
    key = pwd
  }

  let iv = CryptoJS.enc.Hex.parse(key.slice(0, 32))
  key = CryptoJS.enc.Hex.parse(key)
  ciphertext = CryptoJS.lib.CipherParams.create({
    ciphertext: CryptoJS.enc.Hex.parse(ciphertext)
  })

  let encCfg = {
    iv:iv,
    mode: CryptoJS.mode.CFB ,
    padding: CryptoJS.pad.NoPadding
  }

  let ret = null
  try {
    ret = CryptoJS.AES.decrypt(ciphertext, key, encCfg).toString(CryptoJS.enc.Utf8)
  } catch (e) {
    console.error(e)
  }

  return ret
}

export default {
  HmacSHA256,
  AESEnc,
  AESDec,
}