"use client";
import { useState } from "react";
import { Button } from "@/components/ui/button";
import {
    Dialog,
    DialogContent,
    DialogHeader,
    DialogFooter,
    DialogTitle,
    DialogDescription,
} from "@/components/ui/dialog";
import { useToast } from "@/hooks/use-toast";
import { Loader2 } from "lucide-react";
import { deleteUser, deleteProduct } from "@/lib/api/admin";

export const DeleteUser = ({ id }: { id: string }) => {
    const { toast } = useToast();
    const [isDialogOpen, setIsDialogOpen] = useState<boolean>(false);
    const [loadingDelete, setLoadingDelete] = useState<boolean>(false);
    const deleteUserById = deleteUser.bind(null, id);
    const handleDeleteUser = async () => {
        setLoadingDelete(true);
        try {
            await deleteUserById();
            toast({
                title: "Success",
                description: "User deleted successfully",
            });
        } catch (error) {
            toast({
                title: "Error",
                description: "Failed to delete user",
            });
            console.error(error);
        } finally {
            setLoadingDelete(false);
            setIsDialogOpen(false);
        }
        setLoadingDelete(false);
    };

    return (
        <>
            {isDialogOpen && (
                <Dialog open={isDialogOpen} onOpenChange={setIsDialogOpen}>
                    <DialogContent>
                        <DialogHeader>
                            <DialogTitle>
                                Are you sure to delete this user?
                            </DialogTitle>
                        </DialogHeader>

                        <DialogDescription>
                            This action cannot be undone.
                        </DialogDescription>

                        <DialogFooter>
                            <Button
                                variant="outline"
                                onClick={() => setIsDialogOpen(false)}
                            >
                                Cancel
                            </Button>
                            <Button
                                variant="destructive"
                                disabled={loadingDelete}
                                onClick={handleDeleteUser}
                            >
                                {loadingDelete && (
                                    <Loader2 className="animate-spin" />
                                )}
                                Delete
                            </Button>
                        </DialogFooter>
                    </DialogContent>
                </Dialog>
            )}

            <Button variant="destructive" onClick={() => setIsDialogOpen(true)}>
                Delete
            </Button>
        </>
    );
};

export const DeleteItem = ({
    id,
    type,
}: {
    id: string;
    type: "user" | "product";
}) => {
    const { toast } = useToast();
    const [isDialogOpen, setIsDialogOpen] = useState<boolean>(false);
    const [loadingDelete, setLoadingDelete] = useState<boolean>(false);
    let deleteFunction;
    switch (type) {
        case "user":
            deleteFunction = deleteUser;
            break;
        case "product":
            deleteFunction = deleteProduct;
            break;
        default:
            throw new Error("Invalid type");
    }
    const deleteItemById = deleteFunction.bind(null, id);

    const handleDeleteItem = async () => {
        setLoadingDelete(true);
        try {
            await deleteItemById();
            toast({
                title: "Success",
                description: `${type.charAt(0).toUpperCase() + type.slice(1)} deleted successfully`,
            });
        } catch (error) {
            toast({
                variant: "destructive",
                title: "Error",
                description: `Failed to delete ${type}`,
            });
            console.error(error);
        } finally {
            setLoadingDelete(false);
            setIsDialogOpen(false);
        }
        setLoadingDelete(false);
    };

    return (
        <>
            {isDialogOpen && (
                <Dialog open={isDialogOpen} onOpenChange={setIsDialogOpen}>
                    <DialogContent>
                        <DialogHeader>
                            <DialogTitle>
                                Are you sure to delete this {type}?
                            </DialogTitle>
                        </DialogHeader>

                        <DialogDescription>
                            This action cannot be undone.
                        </DialogDescription>

                        <DialogFooter>
                            <Button
                                variant="outline"
                                onClick={() => setIsDialogOpen(false)}
                            >
                                Cancel
                            </Button>
                            <Button
                                variant="destructive"
                                disabled={loadingDelete}
                                onClick={handleDeleteItem}
                            >
                                {loadingDelete && (
                                    <Loader2 className="animate-spin" />
                                )}
                                Delete
                            </Button>
                        </DialogFooter>
                    </DialogContent>
                </Dialog>
            )}

            <Button variant="destructive" onClick={() => setIsDialogOpen(true)}>
                Delete
            </Button>
        </>
    );
};
