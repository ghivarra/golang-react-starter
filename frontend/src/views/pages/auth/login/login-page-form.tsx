import { CardContent } from "@/components/ui/card"
import { useNavigate } from "react-router"
import CustomDynamicIcon from "@/components/custom-dynamic-icon"
import { useState } from "react"
import type { IconCollection } from "@/lib/icon-collection"
import { LogIn } from "lucide-react"
import { z } from "zod"
import { useForm } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import { Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import routeCollection from "@/lib/route-collection"
import { setCookie } from "@/lib/common"
import type { AxiosError, AxiosResponse } from "axios"
import { toast } from "sonner"
import type { APIResponse, UserData } from "@/types"
import CustomErrorAlert from "@/components/custom-error-alert"
import { fetchApi, setApiStatus } from "@/lib/api"

// set form schema
const LoginFormSchema = z.object({
    username: z.string().min(1, { message: "Username tidak boleh kosong" }),
    password: z.string().min(1, { message: "Password tidak boleh kosong" })
})

interface ErrorInterface {
    username?: string;
    password?: string;
}

interface LoginSuccessData {
    user: UserData;
    accessToken: string;
    refreshToken: string;
}

export function LoginPageForm() {

    // reactive
    const [passwordIcon, setPasswordIcon] = useState<keyof typeof IconCollection>("Eye")
    const [passwordInputType, setPasswordInputType] = useState<"text"|"password">("password")
    const [togglePasswordTitle, setTogglePasswordTitle] = useState("Lihat kata sandi")

    // errors
    const [ error, setError ] = useState<ErrorInterface>({})

    // navigate
    const navigate = useNavigate()

    // toggle password input
    const togglePasswordInput = () => {
        if (passwordInputType === "password") {
            setPasswordInputType("text")
            setPasswordIcon("EyeClosed")
            setTogglePasswordTitle("Sembunyikan kata sandi")
        } else {
            setPasswordInputType("password")
            setPasswordIcon("Eye")
            setTogglePasswordTitle("Lihat kata sandi")
        }
    }

    // init form
    const form = useForm<z.infer<typeof LoginFormSchema>>({
        resolver: zodResolver(LoginFormSchema),
        defaultValues: {
            username: "",
            password: ""
        }
    })

    // init send form
    const sendForm = async () => {

        const input = form.getValues()
        const api = await fetchApi(false)

        if (!api) {
            return
        }

        api.post("auth/authenticate", {
            username: input.username,
            password: input.password
        }).then((response: AxiosResponse) => {
            const res = response.data as APIResponse
            const data = res.data as LoginSuccessData

            // put into cookie
            const refreshName = import.meta.env.VITE_REFRESH_COOKIE_NAME as string
            const refreshTime = import.meta.env.VITE_REFRESH_TOKEN_EXPIRATION as number
            const accessName = import.meta.env.VITE_ACCESS_COOKIE_NAME as string
            const accessTime = import.meta.env.VITE_ACCESS_TOKEN_EXPIRATION as number

            // set cookie
            setCookie(accessName, data.accessToken, accessTime)
            setCookie(refreshName, data.refreshToken, refreshTime)

            // set api clear
            setApiStatus("clear")

            // login
            navigate(routeCollection.panel_dashboard)

        }).catch((err) => {
            const error = err as AxiosError
            if (error.response && error.response.data) {
                const data = error.response.data as APIResponse
                toast.error(data.message)

                if (data.errors) {

                    // set new error
                    const newError: ErrorInterface = JSON.parse(JSON.stringify(error))

                    // input error
                    Object.keys(data.errors).forEach((key) => {
                       const field = key as "username"|"password"
                       newError[field] = data.errors![field].join(" ")
                    })

                    // set error
                    setError(newError)
                }
            } else {
                toast.error(error.message)
            }
        })
    }

    return (
        <CardContent>
            <section className="mb-6">
                <Form {...form}>
                    <FormField 
                        control={form.control}
                        name="username"
                        render={({field}) => (
                            <FormItem className="mb-4">
                                <FormLabel htmlFor="username" className="mb-2 font-bold text-gray-800">Nama Pengguna</FormLabel>
                                
                                <FormControl className="">
                                    <Input id="username" placeholder="@username" required {...field} />
                                </FormControl>
                                {
                                    (error.username && error.username.length > 0) ? 
                                    <CustomErrorAlert className="my-2" message={error.username} /> :
                                    <FormDescription className="mb-2">Gunakan hanya nama pengguna yang sudah terdaftar</FormDescription>
                                }
                                <FormMessage className="mb-4" />
                            </FormItem>
                        )}
                    />
                    <FormField 
                        control={form.control}
                        name="password"
                        render={({field}) => (
                            <FormItem className="mb-4">
                                <FormLabel htmlFor="password" className="mb-2 font-bold text-gray-800">Kata Sandi</FormLabel>
                                <div className="relative">
                                    <FormControl className="mb-4">
                                        <Input id="password" type={passwordInputType} placeholder="********" required {...field} />
                                    </FormControl>
                                    {
                                        (error.password && error.password.length > 0) ? 
                                        <CustomErrorAlert className="my-2" message={error.password} /> :
                                        <FormDescription className="mb-2">Kata sandi yang sesuai dengan nama pengguna</FormDescription>
                                    }
                                    <FormMessage className="mb-4" />
                                    <button onClick={togglePasswordInput} title={togglePasswordTitle} type="button" className="p-0 absolute top-2 right-4">
                                        <CustomDynamicIcon name={passwordIcon} />
                                    </button>
                                </div>
                                
                            </FormItem>
                        )}
                    />
                </Form>
            </section>
            <section>
                <div className="mb-2">
                    <Button onClick={sendForm} type="button">
                        <LogIn className="text-white mr-1" />
                        Login
                    </Button>
                </div>
                Belum punya akun? <Button onClick={() => { navigate(routeCollection.user_register) }} className="font-bold p-0" variant="link">Daftar di sini</Button>
            </section>
        </CardContent>
    )
}