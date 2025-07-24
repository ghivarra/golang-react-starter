import type { AxiosError } from "axios"
import { fetchApi } from "./common"

export const authCheck = async () => {
    // init axios
    const api = await fetchApi(true)

    if (!api) {
        return
    }

    try {
        const response = await api.get("/auth/check")
        return Promise.resolve(response)
    } catch(err: unknown) {
        const error = err as AxiosError
        return Promise.reject(error)
    }
}