"use client";
import Link from "next/link";
import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
    Form,
    FormControl,
    FormField,
    FormItem,
    FormLabel,
    FormMessage,
} from "@/components/ui/form";
import { Textarea } from "@/components/ui/textarea";
import { useRouter } from "next/navigation";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { signIn } from "next-auth/react"; // import từ next-auth

const SignInSchema = z.object({
    email: z.string().email(),
    password: z.string().min(6),
});

const SignInPage = () => {
    const [error, setError] = useState<string | null>(null);
    const router = useRouter();

    const form = useForm<z.infer<typeof SignInSchema>>({
        resolver: zodResolver(SignInSchema),
    });

    const onSubmit = async (values: z.infer<typeof SignInSchema>) => {
        // Simulate sign-in logic (replace with actual API call)
        if (
            values.email === "test@example.com" &&
            values.password === "password123"
        ) {
            // Redirect to dashboard after successful login
            router.push("/dashboard");
        } else {
            setError("Invalid email or password.");
        }
    };

    const handleOAuthSignIn = (provider: string) => {
        signIn(provider); // Hàm đăng nhập qua OAuth của next-auth
    };

    return (
        <div className="flex items-center justify-center min-h-screen bg-gray-100">
            <div className="bg-white p-8 rounded-lg shadow-md w-full max-w-sm">
                <h1 className="text-2xl font-semibold text-center mb-6">
                    Sign In
                </h1>
                <Form {...form}>
                    <form
                        onSubmit={form.handleSubmit(onSubmit)}
                        className="space-y-4"
                    >
                        <FormField
                            control={form.control}
                            name="email"
                            render={({ field }) => (
                                <FormItem>
                                    <FormLabel>Email</FormLabel>
                                    <FormControl>
                                        <Input placeholder="Email" {...field} />
                                    </FormControl>
                                    <FormMessage />
                                </FormItem>
                            )}
                        />

                        <FormField
                            control={form.control}
                            name="password"
                            render={({ field }) => (
                                <FormItem>
                                    <FormLabel>Password</FormLabel>
                                    <FormControl>
                                        <Input
                                            type="password"
                                            placeholder="Password"
                                            {...field}
                                        />
                                    </FormControl>
                                    <FormMessage />
                                </FormItem>
                            )}
                        />
                        {error && (
                            <p className="text-red-500 text-sm">{error}</p>
                        )}
                        <Button type="submit" className="w-full">
                            Sign In
                        </Button>
                    </form>
                </Form>

                

                <div className="mt-4 text-center space-y-2">
                    <Button
                        onClick={() => handleOAuthSignIn("google")}
                        className="w-full bg-red-500 hover:bg-red-600"
                    >
                        Sign in with Google
                    </Button>
                    <Button
                        onClick={() => handleOAuthSignIn("github")}
                        className="w-full bg-gray-800 hover:bg-gray-900"
                    >
                        Sign in with GitHub
                    </Button>
                </div>

                <div className="mt-4 text-center">
                    <p className="text-sm text-gray-500">
                        Don't have an account?
                    </p>
                    <a href="/signup" className="text-blue-500 hover:underline">
                        Sign up here
                    </a>
                </div>
            </div>
        </div>
    );
};

export default SignInPage;
