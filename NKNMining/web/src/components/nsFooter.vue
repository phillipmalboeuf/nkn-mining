<template>
    <footer class="nkn-footer">
        <i18n path="nsFooter.copyright" tag="div" class="nkn-copyright" >
            <span place="i">&copy;</span>
        </i18n>
        <div class="nkn-attention">{{$t("nsFooter.attention")}}</div>
    </footer>
</template>

<script>
  export default {
    name: "ns-footer"
  }
</script>

<style scoped>
    .nkn-footer {
        position: relative;
        height: 80px;
        width: 100%;
        bottom: 0;
        background: #ffffff;
        line-height: 80px;
        text-align: center;
    }

    .nkn-footer .nkn-copyright {
        color: #aaaaaa;
        bottom: 0;
        left: 0;
        right: 0;
        height: 40px;
        line-height: 40px;
        font-size: 12px;
        font-weight: 100;
    }

    .nkn-footer .nkn-attention {
        color: #aaaaaa;
        bottom: 0;
        left: 0;
        right: 0;
        height: 40px;
        line-height: 20px;
        font-size: 12px;
        font-weight: 100;
    }
</style>