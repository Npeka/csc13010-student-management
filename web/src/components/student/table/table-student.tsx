"use client";
import { useState, useEffect } from "react";
import { TableData } from "@/components/admin/table/table.data";
import { TablePagination } from "@/components/admin/table/table.pagination";
import { TableFilterRow } from "@/components/admin/table/table.action";
import {
  getCoreRowModel,
  getFilteredRowModel,
  getSortedRowModel,
  useReactTable,
  type ColumnDef,
  type ColumnFiltersState,
  type SortingState,
  getPaginationRowModel,
  VisibilityState,
  type Table,
} from "@tanstack/react-table";
import { OptionDTO, StudentResponseDTO } from "@/types/student";
import { StudentColumns } from "./table-student-columns";
import {
  DropdownMenu,
  DropdownMenuCheckboxItem,
  DropdownMenuContent,
  DropdownMenuTrigger,
  DropdownMenuItem,
} from "@/components/ui/dropdown-menu";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogTrigger,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogDescription,
} from "@/components/ui/dialog";
import { useDropzone } from "react-dropzone";
import { Upload, Download, File, Loader2 } from "lucide-react";
import { useToast } from "@/hooks/use-toast";
import {
  useImportFileMutation,
  useLazyExportFileQuery,
} from "@/services/fileprocessor-service";

export const StudentTable = ({
  data,
  options,
}: {
  data: StudentResponseDTO[];
  options: OptionDTO;
}) => {
  const [sorting, setSorting] = useState<SortingState>([]);
  const [columnFilters, setColumnFilters] = useState<ColumnFiltersState>([]);
  const [columnVisibility, setColumnVisibility] = useState<VisibilityState>({});

  const table = useReactTable({
    data,
    columns: StudentColumns({ options }),
    getCoreRowModel: getCoreRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
    onSortingChange: setSorting,
    getSortedRowModel: getSortedRowModel(),
    onColumnFiltersChange: setColumnFilters,
    getFilteredRowModel: getFilteredRowModel(),
    onColumnVisibilityChange: setColumnVisibility,
    state: { sorting, columnFilters, columnVisibility },
  });

  useEffect(() => {
    const hideColumns = (keys: string[]): { [key: string]: boolean } => {
      return keys.reduce((acc, key) => {
        acc[key] = false;
        return acc;
      }, {} as { [key: string]: boolean });
    };

    table.setColumnVisibility(hideColumns(["email", "birth_date", "address"]));
  }, [table]);

  return (
    <>
      <div className="flex justify-between items-center gap-2">
        <TableFilterRow table={table} by="full_name" />
        <div className="flex gap-2">
          <ImportFileButton />
          <TableExportData />
          <TableVisibleColumns table={table} />
        </div>
      </div>
      <TableData table={table} />
      <TablePagination table={table} />
    </>
  );
};

const TableExportData = () => {
  return (
    <DropdownMenu>
      <DropdownMenuTrigger>
        <Button variant="outline" className="ml-auto">
          <Download />
          Export
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent>
        <ExportFileButton text="CSV" format="csv" />
        <ExportFileButton text="JSON" format="json" />
      </DropdownMenuContent>
    </DropdownMenu>
  );
};

const ImportFileButton = () => {
  const { toast } = useToast();
  const [files, setFiles] = useState<File[]>([]);
  const [open, setOpen] = useState(false);
  const [triggerImportFile, { isLoading }] = useImportFileMutation();

  const onDrop = (acceptedFiles: File[]) => {
    setFiles(acceptedFiles);
  };

  const { getRootProps, getInputProps, isDragActive } = useDropzone({
    accept: { "text/csv": [".csv"], "application/json": [".json"] },
    multiple: false,
    onDrop,
  });

  const handleUpload = async () => {
    if (files.length === 0) {
      toast({
        title: "Error",
        description: "Please select a CSV or JSON file!",
        variant: "destructive",
      });
      return;
    }

    try {
      await triggerImportFile({
        file: files[0],
        format: files[0].type === "application/json" ? "json" : "csv",
        module: "students",
      }).unwrap();
      toast({ title: "Success", description: "Import success!" });
      setFiles([]);
      setOpen(false);
    } catch (error) {
      toast({
        title: "Error",
        description: "Error uploading file!",
        variant: "destructive",
      });
    }
  };

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <Button variant="outline">
          <Upload />
          Import
        </Button>
      </DialogTrigger>
      <DialogContent className="max-w-lg">
        <DialogHeader>
          <DialogTitle>Import Students</DialogTitle>
          <DialogDescription>
            Select or drag and drop a CSV / JSON file to import.
          </DialogDescription>
        </DialogHeader>
        <div
          {...getRootProps()}
          className="border-2 border-dashed rounded-lg p-6 text-center cursor-pointer hover:bg-gray-100 transition"
        >
          <input {...getInputProps()} />
          {isDragActive ? (
            <p className="text-lg font-semibold">Drop the file here...</p>
          ) : (
            <p className="text-lg font-semibold text-gray-600">
              Drag and drop a file or click to select
            </p>
          )}
        </div>
        {files.length > 0 && (
          <div className="mt-4 p-3 border rounded-lg flex items-center justify-between bg-gray-100">
            <File className="w-5 h-5 text-gray-600" />
            <span className="text-sm font-medium">{files[0].name}</span>
            <Button variant="ghost" size="sm" onClick={() => setFiles([])}>
              Remove
            </Button>
          </div>
        )}
        <Button
          className="w-full mt-4"
          onClick={handleUpload}
          disabled={files.length === 0 || isLoading}
        >
          {isLoading ? (
            <Loader2 className="w-4 h-4 animate-spin mr-2" />
          ) : (
            "Upload File"
          )}
        </Button>
      </DialogContent>
    </Dialog>
  );
};

const ExportFileButton = ({
  text,
  format,
}: {
  text: string;
  format: "csv" | "json";
}) => {
  const [triggerExportFile, { isLoading }] = useLazyExportFileQuery();

  const handleClick = async () => {
    try {
      const blob = await triggerExportFile({
        format,
        module: "students",
      }).unwrap();
      const url = window.URL.createObjectURL(new Blob([blob]));
      const a = document.createElement("a");
      a.href = url;
      a.download = `export.${format}`;
      document.body.appendChild(a);
      a.click();
      a.remove();
      window.URL.revokeObjectURL(url);
    } catch (error) {
      console.error("Export file failed:", error);
    }
  };

  return (
    <DropdownMenuItem onClick={handleClick} disabled={isLoading}>
      {text}
    </DropdownMenuItem>
  );
};

const TableVisibleColumns = ({ table }: { table: Table<any> }) => {
  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button variant="outline" className="ml-auto">
          Columns
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end">
        {table
          .getAllColumns()
          .filter((column) => column.getCanHide())
          .map((column) => (
            <DropdownMenuCheckboxItem
              key={column.id}
              className="capitalize"
              checked={column.getIsVisible()}
              onCheckedChange={(value) => column.toggleVisibility(!!value)}
            >
              {column.id}
            </DropdownMenuCheckboxItem>
          ))}
      </DropdownMenuContent>
    </DropdownMenu>
  );
};
