import g from 'guark'
import Vue from 'vue'
import App from './App.vue'
import store from './store'
import VueRouter from 'vue-router'

import Alarms from './components/alarm/Alarms.vue'
import AlarmDetail from './components/alarm/AlarmDetail.vue'

// TODO: add home view
// TODO: add radio view


// const AlarmDetail = {
// 	template: `<div>
// 	  {{$route.params.alarmidx}} 
// 	</div>`
// };

// const Alarms = {
// 	template: `<alarms> </alarms>`
// }


const routes = [{ path: "/alarms", component: Alarms }, {path: "/alarm/:idx", component: AlarmDetail}];
const router = new VueRouter({
	routes
});

Vue.config.productionTip = false
Vue.use(VueRouter)
new Vue({
	store,
	router,
	render: h => h(App),
	created: () => g.hook("created"),
	mounted: () => g.hook("mounted"),
}).$mount('#app')