import { SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar";
import { AppSidebar } from "@/components/app-sidebar";
import { AppBreadcrumb } from "@/components/app-breadcrumb";

export default function UserLayout({
    children,
}: {
    children: React.ReactNode;
}) {
    return (
        <SidebarProvider>
            <AppSidebar role="user" />
            <main className="w-full p-2">
                <SidebarTrigger />
                <div className="p-2">
                    <AppBreadcrumb />
                    {children}
                </div>
            </main>
        </SidebarProvider>
    );
}
