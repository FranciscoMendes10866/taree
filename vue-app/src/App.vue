<template>
  <div>
    <Navbar />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Navbar from "./components/Navbar.vue";

export default defineComponent({
  name: "App",
  components: {
    Navbar
  }
});
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Titillium+Web&display=swap');
* {
  font-family: 'Titillium Web', sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>
