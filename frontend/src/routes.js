import Home from './components/HelloWorld.vue';
import Simple from './components/Simple.vue'
import Wails from './components/Wails.vue'
import CPU from './components/CPUUsage.vue'
import Stats from './components/CPUStats.vue'

const routes = [
  { title: 'Home', path: '/', component: Home, icon: 'home' },
  { title: 'About', path: '/about', component: Simple, icon: 'lock' },
  { title: 'Foo', path: '/foo', component: Home, icon: 'event' },
  { title: 'CPU', path: '/cpu', component: CPU, icon: 'computer' },
  { title: 'Wails', path: '/wails', component: Wails, icon: 'lock' },
  { title: 'Stats', path: '/stats', component: Stats, icon: 'info' }
]

export default routes;