import { Suspense } from "react";
import { ArraySkeleton, TableRowSkeleton } from "@/components/skeleton";
import ProductTable from "@/components/admin/product.table";
import { getProducts } from "@/lib/api/admin";
import { PaginationWithLinks } from "@/components/ui/pagination-with-links";

export default async function AdminProductsPage({
    searchParams,
}: {
    searchParams: Promise<{
        page: string;
        limit: string;
        sort: string;
        search: string;
    }>;
}) {
    const page = parseInt((await searchParams).page || "1", 10);
    const limit = parseInt((await searchParams).limit || "5", 10);
    const search = (await searchParams).search || "";
    const sort = (await searchParams).sort || "desc";
    const { data: products, total } = await getProducts({
        page,
        limit,
        search,
        sort,
    });

    return (
        <>
            <h1 className="mb-4 text-2xl font-bold w-full">
                Admin Products Page
            </h1>

            <Suspense
                fallback={
                    <ArraySkeleton num={limit}>
                        <TableRowSkeleton />
                    </ArraySkeleton>
                }
            >
                <ProductTable products={products} />
            </Suspense>

            <PaginationWithLinks
                page={page}
                pageSize={limit}
                totalCount={total}
            />
        </>
    );
}
