import { Toaster } from "@/components/ui/toaster";
import * as Toast from "@radix-ui/react-toast";

export const ThemeProvider = ({ children }: { children: React.ReactNode }) => {
    return (
        <>
            <Toaster />
            <Toast.Provider duration={3000}>{children}</Toast.Provider>
        </>
    );
};
