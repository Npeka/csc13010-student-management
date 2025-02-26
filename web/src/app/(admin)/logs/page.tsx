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

export const TableAuditlog = () => {
  const { data, isLoading, error } = useGetAuditLogsQuery();

  if (isLoading) return <Skeleton className="h-60 w-full" />;
  if (error)
    return <p className="text-center text-red-500">Failed to load data</p>;

  return <TableDataAuditlog data={data || []} />;
};

const TableDataAuditlog = ({ data }: { data: AuditLog[] }) => {
  const columns: ColumnDef<AuditLog>[] = [
    { header: "ID", accessorKey: "id" },
    { header: "Table Name", accessorKey: "table_name" },
    { header: "Record ID", accessorKey: "record_id" },
    {
      header: "Changed Fields",
      accessorKey: "changed_fields",
      cell: ({ row }) => <JsonDialog jsonData={row.original.changed_fields} />,
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
      header: "Changed By",
      accessorKey: "changed_by",
      cell: ({ row }) => (
        <Badge
          variant={
            row.original.changed_by === "ADMIN" ? "secondary" : "default"
          }
        >
          {row.original.changed_by}
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

  return (
    <Card>
      <CardHeader>
        <CardTitle>Audit Logs</CardTitle>
      </CardHeader>
      <CardContent>
        <TableData table={table} />
      </CardContent>
    </Card>
  );
};

// Component hiển thị JSON trong popup
const JsonDialog = ({ jsonData }: { jsonData: string }) => {
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
};

export default function LogsPage() {
  return (
    <>
      <PageTitle title="Logs" />
      <TableAuditlog />
    </>
  );
}
