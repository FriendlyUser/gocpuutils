<template>
  <v-container fluid class="px-0">
    <v-layout>
      <v-flex xs12 sm6 offset-sm3>
        <v-card raised="raised" class="pa-4 ma-4">
          <v-layout justify-center align-center class="pa-4 ma-4">
            <v-img :src="require('../assets/images/logo.png')"></v-img>
          </v-layout>
        </v-card>
      </v-flex>
    </v-layout>
    <div class="text-xs-center">
      <v-dialog v-model="dialog" width="500">
        <v-card>
          <v-card-title class="headline" primary-title>Message from Go</v-card-title>
          <v-card-text>{{message}}</v-card-text>
          <v-divider></v-divider>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="primary" flat @click="dialog = false">Awesome</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </div>
  </v-container>
</template>
