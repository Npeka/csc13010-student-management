"use client";
import { useState } from "react";
import { TableData } from "@/components/admin/table/table.data";
import {
  getCoreRowModel,
  useReactTable,
  ColumnDef,
} from "@tanstack/react-table";
import {
  useGetFacultiesQuery,
  useUpdateFacultyMutation,
  useDeleteFacultyMutation,
} from "@/services/faculty-service";
import { Faculty } from "@/types/student";
import { ActionCell, EditModal } from "./table-ui";

export const TableFaculty = () => {
  const { data, isLoading } = useGetFacultiesQuery();
  if (isLoading || !data) return <div>Loading...</div>;
  return <TableDataFaculty data={data.data} />;
};

const TableDataFaculty = ({ data }: { data: Faculty[] }) => {
  const [editFaculty, setEditFaculty] = useState<Faculty | null>(null);
  const [updateFaculty] = useUpdateFacultyMutation();
  const [deleteFaculty] = useDeleteFacultyMutation();

  const columns: ColumnDef<Faculty>[] = [
    { header: "Name", accessorKey: "name" },
    {
      header: () => <div className="text-right mr-2">Action</div>,
      accessorKey: "action",
      cell: ({ row }) => (
        <div className="float-right mr-2">
          <ActionCell
            item={row.original}
            onEdit={setEditFaculty}
            onDelete={(id) => deleteFaculty(id)}
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
        item={editFaculty}
        setItem={setEditFaculty}
        fields={[{ key: "name", label: "Name of Faculty" }]}
        onSave={(updatedFaculty) => updateFaculty(updatedFaculty)}
      />
    </>
  );
};
