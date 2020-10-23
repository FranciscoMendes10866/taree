<template>
  <div>
    <Navbar />
    <h1>{{ firstAPI }}</h1>
    <h1>{{ secondAPI }}</h1>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from "vue";
import axios from "axios";
import Navbar from "./components/Navbar.vue";

export default defineComponent({
  name: "App",
  components: {
    Navbar
  },
  setup() {
    // init state
    const firstAPI = ref({})
    const secondAPI = ref({})
    // http request
    async function fetchAPIS() {
      await axios.get('http://localhost:8538')
        .then((res) => {
          firstAPI.value = res.data
        })
      await axios.get('http://localhost:6518')
        .then((res) => {
          secondAPI.value = res.data
        })
    }
    // when the page mounts
    onMounted(fetchAPIS)
    return {
      fetchAPIS
    }
  }
});
</script>

<style src="./assets/styles.css"></style>
