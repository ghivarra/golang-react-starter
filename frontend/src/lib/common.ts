import type { APIResponse } from "@/types"
import axios, { AxiosError, type AxiosInstance, type AxiosRequestConfig } from "axios"
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

export const sleep = (ms: number) => {
    return new Promise(resolve => setTimeout(resolve, ms))
}

const prepareAxios = (withCredential: boolean, config: AxiosRequestConfig): AxiosRequestConfig => {
    // set more config
    if (typeof config.baseURL === 'undefined') {
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

    // return
    return config
}

export const refreshToken = async (): Promise<boolean> => {

    // check
    if (!checkAxiosClear()) {
        return false
    }

    // set to busy
    setAxiosStatus("busy")

    // get data
    const token = {
        access_token: getCookie("access_token"),
        refresh_token: getCookie("refresh_token"),
    }

    if (token.refresh_token == "" || token.access_token == "") {
        // set loggedout
        setAxiosStatus("loggedOut")
        return false
    }

    // new axios instance
    const axiosConfig = prepareAxios(false, {
        url: "/auth/refresh-token",
        method: "POST",
        data: token,
    })

    // refresh token
    const instance = axios.create()

    // refresh interface
    interface refreshResponseData {
        accessToken: string;
        refreshToken: string;
    }

    let status: boolean

    // return promise
    try {

        // set
        const axiosInstance = await instance(axiosConfig)

        const res = axiosInstance.data as APIResponse
        const data = res.data as refreshResponseData

        // set cookie
        setCookie("access_token", data.accessToken, import.meta.env.ACCESS_TOKEN_EXPIRATION as number)
        setCookie("refresh_token", data.refreshToken, import.meta.env.REFRESH_TOKEN_EXPIRATION as number)

        // status
        status = true

        // set clear
        setAxiosStatus("clear")

    } catch (err: unknown) {

        const error = err as AxiosError
        
        if (error.response) {
            if (error.response.data) {
                const res = error.response.data as APIResponse
                console.warn(res.message)
                toast.warning("Sesi anda sudah kedaluwarsa atau anda belum login")
            } else {
                console.warn(error.message)
                toast.warning("Sesi anda sudah kedaluwarsa atau anda belum login")
            }
        } else {
            const message = "Failed to rotate token, check your internet connection" 
            console.error(message)
            toast.error(message)
        }

        // set
        status = false

        // set loggedout
        setAxiosStatus("loggedOut")
    }

    // return
    return status
}

export const checkAxiosClear = (): boolean => {
    // check storage
    const xhrStatus = localStorage.getItem("xhr_status")

    // check and create local storage
    if (xhrStatus !== "clear" && xhrStatus !== "busy") {
        localStorage.setItem("xhr_status", "clear")
        return true
    }

    // if xhr status == clear then go through
    return (xhrStatus === "clear")
}

export const setAxiosStatus = (status: "clear"|"busy"|"loggedOut") => {
    localStorage.setItem("xhr_status", status)
}

export const fetchApi = async (withCredential: boolean, config: AxiosRequestConfig = {}, retryCount: number = 0): Promise<AxiosInstance|void> => {

    if (localStorage.getItem("xhr_status") === "loggedOut") {
        setAxiosStatus("clear")
        return
    }

    // set max retry to 10
    if (retryCount > 9) {
        setAxiosStatus("clear")
        return
    }

    // check only if we use credential or access token
    if (withCredential) {
        // if axios is not ready/clear yet
        // wait for 1 seconds, then try again
        if (!checkAxiosClear()) {
            console.info("Axios waiting for axios status to be cleared...")
            await sleep(1000)
            return fetchApi(withCredential, config, (retryCount + 1))
        }
    }

    // send config
    const sentConfig = prepareAxios(withCredential, config)
    
    // create axios
    const instance = axios.create(sentConfig)

    // set interceptors
    instance.interceptors.response.use(
        response => response,        
        async error => {

            // check error
            if (error instanceof Error) {
                if (typeof error.message === 'undefined') {
                    console.error(error)
                    toast.error("Koneksi gagal, cek internet anda")
                } else {
                    const axiosErr = error as AxiosError

                    // check if request was unauthorized because access token is expired
                    if (typeof axiosErr.response !== "undefined") {
                        if (axiosErr.response.status === 401) {
                            // send warn to console
                            console.warn("Unauthorized. Need to rotate the access token.")

                            // refresh and rotate token
                            // retry again after waiting 0.25 second again
                            const refresh = await refreshToken()

                            // set now
                            if (refresh && axiosErr.config) {
                                await sleep(250)
                                instance(axiosErr.config)
                            }

                            if (!refresh) {
                                fetchApi(withCredential, axiosErr.config, 1)
                            }
                        }
                    }
                }
                
            }  else {
                console.warn(error)
            }

            // return and reject with error
            return Promise.reject(error)
        }
    )

    // return axios
    return instance
}