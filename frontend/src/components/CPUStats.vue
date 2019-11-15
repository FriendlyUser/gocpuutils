<template>
  <v-container grid-list-md text-xs-center>
  <v-layout row wrap>
    <v-flex xs4>
    <v-alert
      :value="true"
      type="success"
    >
      <h3>Num Cores</h3>
      <p>{{cores}}</p>
    </v-alert>
    </v-flex>
    <v-flex
      v-for="cpu in cpuInfo"
      :key="cpu.cpu"
      xs12
      md12
    >
      <v-card>
          <v-card-title primary-title>
            <div>
              <h3 class="headline mb-0">Cpu Id: {{ cpu.cpu }}</h3>
            </div>
          </v-card-title>
          <v-card-text>
            <ul>
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
            </ul>
          </v-card-text>
        </v-card>
      </v-flex>
      <h3> 
        Processors
      </h3>
      <v-flex
        v-for="cpu in cpuTimes"
        :key="cpu.cpu"
        xs12
        md12
      >
        <v-card>
          <v-card-title primary-title>
            <div>
              <h3 class="headline mb-0">Cpu Id: {{ cpu.cpu }}</h3>
            </div>
          </v-card-title>
          <v-card-text>
            <ul>
              <li>
                Guest: {{cpu.guest}} 
              </li> 
              <li> 
                guestNice: {{cpu.guestNice}}
              </li> 
              <li> 
                Idle: {{cpu.Idle}}
              </li>
              <li>
                ioWait: {{cpu.ioWait}} 
              </li> 
              <li> 
                IRQ: {{cpu.irq}}
              </li>
              <li> 
                IRQ: {{cpu.softirq}}
              </li>
              <li> 
                Steal: {{cpu.steal}}
              </li>
              <li> 
                Guest: {{cpu.guest}}
              </li>
              <li> 
                GuestNice: {{cpu.guestNice}}
              </li>
            </ul>
          </v-card-text>
        </v-card>
      </v-flex>
    </v-layout>
  </v-container>
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
      ],
      cpuTimes: [
        {

        }
      ]
    };
  },
  mounted: function() {
    window.backend.Stats.GetDiskSerialNum(true)
    .then(result => {
      if (result) this.cores = result.disk
    })
    .catch(err => {
      console.log(error)
    })
    window.backend.Stats.GetInfo()
    .then(result => {
      if (result) this.cpuInfo = result
    })
    window.backend.Stats.GetTimes()
    .then(result => {
      console.log(result)
      if (result) this.cpuTimes = result
    })
    .catch(err => {
      console.log(err)
    })
  }
};
</script>