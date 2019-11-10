import 'babel-polyfill';
import Vue from 'vue';

// Setup Vuetify
import Vuetify from 'vuetify'
import 'vuetify/dist/vuetify.min.css'
Vue.use(Vuetify);
import 'material-design-icons-iconfont';

import VueApexCharts from 'vue-apexcharts'

Vue.use(VueApexCharts)
Vue.component('apexchart', VueApexCharts)
import App from './App.vue';

Vue.config.productionTip = false;
Vue.config.devtools = true;

import VueRouter from 'vue-router';
Vue.use(VueRouter)

import routes from './routes';

const router = new VueRouter({
	mode: 'history',
  routes
})
import * as Wails from '@wailsapp/runtime';
Wails.Init(() => {
	new Vue({
		router,
		render: h => h(App)
	}).$mount('#app')
});
