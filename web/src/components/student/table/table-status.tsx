"use client";
import { useState } from "react";
import { TableData } from "@/components/admin/table/table.data";
import {
  getCoreRowModel,
  useReactTable,
  ColumnDef,
} from "@tanstack/react-table";
import {
  useGetStatusesQuery,
  useUpdateStatusMutation,
  useDeleteStatusMutation,
} from "@/services/status-service";
import { Status } from "@/types/student";
import { ActionCell, EditModal } from "./table-ui";

export const TableStatus = () => {
  const { data, isLoading } = useGetStatusesQuery();
  if (isLoading || !data) return <div>Loading...</div>;
  return <TableDataStatus data={data.data} />;
};

const TableDataStatus = ({ data }: { data: Status[] }) => {
  const [editStatus, setEditStatus] = useState<Status | null>(null);
  const [updateStatus] = useUpdateStatusMutation();
  const [deleteStatus] = useDeleteStatusMutation();

  const columns: ColumnDef<Status>[] = [
    { header: "Name", accessorKey: "name" },
    {
      header: () => <div className="text-right mr-2">Action</div>,
      accessorKey: "action",
      cell: ({ row }) => (
        <div className="float-right mr-2">
          <ActionCell
            item={row.original}
            onEdit={setEditStatus}
            onDelete={(id) => deleteStatus(id)}
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
        item={editStatus}
        setItem={setEditStatus}
        fields={[{ key: "name", label: "Name of Status" }]}
        onSave={(updatedStatus) => updateStatus(updatedStatus)}
      />
    </>
  );
};
