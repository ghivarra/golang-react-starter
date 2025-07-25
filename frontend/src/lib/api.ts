import type { APIResponse } from "@/types"
import axios, { AxiosError, type AxiosInstance, type AxiosRequestConfig } from "axios"
import { toast } from "sonner"
import { getCookie, setCookie, sleep } from "./common"
import { authLogout } from "./auth"

// axios queue
const apiStatusKey = "api_status"


export const prepareAxios = (withCredential: boolean, config: AxiosRequestConfig): AxiosRequestConfig => {
    // set more config
    if (typeof config.baseURL === 'undefined') {
        config.baseURL = import.meta.env.VITE_API_BASE as string
    }

    // set bearer token
    if (withCredential) {
        if (typeof config.headers === 'undefined' || typeof config.headers.Authorization === 'undefined') {
            const token = getCookie(import.meta.env.VITE_ACCESS_COOKIE_NAME as string)
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
    if (!isApiClear()) {
        return false
    }

    // set to busy
    setApiStatus("busy")

    // get data
    const token = {
        access_token: getCookie(import.meta.env.VITE_ACCESS_COOKIE_NAME as string),
        refresh_token: getCookie(import.meta.env.VITE_REFRESH_COOKIE_NAME as string),
    }

    if (token.refresh_token == "" || token.access_token == "") {
        // set loggedout & return false
        authLogout()
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

    // set variable
    const accessCookieName = import.meta.env.VITE_ACCESS_COOKIE_NAME as string
    const refreshCookieName = import.meta.env.VITE_REFRESH_COOKIE_NAME as string

    // return promise
    try {

        // set
        const axiosInstance = await instance(axiosConfig)

        const res = axiosInstance.data as APIResponse
        const data = res.data as refreshResponseData

        // set cookie
        setCookie(accessCookieName, data.accessToken, import.meta.env.VITE_ACCESS_TOKEN_EXPIRATION as number)
        setCookie(refreshCookieName, data.refreshToken, import.meta.env.VITE_REFRESH_TOKEN_EXPIRATION as number)

        // status
        status = true

        // set clear
        setApiStatus("clear")

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

        // logout
        authLogout()
    }

    // return
    return status
}

export const isApiClear = (): boolean => {
    const value = localStorage.getItem(apiStatusKey)

    if (value === "") {
        localStorage.setItem(apiStatusKey, "clear")
        return true
    }

    return (value === "clear")
}

export const setApiStatus = (status: "clear"|"busy"|"loggedOut" = "clear") => {
    switch (status) {
        case "busy":
            localStorage.setItem(apiStatusKey, "busy")
            break;
        
        case "loggedOut":
            localStorage.setItem(apiStatusKey, "loggedOut")
            break;
    
        default:
            localStorage.setItem(apiStatusKey, "clear")
            break;
    }
}

export const fetchApi = async (withCredential: boolean, config: AxiosRequestConfig = {}, retryNumber: number = 0): Promise<AxiosInstance|void> => {

    // if retry number exceed x numbers then withdraw
    const retryMax = 50;
    if (retryNumber >= retryMax) {
        return
    }

    // check only if we use credential or access token
    if (withCredential) {
        if (localStorage.getItem(apiStatusKey) === "busy") {
            // sleep for 1 seconds then retry
            await sleep(1000)
            return fetchApi(withCredential, config, (retryNumber + 1))
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
                if (!error.message) {
                    console.error(error)
                    toast.error(error.message)
                } else {
                    const axiosErr = error as AxiosError

                    // check if request was unauthorized because access token is expired
                    if (axiosErr.response) {
                        if (axiosErr.response.status === 401) {

                            // send warn to console
                            console.warn("Unauthorized. Need to rotate the access token if authenticated.")

                            // refresh and rotate token
                            if (isApiClear()) {
                                const refreshed = await refreshToken()
                                if (refreshed && axiosErr.config) {
                                    // set new token
                                    const newToken = getCookie(import.meta.env.VITE_ACCESS_COOKIE_NAME)
                                    const newConfig = axiosErr.config
                                    newConfig.headers!.Authorization = `Bearer ${newToken}`
                                    return instance(newConfig)
                                }
                            } else {
                                console.log("Token refresh has been handled by another instance or you are already logged out")
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