<template>
    <footer class="nkn-footer">
        <div class="nkn-copyright">Copyright &copy; 2017-2018 NKN    |    All rights reserved</div>
        <div class="nkn-attention">This software is in the early development stage and may not have all functions working properly. It should be used only for testing now. </div>
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