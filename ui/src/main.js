import g from 'guark'
import Vue from 'vue'
import App from './App.vue'
import store from './store'
import VueRouter from 'vue-router'
import { ToastPlugin } from 'bootstrap-vue'

import Alarms from './components/alarm/Alarms.vue'
import AlarmDetail from './components/alarm/AlarmDetail.vue'
import Radio from './components/radio/Radio.vue'
import { BootstrapVue } from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import './app.scss'
import helpers from '@/helper/helper'

const routes = [{ path: "/alarms", component: Alarms }, { path: "/alarm/new", component: AlarmDetail }, { path: "/alarm/:idx", component: AlarmDetail }, { path: "/radio", component: Radio }];
const router = new VueRouter({
	routes
});

Vue.config.productionTip = false
Vue.prototype.helpers = helpers
Vue.helpers = helpers

Vue.use(VueRouter)
Vue.use(BootstrapVue)
Vue.use(ToastPlugin)
Vue.use(helpers)

new Vue({
	store,
	router,
	render: h => h(App),
	created: () => g.hook("created"),
	mounted: () => g.hook("mounted"),
}).$mount('#app')