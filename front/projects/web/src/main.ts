import "./styles/_all.scss"
import { createApp } from "vue"
import App from "./App.vue"
import i18n from "@/locales"
import router from "./router"
import useUserStore from "@/stores/use-user-store.ts"
import { init } from "@/app/app.ts"

const app = createApp(App).use(stores).use(i18n)

await init()

app.config.errorHandler = (err: any, vm: any, info: any) => {
  if (err.isAxiosError) {
    return
  }
  if (err.toString().includes("Validating error")) {
    return
  }
  console.error(err)
}

useUserStore()
  .init()
  .finally(() => {
    app.use(router).mount("#app")
  })
