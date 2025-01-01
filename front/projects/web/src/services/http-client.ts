import axios, { type AxiosRequestConfig, type AxiosResponse, type Method } from "axios"
import { isEmpty } from "lodash"
import useUserStore from "@/stores/use-user-store.ts"
import { GetCurrentLocale, GetI18nMessage } from "@/locales"
import { disposeStore } from "@/stores"
import globalLoading from "@/utils/loading.ts"

export interface HttpRequestOptions extends AxiosRequestConfig {
  disableLoading?: boolean | false
  disableErrMsg?: boolean | false
  loadingMessage?: string
  isFile?: boolean
  timeout?: number
}

const _loadingCount = ref(0)
const loading = computed({
  get() {
    return _loadingCount.value > 0
  },
  set(val: boolean) {
    _loadingCount.value += val ? 1 : -1
    _loadingCount.value = Math.max(0, _loadingCount.value)
  }
})

class HttpClientService {
  public async get<T>(url: string, options?: HttpRequestOptions): Promise<AxiosResponse<T>> {
    return this.request("GET", url, null, options)
  }

  public async put<T>(url: string, data: any, options?: HttpRequestOptions): Promise<AxiosResponse<T>> {
    return this.request("PUT", url, data, options)
  }

  public async post<T>(url: string, data: any, options?: HttpRequestOptions): Promise<AxiosResponse<T>> {
    return this.request("POST", url, data, options)
  }

  public async delete<T>(url: string, options?: HttpRequestOptions): Promise<AxiosResponse<T>> {
    return this.request("DELETE", url, null, options)
  }

  private handleInnerErr: any = (options?: HttpRequestOptions) => (err: any) => {
    console.error(err)
    if (err.response) {
      let body = err.response.data
      if (!isEmpty(body)) {
        if (body.statusCode === 401) {
          const userStore = useUserStore()
          switch (body.code) {
            case "R40101":
              ShowErrMsg(body.errMsg)
              userStore.clearUserInfo()
              disposeStore()
              window.location.href = "/pub/sign-in"
              break
            case "R40102":
              userStore.clearUserInfo()
              disposeStore()
              console.error(body.errMsg)
              break
            default:
              ShowErrMsg(body.errMsg)
          }
          throw err
        }

        if (!options || !options.disableErrMsg) {
          if (body.statusCode >= 500) {
            console.error(body.errMsg)
            ShowSystemErrMsg()
          } else {
            ShowErrMsg(body.errMsg)
          }
        }
      } else {
        ShowSystemErrMsg()
      }
    } else if (err.message && err.name === "AxiosError") {
      const message = err.message as string
      if (message.toLowerCase().includes("timeout")) {
        ShowErrMsg(GetI18nMessage("err.timeout"))
      } else {
        ShowErrMsg(message)
      }
    } else {
      ShowSystemErrMsg()
    }
    throw err
  }

  private async request<T>(method: Method, url: string, data: any, options?: HttpRequestOptions): Promise<AxiosResponse<T>> {
    if (!options) {
      options = {} as HttpRequestOptions
    }

    const headers: any = {
      "Accept": "application/json, text/javascript, */*",
      "Accept-Language": GetCurrentLocale(),
      "Content-Language": GetCurrentLocale()
    }

    if (options.isFile) {
      headers["Content-Type"] = "multipart/form-data"
    } else if (method !== "GET" && method !== "DELETE") {
      headers["Content-Type"] = "application/json"
    }

    if (options && options.headers && typeof options.headers === "object") {
      for (const key in options.headers) {
        headers[key] = options.headers[key]
      }
    }

    if (!options || !options.disableLoading) {
      loading.value = true
      if (!loading.value) {
        globalLoading.show(options.loadingMessage)
      }
    }

    return await axios
      .request<T>({
        url: url,
        method: method,
        data: data,
        headers: headers,
        responseType: options.responseType,
        params: Object.assign({}, options.params, { _t: new Date().valueOf() }),
        timeout: options.timeout || 60000
      })
      .catch<AxiosResponse<T>>(this.handleInnerErr(options))
      .finally(() => {
        if (!options || !options.disableLoading) {
          loading.value = false
          if (!loading.value) {
            globalLoading.close()
          }
        }
      })
  }
}

export const httpClient = new HttpClientService()
