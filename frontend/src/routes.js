import Home from './components/HelloWorld.vue';
import Simple from './components/Simple.vue'
import Wails from './components/Wails.vue'
import CPU from './components/CPUUsage.vue'
import Stats from './components/CPUStats.vue'

const routes = [
  { path: '/', component: Home },
  { path: '/about', component: Simple },
  { path: '/foo', component: Home },
  { path: '/cpu', component: CPU },
  { path: '/wails', component: Wails },
  { path: '/stats', component: Stats, icon: 'megaevolve' }
]

export default routes;