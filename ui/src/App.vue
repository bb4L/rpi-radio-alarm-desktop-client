<script src="https://unpkg.com/vue-router/dist/vue-router.js"></script>

<template>
  <div id="app">
    <nav class="navbar navbar-dark">
      <router-link class="nav-brand text-primary" to="/"
        >Rpi Radio Alarm</router-link
      >
      <router-link class="nav-item text-primary" to="/alarms"
        >Alarms</router-link
      >
      <router-link class="nav-item text-primary" to="/radio">Radio</router-link>
    </nav>
    <div class="container bg-dark">
      <div class="row">
        <div class="card col">
          <router-view>
          </router-view>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import g from "guark";

export default {
  name: "App",

  created() {
    g.env().then((env) => {
      // Disable right click menu on production, you can implement your own!
      if (!!env.dev_mode) {
        document.addEventListener("contextmenu", (e) => e.preventDefault());
      }
    });
  },
};
</script>