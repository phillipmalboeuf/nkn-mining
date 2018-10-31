<template>
    <div id="g-loading-page">
        <img :src="'./static/img/spin.svg'" width="64" height="64">
    </div>
</template>

<script>
  export default {
    name: "ns-loading",
  }
</script>

<style scoped>
    #g-loading-page {
        position: fixed;
        left: 0;
        top: 0;
        width: 100%;
        height: 100%;
        background-color: rgba(255, 255, 255,0.8);
        /*visibility: hidden;*/
    }

    #g-loading-page > img {
        display: block;
        margin: 0 auto;
        margin-top: 300px;
    }
</style>