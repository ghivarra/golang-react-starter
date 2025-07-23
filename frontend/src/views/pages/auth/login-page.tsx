import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { LogIn } from "lucide-react"
import { z } from "zod"
import { useForm } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form"
import { Input } from "@/components/ui/input"

// set form schema
const LoginFormSchema = z.object({
    username: z.string().min(1, { message: "Username tidak boleh kosong" }),
    password: z.string().min(1, { message: "Password tidak boleh kosong" })
})

export default function LoginPage() {

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
        <main role="main" className="w-full max-w-[380px]">
            <Card className="mb-4 w-full">
                <CardHeader className="text-center mb-4">
                    <CardTitle className="text-xl">Login</CardTitle>
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
                                        <FormLabel htmlFor="username" className="mb-2">Nama Pengguna</FormLabel>
                                        <FormControl className="mb-4">
                                            <Input id="username" placeholder="@username" required {...field} />
                                        </FormControl>
                                        <FormMessage />
                                    </FormItem>
                                )}
                            />
                            <FormField 
                                control={form.control}
                                name="password"
                                render={({field}) => (
                                    <FormItem className="mb-4">
                                        <FormLabel htmlFor="password" className="mb-2">Kata Sandi</FormLabel>
                                        <FormControl className="mb-4">
                                            <Input id="password" type="password" placeholder="********" required {...field} />
                                        </FormControl>
                                        <FormMessage />
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
                        Belum punya akun? <Button className="font-bold p-0" variant="link">Daftar di sini</Button>
                    </section>
                </CardContent>
            </Card>
            {/*<div className="text-muted-foreground *:[a]:hover:text-primary text-center text-xs text-balance *:[a]:underline *:[a]:underline-offset-4">
                By clicking continue, you agree to our <a href="#">Terms of Service</a>{" "}and <a href="#">Privacy Policy</a>.
            </div>*/}
        </main>
    )
}