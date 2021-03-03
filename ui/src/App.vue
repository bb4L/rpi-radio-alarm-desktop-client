<script src="https://unpkg.com/vue-router/dist/vue-router.js"></script>

<template>
  <div id="app">
    <div>
      <router-link class="btn" to="/alarms">Alarms</router-link>
      <router-link class="btn" to="/radio">Radio</router-link>
	  <div> <h2>Router view</h2> </div>
      <router-view></router-view>
	  <div> <h2>End Router view</h2> </div>
    </div>
  </div>
</template>

<style lang="scss">
body,
html {
  width: 100%;
  height: 100%;
  padding: 0;
  margin: 0;
  border: 0;
}

#app {
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica,
    Arial, sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol";
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  height: 100%;
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  color: rgba(245, 64, 32, 0.8);
  background-color: black;
//   background-image: linear-gradient(to top, #fad0c4 0%, #ffd1ff 100%);
}

#nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}

// Disable image drag and drop.
img {
  -webkit-user-drag: none;
  -khtml-user-drag: none;
  -moz-user-drag: none;
  -o-user-drag: none;
  user-drag: none;
}

* {
  user-select: none;
  -moz-user-select: none;
  -webkit-user-select: none;
  -ms-user-select: none;
  touch-callout: none;
  -webkit-touch-callout: none;
}

input {
  user-select: auto;
  -moz-user-select: auto;
  -webkit-user-select: auto;
  -ms-user-select: auto;
}

h3 {
	margin: 40px 0 0;
}

ul {
	list-style-type: none;
	padding: 0;
}

li {
	display: inline-block;
	margin: 0 10px;
}

a {
	color: #42b983;
}

.btn {
	margin: 3px;
	background: #5556e4;
	padding: 12px;
	border-radius: 10px;
	cursor: pointer;
	border: none;
	outline: none;
	font-size: 15px;
	font-weight: 500;
	color: #FFF;
	box-shadow: 0pt 1px 4px 1px rgba(85, 86, 228,0.5);
}

</style>

<script>
import g from "guark";

export default {
  name: "App",

//   components: {
//     HelloWorld,
//   },

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