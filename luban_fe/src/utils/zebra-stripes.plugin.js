// Plugin for drawing stripe bars on top of any timeseries barchart
// Based on cover DIV element with repeating-linear-gradient style
import get from 'lodash/get'
import moment from 'moment'
const defaultOptions = {
  stripeColor: '#ffffff08'
}
export const ZebraStripes = {
  updated: null,
  options: {},
  getOptions(chart) {
    return get(chart, 'options.plugins.ZebraStripes')
  },
  getLastUpdate(chart) {
    const data = chart.data.datasets[0].data[0]
    return moment.unix(parseInt(data.x))
  },
  getStripesElem(chart) {
    return chart.canvas.parentElement.querySelector('.zebra-cover')
  },
  removeStripesElem(chart) {
    const elem = this.getStripesElem(chart)
    if (!elem) return
    chart.canvas.parentElement.removeChild(elem)
  },
  updateOptions(chart) {
    this.options = Object.assign(
      Object.assign({}, defaultOptions),
      this.getOptions(chart)
    )
  },
  getStripeMinutes() {
    return this.options.interval < 10 ? 0 : 10
  },
  renderStripes(chart) {
    if (!chart.data.datasets.length) return
    const { interval, stripeColor } = this.options
    const { top, left, bottom, right } = chart.chartArea
    const step = (right - left) / interval
    const stripeWidth = step * this.getStripeMinutes()
    const cover = document.createElement('div')
    const styles = cover.style
    if (this.getStripesElem(chart)) return
    cover.className = 'zebra-cover'
    styles.width = `${right - left}px`
    styles.left = `${left}px`
    styles.top = `${top}px`
    styles.height = `${bottom - top}px`
    styles.backgroundImage = `
      repeating-linear-gradient(to right, ${stripeColor} 0px, ${stripeColor} ${stripeWidth}px,
      transparent ${stripeWidth}px, transparent ${stripeWidth * 2 + step}px)
     `
    chart.canvas.parentElement.appendChild(cover)
  },
  afterInit(chart) {
    if (!chart.data.datasets.length) return
    this.updateOptions(chart)
    this.updated = this.getLastUpdate(chart)
  },
  afterUpdate(chart) {
    this.updateOptions(chart)
    this.renderStripes(chart)
  },
  resize(chart) {
    this.removeStripesElem(chart)
  },
  afterDatasetUpdate(chart) {
    if (!this.updated) this.updated = this.getLastUpdate(chart)
    const { interval } = this.options
    const { left, right } = chart.chartArea
    const step = (right - left) / interval
    const diff = moment(this.updated).diff(this.getLastUpdate(chart), 'minutes')
    const minutes = Math.abs(diff)
    this.removeStripesElem(chart)
    this.renderStripes(chart)
    if (minutes > 0) {
      // Move position regarding to difference in time
      const cover = this.getStripesElem(chart)
      cover.style.backgroundPositionX = `${-step * minutes}px`
    }
  }
}
