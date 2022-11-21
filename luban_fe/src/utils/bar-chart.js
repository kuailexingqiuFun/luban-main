import Color from 'color'
import merge from 'lodash/merge'
import moment from 'moment'
import { Line, mixins } from 'vue-chartjs'
import { ZebraStripes } from './zebra-stripes.plugin'

const getBarColor = ({ dataset }) => {
  const color = dataset.borderColor
  return Color(color)
    .alpha(0.2)
    .string()
}

const formatTimeLabels = (timestamp, index) => {
  const label = moment(parseInt(timestamp)).format('HH:mm')
  const offset = ' '
  if (index == 0) return offset + label
  if (index == 60) return label + offset
  return index % 10 == 0 ? label : ''
}

const { reactiveProp } = mixins

export default {
  extends: Line,
  mixins: [reactiveProp],
  props: {
    options: {
      type: Object,
      default: null
    },
    showLegend: {
      type: Boolean,
      default: true
    },
    height: {
      default: 300,
      type: Number
    },
    plugins: [ZebraStripes],
    timeLabelStep: 10
  },
  mounted() {
    // this.chartData 在 mixin 创建.
    // 如果你需要替换 options , 请创建本地的options对象;
    this.renderChart(this.chartData, this.optionsData)
  },
  computed: {
    optionsData() {
      const barOptions = {
        maintainAspectRatio: false,
        responsive: true,
        tooltips: {
          mode: 'index'
        },
        legend: {
          display: this.showLegend,
          position: 'bottom',
          labels: {}
        },
        scales: {
          xAxes: [
            {
              type: 'time',
              offset: true,
              gridLines: {
                display: false
              },
              stacked: true,
              ticks: {
                callback: formatTimeLabels,
                autoSkip: false,
                source: 'data',
                backdropColor: 'white',
                fontColor: '#424242',
                fontSize: 11,
                maxRotation: 0,
                minRotation: 0
              },
              bounds: 'data',
              time: {
                unit: 'minute',
                displayFormats: { minute: 'x' },
                parser: (timestamp) => moment.unix(parseInt(timestamp))
              }
            }
          ],
          yAxes: [
            {
              position: 'right',
              gridLines: {
                zeroLineWidth: 0
              },
              ticks: {
                maxTicksLimit: 6,
                fontColor: '#424242',
                fontSize: 11,
                padding: 8,
                min: 0
              }
            }
          ]
        },
        animation: {
          duration: 0
        },
        elements: {
          rectangle: {
            backgroundColor: getBarColor.bind(null)
          }
        },
        plugins: {
          ZebraStripes: {
            stripeColor: '#424242',
            interval: this.chartData.datasets[0].data.length
          }
        }
      }
      return merge(barOptions, this.options)
    }
  }
}
