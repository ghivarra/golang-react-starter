import type { AxiosError } from "axios"
import { fetchApi } from "./api"

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