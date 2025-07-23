import { Outlet } from "react-router";

export default function AuthView() {
    return (
        <>
            <div className="min-w-dvw min-h-dvh bg-muted flex items-center justify-center">
                <Outlet />
            </div>
        </>
    )
}