import { createSlice } from "@reduxjs/toolkit"

// create and export slices
const loaderSlice = createSlice({
    name: "loader",
    initialState: {
        show: false
    },
    reducers: {
        hideLoader: (state) => {
            state.show = false
        },
        showLoader: (state) => {
            state.show = true
        },
    }
})

// export actions
export const { hideLoader, showLoader } = loaderSlice.actions

// export slices
export default loaderSlice