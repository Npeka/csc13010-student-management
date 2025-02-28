"use client";
import { TableData } from "@/components/admin/table/table.data";
import {
  getCoreRowModel,
  useReactTable,
  ColumnDef,
} from "@tanstack/react-table";
import {
  useGetFacultiesQuery,
  useDeleteFacultyMutation,
} from "@/services/faculty-service";
import { Faculty } from "@/types/student";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Button } from "@/components/ui/button";
import { MoreHorizontal } from "lucide-react";

export const TableFaculty = () => {
  const { data, isLoading } = useGetFacultiesQuery();

  if (isLoading || !data) return <div>Loading...</div>;

  return <TableDataFaculty data={data} />;
};

const TableDataFaculty = ({ data }: { data: Faculty[] }) => {
  const columns: ColumnDef<Faculty>[] = [
    { header: "Name", accessorKey: "name" },
    {
      header: () => <div className="text-right mr-2">Action</div>,
      accessorKey: "action",
      cell: ({ row }) => (
        <div className="float-right mr-2">
          <ActionCell faculty={row.original} />,
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

const ActionCell = ({ faculty }: { faculty: Faculty }) => (
  <DropdownMenu>
    <DropdownMenuTrigger asChild>
      <Button variant="ghost" className="h-8 w-8 p-0">
        <span className="sr-only">Open menu</span>
        <MoreHorizontal className="h-4 w-4" />
      </Button>
    </DropdownMenuTrigger>
    <DropdownMenuContent align="end">
      <DropdownMenuLabel>Actions</DropdownMenuLabel>
      <ActionDeleteDropdown facultyId={String(faculty.id)} />
    </DropdownMenuContent>
  </DropdownMenu>
);

const ActionDeleteDropdown = ({ facultyId }: { facultyId: string }) => {
  const [deleteFaculty] = useDeleteFacultyMutation();

  return (
    <DropdownMenuItem
      onClick={() => {
        deleteFaculty(facultyId);
      }}
    >
      Delete
    </DropdownMenuItem>
  );
};
