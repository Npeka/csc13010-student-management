"use client";
import { TableData } from "@/components/admin/table/table.data";
import {
  getCoreRowModel,
  useReactTable,
  ColumnDef,
} from "@tanstack/react-table";
import {
  useGetProgramsQuery,
  useDeleteProgramMutation,
} from "@/services/program-service";
import { Program } from "@/types/student";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Button } from "@/components/ui/button";
import { MoreHorizontal } from "lucide-react";

export const TableProgram = () => {
  const { data, isLoading } = useGetProgramsQuery();

  if (isLoading || !data) return <div>Loading...</div>;

  return <TableDataProgram data={data} />;
};

const TableDataProgram = ({ data }: { data: Program[] }) => {
  const columns: ColumnDef<Program>[] = [
    // { header: "ID", accessorKey: "id" },
    { header: "Name", accessorKey: "name" },
    {
      header: () => <div className="text-right mr-2">Action</div>,
      accessorKey: "action",
      cell: ({ row }) => (
        <div className="float-right mr-2">
          <ActionCell program={row.original} />,
        </div>
      ),
    },
  ];

  const table = useReactTable({
    data,
    columns,
    getCoreRowModel: getCoreRowModel(),
  });

  return <TableData table={table} />;
};

const ActionCell = ({ program }: { program: Program }) => (
  <DropdownMenu>
    <DropdownMenuTrigger asChild>
      <Button variant="ghost" className="h-8 w-8 p-0">
        <span className="sr-only">Open menu</span>
        <MoreHorizontal className="h-4 w-4" />
      </Button>
    </DropdownMenuTrigger>
    <DropdownMenuContent align="end">
      <DropdownMenuLabel>Actions</DropdownMenuLabel>
      <ActionDeleteDropdown programId={String(program.id)} />
    </DropdownMenuContent>
  </DropdownMenu>
);

const ActionDeleteDropdown = ({ programId }: { programId: string }) => {
  const [deleteProgram] = useDeleteProgramMutation();

  return (
    <DropdownMenuItem
      onClick={() => {
        deleteProgram(programId);
      }}
    >
      Delete
    </DropdownMenuItem>
  );
};
