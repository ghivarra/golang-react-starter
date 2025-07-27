import { Sidebar, SidebarContent, SidebarFooter, SidebarHeader } from "../ui/sidebar"
import NavFooter from "./nav-footer"
import NavHeader from "./nav-header"
import NavMain from "./nav-main"

interface NavLayoutProps {
    collapsible?: "offcanvas" | "icon" | "none";
    variant?: "inset" | "sidebar" | "floating";
}

export default function NavLayout(props: NavLayoutProps) {

    // render
    return (
        <Sidebar {...props}>
            <SidebarHeader>
                <NavHeader />
            </SidebarHeader>
            <SidebarContent>
                <NavMain />
            </SidebarContent>
            <SidebarFooter>
                <NavFooter />
            </SidebarFooter>
        </Sidebar>
    )
}