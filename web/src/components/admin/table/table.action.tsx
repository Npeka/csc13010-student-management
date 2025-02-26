import { ArrowUpDown } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { type Table as TableType, type Column } from "@tanstack/react-table";

export const TableSortingCol = ({
  column,
  by,
}: {
  column: Column<any, any>;
  by: string;
}) => {
  return (
    <Button
      className="px-0"
      variant="ghost"
      onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
    >
      {by}
      <ArrowUpDown className="ml-2 h-4 w-4" />
    </Button>
  );
};

export const TableFilterRow = ({
  table,
  by,
}: {
  table: TableType<any>;
  by: string;
}) => {
  return (
    <Input
      className="max-w-sm"
      placeholder={`Filter by ${by}`}
      value={(table.getColumn(by)?.getFilterValue() as string) ?? ""}
      onChange={(event) =>
        table.getColumn(by)?.setFilterValue(event.target.value)
      }
    />
  );
};
