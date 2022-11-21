<template>
  <div class="resource-metrics">
    <el-row justify="space-around">
      <el-col align="middle">
        <el-radio-group
            mandatory
            centered
            v-model="selection"
            size="small"
            active-class="primary--text">
          <el-radio-button  border @change="handleClick(idx)" v-for="(item, idx) in items" :key="item" :label="idx" :value="idx">
            {{ item }}
          </el-radio-button>
        </el-radio-group>
      </el-col>
    </el-row>
    <div v-if="metrics">
      <div v-for="(item, idx) in items" :key="item" style="width:100%;">
        <BarChart
            v-if="idx === selection"
            :options="options[idx]"
            :chart-data="getChartData(idx)"/>
      </div>
    </div>
    <div v-else class="no-metrics">
      <el-icon>mdi-information</el-icon>
      <span style="margin-left: 5px;">Metrics not available at the moment</span>
    </div>
  </div>
</template>
<script>
import BarChart from '@/utils/bar-chart'
import interval from '@/utils/interval'
import { mapValues } from 'lodash'
import { normalizeMetrics } from '@/utils/metrics'
import { bytesToUnits } from '@/utils/convertMemory'
import {GetMetrics} from '@/api/kubernetes/metrics'
const cpuOptions = {
  scales: {
    yAxes: [
      {
        ticks: {
          callback: (value) => {
            const float = parseFloat(`${value}`)
            if (float == 0) return '0'
            if (float < 10) return float.toFixed(3)
            if (float < 100) return float.toFixed(2)
            return float.toFixed(1)
          }
        }
      }
    ]
  },
  tooltips: {
    callbacks: {
      label: ({ datasetIndex, index }, { datasets }) => {
        const { label, data } = datasets[datasetIndex]
        const value = data[index]
        return `${label}: ${parseFloat(value.y).toPrecision(2)}`
      }
    }
  }
}
const memoryOptions = {
  scales: {
    yAxes: [
      {
        ticks: {
          callback: (value) => {
            if (typeof value == 'string') {
              const float = parseFloat(value)
              if (float < 1) {
                return float.toFixed(3)
              }
              return bytesToUnits(parseInt(value))
            }
            return bytesToUnits(value)
          },
          stepSize: 1
        }
      }
    ]
  },
  tooltips: {
    callbacks: {
      label: ({ datasetIndex, index }, { datasets }) => {
        const { label, data } = datasets[datasetIndex]
        const value = data[index]
        return `${label}: ${bytesToUnits(parseInt(value.y.toString()), 3)}`
      }
    }
  }
}
export default {
  name: 'MonitorBlock',
  props: {
    form: {
      type: Object,
      default() {
        return {}
      }
    }
  },
  components: {
    BarChart,
  },
  data() {
    return {
      poddetail: this.form,
      selection: 0,
      metrics: null,
      metricsWatcher: null,
      metricsData: this.metrics,
      items: ['CPU', 'Memory', 'Network', 'Filesystem']
    }
  },
  beforeDestroy() {
    this.metricsWatcher.stop()
  },
  created() {
    this.state = 1
    this.getMetrics()
    this.metricsWatcher = interval(60, () => {this.getMetrics()},
        true,
        true
    )
  },
  computed: {
    options() {
      return [cpuOptions, memoryOptions, memoryOptions, memoryOptions]
    },
    // 解析传入的指标数据属性
    chartDataSet() {
      return mapValues(
          this.metrics,
          (metric) => normalizeMetrics(metric).data.result[0].values
      )
    }
  },
  watch: {
    metrics(val) {
      this.metricsData = val
    }
  },
  methods: {
    submitForm() {
      this.$emit('submit', this.form)
    },
    getChartDataSet() {
      const {
        cpuUsage,
        cpuRequests,
        cpuLimits,
        memoryUsage,
        memoryRequests,
        memoryLimits,
        fsUsage,
        fsWrite,
        fsRead,
        networkReceive,
        networkTransmit,
      } = this.chartDataSet
      return [
        // CPU
        {
          datasets: [
            {
              id: `cpuUsage`,
              label: `Usage`,
              tooltip: `CPU cores usage`,
              borderColor: '#3D90CE',
              data: cpuUsage.map(([x, y]) => ({ x, y }))
            },
            {
              id: `cpuRequests`,
              label: `Requests`,
              tooltip: `CPU requests`,
              borderColor: '#30b24d',
              data: cpuRequests.map(([x, y]) => ({ x, y }))
            },
            {
              id: `cpuLimits`,
              label: `Limits`,
              tooltip: `CPU limits`,
              borderColor: '#909399',
              data: cpuLimits.map(([x, y]) => ({ x, y }))
            },
          ]
        },
        // Memory
        {
          datasets: [
            {
              id: `memoryUsage`,
              label: `Usage`,
              tooltip: `Memory usage`,
              borderColor: '#c93dce',
              data: memoryUsage.map(([x, y]) => ({ x, y }))
            },
            {
              id: 'memoryRequests',
              label: `Requests`,
              tooltip: `Memory requests`,
              borderColor: '#30b24d',
              data: memoryRequests.map(([x, y]) => ({ x, y }))
            },
            {
              id: 'memoryLimits',
              label: `Limits`,
              tooltip: `Memory limits`,
              borderColor: '#909399',
              data: memoryLimits.map(([x, y]) => ({ x, y }))
            },
          ]
        },
        // Network
        {
          datasets: [
            {
              id: `networkReceive`,
              label: `Receive`,
              tooltip: `Bytes received by all containers`,
              borderColor: "#64c5d6",
              data: networkReceive.map(([x, y]) => ({ x, y })),
            },
            {
              id: `networkTransmit`,
              label: `Transmit`,
              tooltip: `Bytes transmitted from all containers`,
              borderColor: "#46cd9e",
              data: networkTransmit.map(([x, y]) => ({ x, y })),
            },
          ]
        },
        // Filesystem
        {
          datasets: [
            {
              id: `fsUsage`,
              label: `Usage`,
              tooltip: `Bytes consumed on this filesystem`,
              borderColor: "#ffc63d",
              data: fsUsage.map(([x, y]) => ({ x, y })),
            },
            {
              id: `fsWrites`,
              label: `Writes`,
              tooltip: `Bytes written on this filesystem`,
              borderColor: "#ff963d",
              data: fsWrite.map(([x, y]) => ({ x, y })),
            },
            {
              id: `fsReads`,
              label: `Reads`,
              tooltip: `Bytes read on this filesystem`,
              borderColor: "#fff73d",
              data: fsRead.map(([x, y]) => ({ x, y })),
            },
          ]
        },
      ]
    },
    getChartData(index) {
      const dataItem = this.getChartDataSet()[index]
      return Object.assign(dataItem, {
        datasets: dataItem.datasets
            .filter((set) => set.data.length)
            .map((item) => {
              return Object.assign(
                  {
                    type: 'bar',
                    borderWidth: { top: 3 },
                    barPercentage: 1,
                    categoryPercentage: 1
                  },
                  item
              )
            })
      })
    },
    async getMetrics() {
      const opts = { category: 'pods',
        pods: this.poddetail.metadata.name,
        namespace: this.poddetail.metadata.namespace,
        cluster_id: this.poddetail.cluster_id,
      }
      const resp = await GetMetrics({
        CpuUsage: opts,
        CpuRequests: opts,
        CpuLimits: opts,
        memoryUsage: opts,
        MemoryRequests: opts,
        MemoryLimits: opts,
        FsUsage: opts,
        FsWrite: opts,
        FsRead: opts,
        NetworkReceive: opts,
        NetworkTransmit: opts,
      })
      this.metrics = resp.data
    },
    handleClick(value){
      this.selection = value
    },
    cancel() {
      this.$emit('cancel')
    }
  }
}
</script>
