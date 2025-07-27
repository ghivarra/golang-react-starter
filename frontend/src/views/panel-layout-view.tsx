import NavLayout from "@/components/panel/nav-layout";
import { SidebarInset, SidebarProvider } from "@/components/ui/sidebar";
import { Outlet } from "react-router";
import { Toaster } from "sonner";

export default function PanelLayoutView() {
    return (
        <>
            <Toaster position="top-right" richColors />
            <SidebarProvider>
                <NavLayout collapsible="offcanvas" variant="inset" />
                <SidebarInset>
                    <Outlet />
                </SidebarInset>
            </SidebarProvider>
        </>
    )
}