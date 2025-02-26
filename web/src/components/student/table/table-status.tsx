"use client";
import { TableData } from "@/components/admin/table/table.data";
import {
  getCoreRowModel,
  useReactTable,
  ColumnDef,
} from "@tanstack/react-table";
import {
  useGetStatusesQuery,
  useDeleteStatusMutation,
} from "@/services/student-service";
import { Status } from "@/types/student";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Button } from "@/components/ui/button";
import { MoreHorizontal } from "lucide-react";

export const TableStatus = () => {
  const { data, isLoading } = useGetStatusesQuery();

  if (isLoading || !data) return <div>Loading...</div>;

  return <TableDataStatus data={data} />;
};

const TableDataStatus = ({ data }: { data: Status[] }) => {
  const columns: ColumnDef<Status>[] = [
    { header: "ID", accessorKey: "id" },
    { header: "Name", accessorKey: "name" },
    {
      header: "Action",
      accessorKey: "action",
      cell: ({ row }) => <ActionCell status={row.original} />,
    },
  ];

  const table = useReactTable({
    data,
    columns,
    getCoreRowModel: getCoreRowModel(),
  });

  return <TableData table={table} />;
};

const ActionCell = ({ status }: { status: Status }) => (
  <DropdownMenu>
    <DropdownMenuTrigger asChild>
      <Button variant="ghost" className="h-8 w-8 p-0">
        <span className="sr-only">Open menu</span>
        <MoreHorizontal className="h-4 w-4" />
      </Button>
    </DropdownMenuTrigger>
    <DropdownMenuContent align="end">
      <DropdownMenuLabel>Actions</DropdownMenuLabel>
      <ActionDeleteDropdown statusId={String(status.id)} />
    </DropdownMenuContent>
  </DropdownMenu>
);

const ActionDeleteDropdown = ({ statusId }: { statusId: string }) => {
  const [deleteStatus] = useDeleteStatusMutation();

  return (
    <DropdownMenuItem
      onClick={() => {
        deleteStatus(statusId);
      }}
    >
      Delete
    </DropdownMenuItem>
  );
};
