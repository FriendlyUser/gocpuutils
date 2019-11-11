<template>
  <div>
    <v-alert
      :value="true"
      type="success"
    >
      <h3>Num Cores</h3>
      <p>{{cores}}</p>
    </v-alert>
    <br />
    <hr />
    <v-card>
      <v-card-title primary-title>
        <div>
          <h3 class="headline mb-0">Num Cores</h3>
          <div> {{ cores }} </div>
        </div>
      </v-card-title>
    </v-card>
    <v-item-group>
      <v-card>
        <v-flex
          v-for="cpu in cpuInfo"
          :key="cpu.cpu"
          xs12
          md12
        >
          <v-card-title primary-title>
            <div>
              <h3 class="headline mb-0">Cpu Id: {{ cpu.cpu }}</h3>
            </div>
          </v-card-title>
          <v-card-text>
            <ol>
              <li>
                Modal Name: {{cpu.modelName}} 
              </li> 
              <li> 
                Physical Id: {{cpu.physicalId}}
              </li> 
              <li> 
                Num Cores: {{cpu.cores}}
              </li>
              <li>
                Vendor Id: {{cpu.vendorId}} 
              </li> 
              <li> 
                Frequency: {{cpu.mhz}} MHz
              </li>
            </ol>
          </v-card-text>
        </v-flex>
      </v-card>
    </v-item-group>
  </div>
</template>

<script>
import runtime from '@wailsapp/runtime';
export default {
  data() {
    return {
      cores: 0,
      cpuInfo: [
        {
          cacheSize: 0,
          coreId: "",
          cores: 4,
          cpu: 0,
          family: "11",
          flags: [],
          mhz: 1101,
          microcode: "",
          model: "",
          modelName: "Intel(R) Pentium(R) Silver N5000 CPU @ 1.10GHz",
          physicalId: "BFEBFBFF000706A1",
          stepping: 0,
          vendorId: "GenuineIntel"
        }
      ]
    };
  },
  mounted: function() {
    window.backend.Stats.GetDiskSerialNum(true)
    .then(result => {
      console.log(result)
      if (result) this.cores = result.disk
    })
    .catch(err => {
      console.log(error)
    })
    window.backend.Stats.GetInfo()
    .then(result => {
      console.log(result)
      if (result) this.cpuInfo = result
    })
    .catch(err => {
      console.log(err)
    })
  }
};
</script>