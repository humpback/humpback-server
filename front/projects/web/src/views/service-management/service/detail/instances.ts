export class DynamicIntervalTimer {
  private delay: number // 当前间隔时间
  private interval: number
  private timerId: NodeJS.Timeout | null
  private isRunning: boolean // 定时器是否正在运行

  constructor(initDelay?: number, interval?: number) {
    this.delay = initDelay || 3000
    this.interval = interval || 3000
    this.timerId = null
    this.isRunning = false
  }

  async start(callback: () => Promise<void>) {
    if (this.isRunning) return
    this.isRunning = true
    await this.run(callback)
  }

  reset(initDelay?: number, interval?: number) {
    this.stop()
    this.delay = initDelay || 3000
    this.interval = interval || 3000
  }

  stop() {
    if (this.timerId) {
      clearTimeout(this.timerId)
      this.timerId = null
    }
    this.isRunning = false
  }

  private async run(callback: () => Promise<void>) {
    if (!this.isRunning) return
    const delayTime = this.delay
    this.timerId = setTimeout(async () => {
      await callback().catch(() => {})
      this.delay = this.interval + delayTime
      await this.run(callback)
    }, delayTime)
  }
}
