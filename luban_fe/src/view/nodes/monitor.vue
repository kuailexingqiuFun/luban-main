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
const podOptions = {
  scales: {
    yAxes: [{
      ticks: {
        callback: value => value,
      },
    }],
  },
  tooltips: {
    callbacks: {
      label: ({ datasetIndex, index }, { datasets }) => {
        const { label, data } = datasets[datasetIndex];
        const value =data[index];

        return `${label}: ${value.y}`;
      },
    },
  },
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
      nodedetail: this.form,
      selection: 0,
      metrics: null,
      metricsWatcher: null,
      metricsData: this.metrics,
      selectTime: {},
      items: ['CPU', 'Memory', 'Disk', 'Pods']
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
      return [cpuOptions, memoryOptions, memoryOptions, podOptions]
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
        memoryRequests,
        memoryUsage,
        memoryCapacity,
        cpuUsage,
        cpuRequests,
        cpuCapacity,
        fsSize,
        fsUsage,
        podUsage,
        podCapacity,
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
              id: `cpuCapacity`,
              label: `Capacity`,
              tooltip: `CPU Capacity`,
              borderColor: '#30b24d',
              data: cpuCapacity.map(([x, y]) => ({ x, y }))
            },
            {
              id: `cpuRequests`,
              label: `Requests`,
              tooltip: `CPU requests`,
              borderColor: "#909399",
              data: cpuRequests.map(([x, y]) => ({ x, y })),
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
              id: 'MemoryCapacity',
              label: `Capacity`,
              tooltip: `Memory Capacity`,
              borderColor: '#909399',
              data: memoryCapacity.map(([x, y]) => ({ x, y }))
            },
            {
              id: "memoryRequests",
              label: `Requests`,
              tooltip: `Memory requests`,
              borderColor: "#30b24d",
              data: memoryRequests.map(([x, y]) => ({ x, y })),
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
              id: `fsSize`,
              label: `Size`,
              tooltip: `Node filesystem size in bytes`,
              borderColor: "#9cd3ce",
              data: fsSize.map(([x, y]) => ({ x, y })),
            },
          ]
        },
        // Pods
        {
          datasets: [
            {
              id: `podUsage`,
              label: `Usage`,
              tooltip: `Number of running Pods`,
              borderColor: "#30b24d",
              data: podUsage.map(([x, y]) => ({ x, y })),
            },
            {
              id: `podCapacity`,
              label: `Capacity`,
              tooltip: `Node Pods capacity`,
              borderColor: "#ffc63d",
              data: podCapacity.map(([x, y]) => ({ x, y })),
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
      const opts = { category: 'nodes',
        nodes: this.nodedetail.metadata.name,
        namespace: this.nodedetail.metadata.namespace,
        cluster_id: this.nodedetail.cluster_id,
        start: this.selectTime.start,
        end: this.selectTime.end,
      }
      const resp = await GetMetrics({
        MemoryRequests: opts,
        MemoryUsage: opts,
        MemoryCapacity: opts,
        CpuUsage: opts,
        CpuRequests: opts,
        CpuCapacity: opts,
        FsSize: opts,
        FsUsage: opts,
        PodUsage: opts,
        PodCapacity: opts,
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
