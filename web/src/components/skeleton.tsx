import { Skeleton } from "@/components/ui/skeleton";
import {
    Card,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/components/ui/card";
import { Fragment } from "react";

interface ArraySkeletonProps {
    num: number;
    children: React.ReactNode;
}

export const ArraySkeleton = ({ num, children }: ArraySkeletonProps) => {
    return (
        <>
            {Array.from({ length: num }).map((_, i) => (
                <Fragment key={i}>{children}</Fragment>
            ))}
        </>
    );
};

export const TableRowSkeleton = () => {
    return <Skeleton className="h-6 w-full mb-2" />;
};

export const ProductCardSkeleton = () => {
    return (
        <Card className=" overflow-hidden hover:shadow-lg transition duration-300 ease-in-out">
            <Skeleton className="h-60 w-full" />
            <CardHeader>
                <CardTitle>
                    <Skeleton className="w-full h-full" />
                </CardTitle>
                <CardDescription>
                    <Skeleton className="w-full h-full" />
                </CardDescription>
            </CardHeader>
            <CardContent className="flex justify-between">
                <Skeleton className="w-1/2 h-12" />
                <Skeleton className="w-1/2 h-12" />
            </CardContent>
        </Card>
    );
};
