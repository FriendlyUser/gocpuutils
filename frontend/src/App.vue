<template>
  <v-app id="inspire">
    <v-navigation-drawer v-model="drawer" clipped fixed app>
      <v-list dense>
        <v-list-tile :to="{ path: '/foo' }">
          <v-list-tile-action>
            <v-icon>dashboard</v-icon>
          </v-list-tile-action>
          <v-list-tile-content>
            <v-list-tile-title>Dashboard</v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
        <v-list-tile :to="{ path: '/bar' }">
          <v-list-tile-action>
            <v-icon>settings</v-icon>
          </v-list-tile-action>
          <v-list-tile-content>
            <v-list-tile-title>Settings</v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
      </v-list>
    </v-navigation-drawer>
    <v-toolbar app fixed clipped-left>
      <v-toolbar-side-icon @click.stop="drawer = !drawer"></v-toolbar-side-icon>
      <v-toolbar-title>Application</v-toolbar-title>
    </v-toolbar>
    <v-content>
      <v-container fluid class="px-0">
        <v-layout justify-center align-center class="px-0">
         <transition>
          <keep-alive>
            <router-view></router-view>
          </keep-alive>
        </transition>
        </v-layout>
      </v-container>
    </v-content>
    <v-footer app fixed>
      <span style="margin-left:1em">&copy; You</span>
    </v-footer>
  </v-app>
</template>

<script>
// import HelloWorld from "./components/HelloWorld.vue";

export default {
  data: () => ({
    drawer: false,
    appTitle: 'Meal Prep',
    items: [
        { title: 'Menu', url: '/menu' },
        { title: 'Profile', url: '/about' },
        { title: 'Sign In', url: '/sign-in' },
        { title: 'Join', url: '/join' }
    ]
  })
};
</script>

<style>
.logo {
  width: 16em;
}
</style>