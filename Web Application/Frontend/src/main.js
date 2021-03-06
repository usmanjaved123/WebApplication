import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify';
import axios from 'axios'
import VueAxios from 'vue-axios'
import 'material-design-icons-iconfont/dist/material-design-icons.css'
import '@mdi/font/css/materialdesignicons.css'
Vue.config.productionTip = false
Vue.use(VueAxios, axios)

new Vue({
  vuetify,
  render: h => h(App)
}).$mount('#app')