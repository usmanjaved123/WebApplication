<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png">
    <h1>Order Details</h1>
    <Order msg="Welcome to Your Vue.js App"/>
  </div>
</template>

<script>
//import Order from './components/Order.vue'
import Order from './components/Order';

export default {
  name: 'App',

  components: {
    Order
  },

  data: () => ({
    //
  }),
}
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
