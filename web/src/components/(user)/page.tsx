"use client";

import { useEffect, useRef, useState, useCallback, use } from "react";
import { getProducts } from "@/lib/api/admin";
import { Product } from "@/lib/model/product";
import { ProductCard } from "@/components/homepage/product-card";
import { ArraySkeleton, ProductCardSkeleton } from "@/components/skeleton";
import { text } from "stream/consumers";

export default function HomePage() {
    const [products, setProducts] = useState<Product[]>([]);
    const [total, setTotal] = useState<number>(0);
    const [page, setPage] = useState<number>(1);
    const [hasMore, setHasMore] = useState<boolean>(true);
    const [loading, setLoading] = useState<boolean>(false);
    const observer = useRef<IntersectionObserver | null>(null);

    const fetchProducts = async () => {
        if (loading || !hasMore) return;

        setLoading(true);
        try {
            const { data, total } = await getProducts({
                page,
                limit: 10,
                search: "",
                sort: "desc",
            });
            setProducts((prevProducts) => [...prevProducts, ...data]);
            setHasMore(data.length > 0);
            setTotal(total);
        } catch (error) {
            console.error("Failed to fetch products:", error);
        } finally {
            setLoading(false);
        }
    };

    const lastElementRef = useCallback(
        (node: HTMLDivElement | null) => {
            if (loading) return;
            if (observer.current) observer.current.disconnect();

            observer.current = new IntersectionObserver((entries) => {
                if (entries[0].isIntersecting && hasMore) {
                    setPage((prevPage) => prevPage + 1);
                }
            });

            if (node) observer.current.observe(node);
        },
        [loading, hasMore],
    );

    useEffect(() => {
        fetchProducts();
    }, [page]);

    return (
        <div className="p-4 grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
            {products.map((product) => (
                <ProductCard
                    key={product.id}
                    name={product.name}
                    // image={product.images[0] || "/placeholder.png"}
                    price={product.price}
                />
            ))}

            {loading && (
                <ArraySkeleton num={Math.min(10, total - products.length)}>
                    <ProductCardSkeleton />
                </ArraySkeleton>
            )}

            {!hasMore && !loading && (
                <p className="col-span-full text-center mt-4">
                    No more products
                </p>
            )}

            <div ref={lastElementRef} className="h-10 col-span-full"></div>
        </div>
    );
}
