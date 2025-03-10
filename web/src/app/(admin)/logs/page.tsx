"use client";
import { PageTitle } from "../page-title";
import { useGetAuditLogsQuery } from "@/services/auditlogs-service";
import { TableData } from "@/components/admin/table/table.data";
import {
  getCoreRowModel,
  useReactTable,
  ColumnDef,
} from "@tanstack/react-table";
import { AuditLog } from "@/types/auditlog";
import { Badge } from "@/components/ui/badge";
import { format } from "date-fns";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Skeleton } from "@/components/ui/skeleton";
import { Button } from "@/components/ui/button";
import { Dialog, DialogContent, DialogTrigger } from "@/components/ui/dialog";
import { ScrollArea } from "@/components/ui/scroll-area";
import { memo } from "react";

export default function LogsPage() {
  return (
    <>
      <PageTitle title="Logs" />
      <TableAuditlog />
    </>
  );
}

const TableAuditlog = memo(() => {
  const { data, isLoading } = useGetAuditLogsQuery();
  if (isLoading) return <Skeleton className="h-60 w-full" />;
  return <TableDataAuditlog data={data?.data || []} />;
});

const TableDataAuditlog = memo(({ data }: { data: AuditLog[] }) => {
  const columns: ColumnDef<AuditLog>[] = [
    { header: "ID", accessorKey: "id" },
    { header: "Table Name", accessorKey: "table_name" },
    { header: "Record ID", accessorKey: "record_id" },
    {
      header: "Old Record",
      accessorKey: "old_record",
      cell: ({ row }) => <JsonDialog jsonData={row.original.old_record} />,
    },
    {
      header: "New Record",
      accessorKey: "new_record",
      cell: ({ row }) => <JsonDialog jsonData={row.original.new_record} />,
    },
    {
      header: "Changed Fields",
      accessorKey: "changed_fields",
      cell: ({ row }) => <JsonDialog jsonData={row.original.field_changes} />,
    },
    {
      header: "Action",
      accessorKey: "action",
      cell: ({ row }) => (
        <Badge
          variant={
            row.original.action === "CREATE"
              ? "default"
              : row.original.action === "UPDATE"
              ? "secondary"
              : "destructive"
          }
        >
          {row.original.action}
        </Badge>
      ),
    },
    {
      header: "Created At",
      accessorKey: "created_at",
      cell: ({ row }) =>
        format(new Date(row.original.created_at), "yyyy-MM-dd HH:mm:ss"),
    },
  ];

  const table = useReactTable({
    data,
    columns,
    getCoreRowModel: getCoreRowModel(),
  });

  return <TableData table={table} />;
});

const JsonDialog = memo(({ jsonData }: { jsonData: string }) => {
  let parsedData;
  try {
    parsedData = JSON.parse(jsonData);
  } catch (error) {
    parsedData = { error: "Invalid JSON" };
  }

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button variant="outline" size="sm">
          View
        </Button>
      </DialogTrigger>
      <DialogContent className="max-w-lg">
        <h3 className="text-lg font-semibold">Changed Fields</h3>
        <ScrollArea className="h-60 border rounded-lg p-2 bg-gray-50">
          <pre className="text-xs bg-gray-100 p-2 rounded-md">
            {JSON.stringify(parsedData, null, 2)}
          </pre>
        </ScrollArea>
      </DialogContent>
    </Dialog>
  );
});
