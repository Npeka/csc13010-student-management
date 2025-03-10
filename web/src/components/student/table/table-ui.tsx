"use client";
import { useEffect, useState } from "react";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogFooter,
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { MoreHorizontal } from "lucide-react";

type EditModalProps<T> = {
  item: T | null;
  setItem: (item: T | null) => void;
  fields: { key: keyof T; label: string }[];
  onSave: (updatedItem: T) => void;
};

export const EditModal = <T,>({
  item,
  setItem,
  fields,
  onSave,
}: EditModalProps<T>) => {
  const [values, setValues] = useState<Partial<T>>({});

  useEffect(() => {
    if (item) setValues(item);
  }, [item]);

  if (!item) return null;

  return (
    <Dialog open={!!item} onOpenChange={() => setItem(null)}>
      <DialogContent>
        <DialogHeader>Edit</DialogHeader>
        {fields.map((field) => (
          <Input
            key={String(field.key)}
            value={(values[field.key] as string) || ""}
            onChange={(e) =>
              setValues((prev) => ({ ...prev, [field.key]: e.target.value }))
            }
            placeholder={field.label}
          />
        ))}
        <DialogFooter>
          <Button variant="outline" onClick={() => setItem(null)}>
            Cancel
          </Button>
          <Button
            onClick={() => {
              onSave({ ...item, ...values } as T);
              setItem(null);
            }}
          >
            Save
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
};

type ActionCellProps<T> = {
  item: T;
  onEdit: (item: T) => void;
  onDelete: (id: string) => void;
  idKey?: keyof T; // Key để lấy ID (mặc định "id")
};

export const ActionCell = <T,>({
  item,
  onEdit,
  onDelete,
  idKey = "id" as keyof T,
}: ActionCellProps<T>) => {
  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button variant="ghost" className="h-8 w-8 p-0">
          <span className="sr-only">Open menu</span>
          <MoreHorizontal className="h-4 w-4" />
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end">
        <DropdownMenuLabel>Actions</DropdownMenuLabel>
        <DropdownMenuItem onClick={() => onEdit(item)}>Edit</DropdownMenuItem>
        <DropdownMenuItem onClick={() => onDelete(String(item[idKey]))}>
          Delete
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
};

type ActionDeleteDropdownProps<T> = {
  idKey?: keyof T;
  onDelete: (id: string) => void;
};

export const ActionDeleteDropdown = ({
  idKey = "id",
  onDelete,
}: ActionDeleteDropdownProps<any>) => {
  return (
    <DropdownMenuItem
      onClick={() => {
        onDelete(String(idKey));
      }}
    >
      Delete
    </DropdownMenuItem>
  );
};
