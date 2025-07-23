import { Outlet } from "react-router";

export default function AuthView() {
    return (
        <>
            <div className="min-w-dvw min-h-dvh bg-linear-to-br from-green-900 to-cyan-800 flex items-center justify-center p-6 max-sm:p-0">
                <Outlet />
            </div>
        </>
    )
}