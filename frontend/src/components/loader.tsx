import store from "@/states/stores"
import { useState } from "react"

export default function Loader() {

    // set class
    const [ className, setClassName ] = useState("hidden opacity-0 transition-opacity duration-200 linear fixed left-0 top-0 min-h-dvh min-w-dvw bg-white z-50 flex items-center justify-center")

    // subscribe to store
    store.subscribe(() => {
        const state = store.getState()
        const hiddenClass = "hidden "
        const opacityZeroClass = "opacity-0 "
        const opacityFullClass = "opacity-95 "

        if (state.loader.show) {
            let newClassName = className.replaceAll(hiddenClass, "")
            setClassName(newClassName)
            setTimeout(() => {
                newClassName = newClassName.replaceAll(opacityZeroClass, opacityFullClass)
                setClassName(newClassName)
            }, 50)
        } else {
            setClassName(opacityZeroClass + className)
            setTimeout(() => {
                setClassName(hiddenClass + opacityZeroClass + className)
            }, 250)
        }
    })

    // render
    return (
        <div className={className}>
            <div className="loader"></div>
        </div>
    )
}