import {
    Card,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import Image from "next/image";
import iphone16 from "@/public/iphone16.png";
import { memo } from "react";

interface ProductCardProps {
    name: string;
    image?: string;
    price: number;
}

export const ProductCard = memo(({ name, image, price }: ProductCardProps) => {
    return (
        <Card className=" overflow-hidden hover:shadow-lg transition duration-300 ease-in-out">
            <Image
                className="w-full h-60 object-cover"
                src={iphone16}
                alt={name}
                width={300}
                height={200}
            />
            <CardHeader>
                <CardTitle>{name}</CardTitle>
                <CardDescription>${price}</CardDescription>
            </CardHeader>
            <CardContent className="flex justify-between">
                <Button variant="outline">View details</Button>
                <Button variant="default">Add to cart</Button>
            </CardContent>
        </Card>
    );
});
