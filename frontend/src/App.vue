<template>
  <v-app id="inspire" dark>
    <v-navigation-drawer v-model="drawer" clipped fixed app>
       <v-list dense>
        <div v-for="item in items" :key="item.title">
          <v-list-tile
            :slot="item"
            :to="item.url"
          >
            <v-list-tile-action>
              <v-icon>{{ item.icon }}</v-icon>
            </v-list-tile-action>
            <v-list-tile-content>
              <v-list-tile-title >{{ item.title }}</v-list-tile-title>
            </v-list-tile-content>
          </v-list-tile>         
        </div>
      </v-list>
    </v-navigation-drawer>
    <v-toolbar app fixed clipped-left>
      <v-toolbar-side-icon @click.stop="drawer = !drawer"></v-toolbar-side-icon>
      <v-toolbar-title>Application</v-toolbar-title>
    </v-toolbar>
    <v-content>
      <transition>
        <keep-alive>
          <router-view></router-view>
        </keep-alive>
      </transition>
    </v-content>
    <v-footer app fixed>
      <span style="margin-left:1em">&copy; You</span>
    </v-footer>
  </v-app>
</template>

<script>
// import HelloWorld from "./components/HelloWorld.vue";
import routes from './routes'
const newRoutes = routes.map(route => {
  route.url = route.path
  return route
})

console.log(newRoutes)
export default {
  data: () => ({
    drawer: false,
    appTitle: 'Meal Prep',
    items: newRoutes
  })
};
</script>

<style>
.logo {
  width: 16em;
}
</style>