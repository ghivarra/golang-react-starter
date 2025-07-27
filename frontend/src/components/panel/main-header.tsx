import { Separator } from "@radix-ui/react-separator"
import { SidebarTrigger } from "../ui/sidebar"

export default function MainHeader() {
    return (
        <header>
            <div className="flex">
                <SidebarTrigger />
                <Separator />
            </div>
        </header>
    )
}