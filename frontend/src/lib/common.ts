import axios, { AxiosError, type AxiosRequestConfig, type AxiosResponse } from "axios"
import { toast } from "sonner"

export const deleteCookie = (name: string) => {
    setCookie(name, '', 0)
}

export const setCookie = (name: string, value: string, expHours: number) => {
    // set expiring time
    const date = new Date()
    const addedTime = (expHours === 0) ? -1 : (expHours*60*60*1000)
    date.setTime(date.getTime() + addedTime)
    const expires = "expires=" + date.toUTCString()

    // build string
    const cookieStr = `${name}=${value};${expires};domain=${location.hostname};samesite=strict;path=/`

    // set cookie
    document.cookie = cookieStr
}

export const getCookie = (name: string): string => {
  const cookieName = name + "=";
  const decodedCookie = decodeURIComponent(document.cookie);
  const ca = decodedCookie.split(';');

  for(let i = 0; i <ca.length; i++) {
    let c = ca[i];
    while (c.charAt(0) == ' ') {
      c = c.substring(1);
    }
    if (c.indexOf(cookieName) == 0) {
      return c.substring(cookieName.length, c.length);
    }
  }

  return "";
}

const prepareConfig = (withCredential: boolean, config: AxiosRequestConfig): AxiosRequestConfig => {
    // set more config
    if (typeof config.baseURL === 'undefined') {
        console.log('HITTT')
        config.baseURL = import.meta.env.VITE_API_BASE as string
    }

    // set bearer token
    if (withCredential) {
        if (typeof config.headers === 'undefined' || typeof config.headers.Authorization === 'undefined') {
            const token = getCookie('access_token')
            if (typeof config.headers === 'undefined') {
                config.headers = {
                    Authorization: `Bearer ${token}`
                }
            } else {
                config.headers.Authorization = `Bearer ${token}`
            }
        }
    }

    console.log(config)

    // return
    return config
}

export const fetchApi = async (withCredential: boolean, config: AxiosRequestConfig): Promise<AxiosResponse|void> => {
    const sentConfig = prepareConfig(withCredential, config)
    try {
        const response = await axios(sentConfig)
        return response
    } catch (err: unknown) {
        if (err instanceof Error) {
            console.warn(err)
        } else if (err instanceof AxiosError) {
            console.warn(err)
            if (typeof err.response?.data !== 'undefined') {
                toast.error(err.response.data.message)
            } else {
                toast.error(err.message)
            }
        } else {
            console.warn(err)
        }
        return
    }
}