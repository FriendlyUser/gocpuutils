<template>
  <v-container fluid class="px-0">
    <v-layout>
      <v-flex xs12 sm6 offset-sm3>
        <v-card raised="raised" class="pa-4 ma-4">
          <v-layout justify-center align-center class="pa-4 ma-4">
            <v-img :src="require('../assets/images/logo.png')"></v-img>
          </v-layout>
          <v-card-actions>
            <v-layout justify-center align-center class="px-0">
              <v-btn color="blue" dark @click="getFilesInDir">Get Files</v-btn>
            </v-layout>
          </v-card-actions>
        </v-card>
        <treemap-chart :data="chartData" :options="options" />
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

<script>
import 'tui-chart/dist/tui-chart.css'
import { treemapChart } from '@toast-ui/vue-chart'
import Wails from '@wailsapp/runtime'
export default {
  components: {
    'treemap-chart': treemapChart
  },
  data() {
    return {
      message: " ",
      raised: true,
      dialog: false,
      chartData: { // for 'data' prop of 'bar-chart'
          series: [
            {
                label: 'Documents',
                children: [
                    {
                        label: 'docs',
                        children: [
                            {
                                label: 'pages',
                                value: 1.3
                            },
                            {
                                label: 'keynote',
                                value: 2.5
                            },
                            {
                                label: 'numbers',
                                value: 1.2
                            }
                        ]
                    },
                    {
                        label: 'photos',
                        value: 5.5
                    },
                    {
                        label: 'videos',
                        value: 20.7
                    }
                ]
            }, {
                label: 'Downloads',
                children: [
                    {
                        label: 'recents',
                        value: 5.3
                    }, {
                        label: '2015',
                        value: 10.1
                    }, {
                        label: '2014',
                        value: 8.2
                    }
                ]
            }, {
                label: 'Application',
                value: 16.4
            }, {
                label: 'Desktop',
                value: 4.5
            }
          ],
      },
      options: {
        series: {
          showLabel: true,
          zoomable: false,
          useLeafLabel: true
        },
        tooltip: {
          suffix: 'GB'
        }
      },
      theme: {
        series: {
          colors: [
              '#83b14e', '#458a3f', '#295ba0', '#2a4175', '#289399',
              '#289399', '#617178', '#8a9a9a', '#516f7d', '#dddddd'
          ],
          borderColor: '#cccccc',
          borderWidth: 5
        }
      }
    }
  },
  methods: {
    getMessage: function() {
      var self = this;
      window.backend.basic().then(result => {
        self.message = result;
        self.dialog = true;
      });
    },
    getFilesInDir: function() {
      var self = this;
      Wails.Log.Info("New Code is Going on");
      // console.log("Doing Something")
      // console.log(window.backend)
      window.backend.Todos.GetFiles().then(result => {
        self.chartData = {
          series: [JSON.parse(result)]
        }
      }).catch(err => {
        // console.log(err)
        console.log(err)
      })
      // window.backend.Todos.getFiles().then(result => {
      //   console.log(result)
      //   console.log("Got Here")
      //   Wails.Log.Info("Do Something Here")
      // });
    }
  },
  mounted() {
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1 {
  margin-top: 2em;
  position: relative;
  min-height: 5rem;
  width: 100%;
}
a:hover {
  font-size: 1.7em;
  border-color: blue;
  background-color: blue;
  color: white;
  border: 3px solid white;
  border-radius: 10px;
  padding: 9px;
  cursor: pointer;
  transition: 500ms;
}
a {
  font-size: 1.7em;
  border-color: white;
  background-color: #121212;
  color: white;
  border: 3px solid white;
  border-radius: 10px;
  padding: 9px;
  cursor: pointer;
}
</style>
