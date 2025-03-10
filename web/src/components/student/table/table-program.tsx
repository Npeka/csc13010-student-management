"use client";
import { useState } from "react";
import { TableData } from "@/components/admin/table/table.data";
import {
  getCoreRowModel,
  useReactTable,
  ColumnDef,
} from "@tanstack/react-table";
import {
  useGetProgramsQuery,
  useUpdateProgramMutation,
  useDeleteProgramMutation,
} from "@/services/program-service";
import { Program } from "@/types/student";
import { ActionCell, EditModal } from "./table-ui";

export const TableProgram = () => {
  const { data, isLoading } = useGetProgramsQuery();
  if (isLoading || !data) return <div>Loading...</div>;
  return <TableDataProgram data={data.data} />;
};

const TableDataProgram = ({ data }: { data: Program[] }) => {
  const [editProgram, setEditProgram] = useState<Program | null>(null);
  const [updateProgram] = useUpdateProgramMutation();
  const [deleteProgram] = useDeleteProgramMutation();

  const columns: ColumnDef<Program>[] = [
    { header: "Name", accessorKey: "name" },
    {
      header: () => <div className="text-right mr-2">Action</div>,
      accessorKey: "action",
      cell: ({ row }) => (
        <div className="float-right mr-2">
          <ActionCell
            item={row.original}
            onEdit={setEditProgram}
            onDelete={(id) => deleteProgram(id)}
          />
        </div>
      ),
    },
  ];

  const table = useReactTable({
    data,
    columns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <>
      <TableData table={table} />
      <EditModal
        item={editProgram}
        setItem={setEditProgram}
        fields={[{ key: "name", label: "Name of Program" }]}
        onSave={(updatedProgram) => updateProgram(updatedProgram)}
      />
    </>
  );
};
