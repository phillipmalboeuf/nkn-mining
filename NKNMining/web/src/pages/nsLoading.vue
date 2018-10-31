<template>
    <div>
        <ns-loading/>
    </div>
</template>

<script>
  import NsLoading from "../components/nsLoading";
  import {loadPage} from "../js/nsLoading";

  export default {
    components: {NsLoading},
    name: "ns-loading-page",
    mounted() {
      loadPage.call(this)
    },
  }
</script>

<style scoped>

</style>