<template>
    <div class="nkn-login-item">
        <div v-if="'' !== config.title" class="item-title text-label-gray" v-html="config.title"></div>
        <div class="item-input input-group">
            <input class="form-control"
                   :maxlength="config.maxSize ? config.maxSize:-1"
                   :type="config.inputType"
                   :class="this.config.inputId"
                   :placeholder="config.placeholder">
            <div v-if="config.hasAppend" class="input-group-append">
                <span class="input-group-text" @click="switchInputType">
                    <img v-if="openEyeToShow" class="item-append-icon" src="../assets/img/icon/eye-open.png"/>
                    <img v-else class="item-append-icon" src="../assets/img/icon/eye-close.png"/>
                </span>
            </div>
        </div>
        <div class="input-error-tooltip" v-text="config.errorInfo ? config.errorInfo : ''"></div>
    </div>
</template>

<script>
  export default {
    name: "ns-input-item",
    props: ["config"],
    methods: {
      switchInputType() {
        this.config.inputType = ('text' === this.config.inputType) ? 'password':'text'
      }
    },
    computed: {
      openEyeToShow() {
        return ('text' === this.config.inputType)
      }
    }
  }
</script>

<style scoped>
    .nkn-login-item {
        height: 65px;
        width: 100%;
        padding-right: 94px;
        margin-bottom: 50px;
    }

    .nkn-login-item > .item-title {
        font-size: 14px;
        height: 20px;
    }

    .nkn-login-item > .item-input {
        height: 45px;
        border-bottom: 1px solid #243a80;
    }

    .nkn-login-item > .item-input input,
    .nkn-login-item > .item-input input:disabled,
    .nkn-login-item > .item-input input:active,
    .nkn-login-item > .item-input input:focus {
        border: 0;
        border-radius: 0;
        padding-left: 0;
        box-shadow: none;
        outline: 0 none;
        font-size: 16px;
        background-color: transparent;
    }

    .nkn-login-item > .item-input input::placeholder {
        color: #c7cbd2;
    }

    .nkn-login-item > .item-input span {
        background: transparent;
        border: 0;
        cursor: pointer;
    }

    .nkn-login-item > .item-input .item-append-icon {
        width: 30px;
    }

    .nkn-login-item > .input-error-tooltip {
        color: #c22;
        padding-top: 4px;
        font-size: 12px;
        font-weight: 300
    }
</style>