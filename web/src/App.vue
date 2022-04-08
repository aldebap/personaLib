<template>
  <HelloWorld msg="personaLib" />
</template>

<script>
import HelloWorld from "./components/HelloWorld.vue";
import personaLib from "./components/personaLib.vue";

export default {
  name: "App",
  components: {
    HelloWorld,
    personaLib,
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
