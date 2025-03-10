import {
  Calendar,
  CreditCard,
  Home,
  Inbox,
  Package,
  Search,
  Settings,
} from "lucide-react";
import Link from "next/link";
import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "@/components/ui/sidebar";
import Image from "next/image";
import hcmusLogo from "@/public/hcmus-logo.jpg";

// Menu items.
const userItems = [
  {
    title: "Home",
    url: "/",
    icon: Home,
  },
  {
    title: "Inbox",
    url: "#",
    icon: Inbox,
  },
  {
    title: "Calendar",
    url: "#",
    icon: Calendar,
  },
  {
    title: "Search",
    url: "#",
    icon: Search,
  },
  {
    title: "Settings",
    url: "#",
    icon: Settings,
  },
];

const adminItems = [
  {
    title: "Students",
    url: "/",
    icon: Home,
  },
  {
    title: "Categories",
    url: "/categories",
    icon: Package,
  },
  {
    title: "Logs",
    url: "/logs",
    icon: CreditCard,
  },
  {
    title: "Settings",
    url: "/settings",
    icon: Settings,
  },
];

interface AppSidebarProps {
  role: "admin" | "user";
}

export function AppSidebar({ role }: AppSidebarProps) {
  const isAdmin = role === "admin";
  const items = isAdmin ? adminItems : userItems;
  return (
    <Sidebar>
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupLabel>
            <div className="flex items-center gap-3 w-full">
              <Image src={hcmusLogo} alt="HCMUS Logo" width={35} height={35} />
              <p>Student Management</p>
            </div>
          </SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarMenu>
              {items.map((item) => (
                <SidebarMenuItem key={item.title}>
                  <SidebarMenuButton asChild>
                    <Link href={item.url}>
                      <item.icon />
                      <span>{item.title}</span>
                    </Link>
                  </SidebarMenuButton>
                </SidebarMenuItem>
              ))}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>
    </Sidebar>
  );
}
