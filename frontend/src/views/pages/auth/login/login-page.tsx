import { Card, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import loginBackground from "@/assets/login-background.jpg"
import { LoginPageForm } from "./login-page-form"
import { useEffect } from "react"
import { authCheck } from "@/lib/auth"
import type { AxiosResponse } from "axios"
import { useNavigate } from "react-router"
import routeCollection from "@/lib/route-collection"

export default function LoginPage() {

    // navigate
    const navigate = useNavigate()

    // use effect
    useEffect(() => {
        authCheck()
            .then((response: AxiosResponse | undefined) => {
                if (response) {
                    navigate(routeCollection.panel_dashboard)
                }
            })
            .catch(() => {
                console.log("account is logged out")
            })
    }, [navigate])

    // render
    return (
        <main role="main" className="w-full flex items-center max-w-[800px] max-sm:h-dvh max-sm:items-center max-sm:bg-white">
            <Card className="w-full py-10 md:max-w-[400px] max-h-[520px] md:rounded-r-none max-sm:rounded-none max-sm:border-none max-sm:shadow-none">
                <CardHeader className="text-center mb-6">
                    <CardTitle className="text-2xl">Masuk ke Akun Anda</CardTitle>
                    <CardDescription>Selamat Datang Kembali!</CardDescription>
                </CardHeader>
                <LoginPageForm />
            </Card>
            <div className="flex items-center w-full max-w-[400px] max-h-[520px] overflow-hidden rounded-r-xl shadow-sm max-md:hidden">
                <img className="h-full max-h-full w-full block object-cover object-top" src={loginBackground} alt="Login Background" />
            </div>
        </main>
    )
}