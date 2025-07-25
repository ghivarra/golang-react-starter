import { configureStore } from "@reduxjs/toolkit";
import loaderSlice from "../slices/loader-slices";

// create store
const store = configureStore({
    reducer: {
        loader: loaderSlice.reducer
    }
})

// export loader store
export default store