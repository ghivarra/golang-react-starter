import type { AxiosError } from "axios"
import { fetchApi, setApiStatus } from "./api"
import { deleteCookie } from "./common"
import routeCollection from "./route-collection"

export const authCheck = async () => {
    // init axios
    const api = await fetchApi(true)

    if (!api) {
        return Promise.reject(undefined)
    }

    try {
        const response = api.get("/auth/check")
        return Promise.resolve(response)
    } catch(err: unknown) {
        const error = err as AxiosError
        return Promise.reject(error)
    }
}

export const authLogout = (routeName?: keyof typeof routeCollection) => {
    deleteCookie(import.meta.env.VITE_ACCESS_COOKIE_NAME as string)
    deleteCookie(import.meta.env.VITE_REFRESH_COOKIE_NAME as string)
    setApiStatus("loggedOut")

    if (routeName) {
        window.location.href = window.location.hostname + routeCollection[routeName]
    }
}