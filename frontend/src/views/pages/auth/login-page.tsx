import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { LogIn } from "lucide-react"
import { z } from "zod"
import { useForm } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import { Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import loginBackground from "@/assets/login-background.jpg"
import { useNavigate } from "react-router"
import CustomDynamicIcon from "@/components/custom-dynamic-icon"
import { useState } from "react"
import type { IconCollection } from "@/lib/icon-collection"

// set form schema
const LoginFormSchema = z.object({
    username: z.string().min(1, { message: "Username tidak boleh kosong" }),
    password: z.string().min(1, { message: "Password tidak boleh kosong" })
})

export default function LoginPage() {

    // reactive
    const [passwordIcon, setPasswordIcon] = useState<keyof typeof IconCollection>("Eye")
    const [passwordInputType, setPasswordInputType] = useState<"text"|"password">("password")
    const [togglePasswordTitle, setTogglePasswordTitle] = useState("Lihat kata sandi")

    // navigate
    const navigate = useNavigate()

    // to register page
    const toRegisterPage = () => {
        navigate("/user/register")
    }

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
    const sendForm = () => {

        console.log(form.getValues())
    }

    return (
        <main role="main" className="w-full flex max-w-[800px] max-sm:h-dvh max-sm:items-center max-sm:bg-white">
            <Card className="w-full py-10 md:max-w-[400px] md:rounded-r-none max-sm:rounded-none max-sm:border-none max-sm:shadow-none">
                <CardHeader className="text-center mb-6">
                    <CardTitle className="text-2xl">Masuk ke Akun Anda</CardTitle>
                    <CardDescription>Selamat Datang Kembali!</CardDescription>
                </CardHeader>
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
                                        <FormDescription className="mb-2">Gunakan hanya nama pengguna yang sudah terdaftar</FormDescription>
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
                                            <FormDescription className="mb-2">Kata sandi yang sesuai dengan nama pengguna</FormDescription>
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
                        Belum punya akun? <Button onClick={toRegisterPage} className="font-bold p-0" variant="link">Daftar di sini</Button>
                    </section>
                </CardContent>
            </Card>
            <div className="w-full max-w-[400px] overflow-hidden rounded-r-xl shadow-sm max-md:hidden">
                <img className="h-full w-full object-cover object-center" src={loginBackground} alt="Login Background" />
            </div>
        </main>
    )
}