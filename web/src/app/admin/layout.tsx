import { SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar";
import { AppSidebar } from "@/components/app-sidebar";

export default function Layout({ children }: { children: React.ReactNode }) {
    return (
        <SidebarProvider>
            <AppSidebar role="admin" />
            <main className="w-full p-2">
                <SidebarTrigger />
                <div className="p-2">{children}</div>
            </main>
        </SidebarProvider>
    );
}
