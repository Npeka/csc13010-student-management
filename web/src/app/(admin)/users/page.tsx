import { Suspense } from "react";
import { ArraySkeleton, TableRowSkeleton } from "@/components/skeleton";
import UserTable from "@/components/admin/user-table";

export default async function AdminUsersPage() {
    return (
        <>
            <h1 className="mb-4 text-2xl font-bold w-full">Admin Users Page</h1>
            <Suspense
                fallback={
                    <ArraySkeleton num={5}>
                        <TableRowSkeleton />
                    </ArraySkeleton>
                }
            >
                <UserTable />
            </Suspense>
        </>
    );
}
