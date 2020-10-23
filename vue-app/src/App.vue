<template>
  <div>
    <section class="hero has-background-white-ter is-fullheight">
      <div class="hero-body">
        <div class="container">
          <div class="columns" v-if="!state.firstAPI || !state.secondAPI">
            <div class="column">
              <div class="card has-background-danger">
                <div class="card-content has-text-centered">
                  <p class="title has-text-white">
                    One of the API's is not working properly.
                  </p>
                </div>
              </div>
            </div>
          </div>
          <div class="columns" v-else>
            <div class="column">
              <div class="card has-background-success">
                <div class="card-content">
                  <p class="title has-text-white">
                    {{ state.firstAPI }}
                  </p>
                </div>
              </div>
            </div>
            <div class="column">
              <div class="card has-background-success">
                <div class="card-content">
                  <p class="title has-text-white">
                    {{ state.secondAPI }}
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, reactive } from "vue";
import axios from "axios";

export default defineComponent({
  setup() {
    // init state
    const state = reactive({
      firstAPI: null,
      secondAPI: null
    });
    // http request
    async function Fetch() {
      await axios.get("http://localhost/first-api").then(res => {
        state.firstAPI = res.data;
      });
      await axios.get("http://localhost/second-api").then(res => {
        state.secondAPI = res.data;
      });
    }
    // when the page mounts
    onMounted(Fetch);
    // returns the state (to be used by the template)
    // and returns the Fetch function
    return {
      state,
      Fetch
    };
  }
});
</script>

<style src="./assets/styles.css"></style>
