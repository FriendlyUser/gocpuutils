import Home from './components/HelloWorld.vue';
import Simple from './components/Simple.vue'

const routes = [
  { path: '/', component: Home },
  { path: '/about', component: Simple },
  { path: '/foo', component: Home },
  { path: '/bar', component: Simple }
]

export default routes;