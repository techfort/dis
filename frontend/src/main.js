import Vue from 'vue';
import App from './App.vue';
import filters from './filters';
import VueNativeSock from 'vue-native-websocket';
import { store } from './store';
import VueRouter from 'vue-router';
import Watch from './components/Watch.vue';
import Event from './components/Event.vue';
import Query from './components/Query.vue';
import Insert from './components/Insert.vue';
import Sets from './components/Sets.vue';

Vue.use(VueNativeSock, 'ws://localhost:1666/ws', { store });
Vue.config.productionTip = false;
Vue.use(VueRouter);

const routes = [
  { path: '/watch', component: Watch },
  { path: '/query', component: Query },
  { path: '/insert', component: Insert },
  { path: '/sets', component: Sets }
];

const router = new VueRouter({
  routes
});

new Vue({
  router,
  store,
  filters,
  render: h => h(App),
  components: {
    Watch,
    Query,
    Event,
    Insert,
    Sets
  }
}).$mount('#app');
