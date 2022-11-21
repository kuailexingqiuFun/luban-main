const interval = (
  timeSec = 1,
  callback,
  autoRun = false,
  immediately = false
) => {
  let count = 0
  let timer = null
  let isRunning = false

  const intervalManager = {
    start(runImmediately = false) {
      if (isRunning) return
      const tick = () => callback(++count)
      isRunning = true
      timer = setInterval(tick, 1000 * timeSec)
      if (runImmediately) tick()
    },

    stop() {
      count = 0
      isRunning = false
      clearInterval(timer)
    },

    restart(runImmediately = false) {
      this.stop()
      this.start(runImmediately)
    },

    isRunning() {
      return isRunning
    }
  }

  if (autoRun) intervalManager.start(immediately)

  return intervalManager
}

export default interval
