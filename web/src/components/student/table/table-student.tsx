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
} from "@tanstack/react-table";
import {
  type ColumnDef,
  type ColumnFiltersState,
  type SortingState,
  getPaginationRowModel,
} from "@tanstack/react-table";

import { OptionDTO, StudentResponseDTO } from "@/types/student";
import { StudentColumns } from "./table-student-columns";
import {
  DropdownMenu,
  DropdownMenuCheckboxItem,
  DropdownMenuContent,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Button } from "@/components/ui/button";
import { VisibilityState } from "@tanstack/react-table";
import type { Table } from "@tanstack/react-table";

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
    data: data,
    columns: StudentColumns({ options }),
    getCoreRowModel: getCoreRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
    onSortingChange: setSorting,
    getSortedRowModel: getSortedRowModel(),
    onColumnFiltersChange: setColumnFilters,
    getFilteredRowModel: getFilteredRowModel(),
    onColumnVisibilityChange: setColumnVisibility,
    state: {
      sorting,
      columnFilters,
      columnVisibility,
    },
  });

  useEffect(() => {
    const hideColunms = (keys: string[]): { [key: string]: boolean } => {
      return keys.reduce((acc, key) => {
        acc[key] = false;
        return acc;
      }, {} as { [key: string]: boolean });
    };

    table.setColumnVisibility(hideColunms(["email", "birth_date", "address"]));
  }, []);

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
import { DropdownMenuItem } from "@/components/ui/dropdown-menu";
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

import {
  Dialog,
  DialogTrigger,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogDescription,
} from "@/components/ui/dialog";
import { useDropzone } from "react-dropzone";
import { Upload, Download, File } from "lucide-react";
import { Loader2 } from "lucide-react";
import { useToast } from "@/hooks/use-toast";

const ImportFileButton = () => {
  const { toast } = useToast();
  const [files, setFiles] = useState<File[]>([]);
  const [open, setOpen] = useState(false);
  const [loading, setLoading] = useState(false);

  const onDrop = (acceptedFiles: File[]) => {
    setFiles(acceptedFiles);
  };

  const { getRootProps, getInputProps, isDragActive } = useDropzone({
    accept: {
      "text/csv": [".csv"],
      "application/json": [".json"],
    },
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

    setLoading(true);
    const formData = new FormData();
    formData.append("students", files[0]);

    try {
      const response = await fetch(
        `${process.env.NEXT_PUBLIC_BACKEND_BASE_URL}/api/v1/students/import`,
        {
          method: "POST",
          body: formData,
        }
      );

      if (!response.ok) throw new Error("Error uploading file!");

      toast({
        title: "Success",
        description: "Import success!",
      });
      setFiles([]);
      setOpen(false);
    } catch (error) {
      toast({
        title: "Error",
        description: "Error uploading file!",
        variant: "destructive",
      });
    } finally {
      setLoading(false);
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

        {/* Drag and drop file */}
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

        {/* Selected file list */}
        {files.length > 0 && (
          <div className="mt-4 p-3 border rounded-lg flex items-center justify-between bg-gray-100">
            <File className="w-5 h-5 text-gray-600" />
            <span className="text-sm font-medium">{files[0].name}</span>
            <Button variant="ghost" size="sm" onClick={() => setFiles([])}>
              Remove
            </Button>
          </div>
        )}

        {/* Upload button */}
        <Button
          className="w-full mt-4"
          onClick={handleUpload}
          disabled={files.length === 0 || loading}
        >
          {loading ? (
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
  const [downloading, setDownloading] = useState(false);

  const handleDownload = async () => {
    setDownloading(true);
    try {
      const response = await fetch(
        `${process.env.NEXT_PUBLIC_BACKEND_BASE_URL}/api/v1/students/export?format=${format}`,
        {
          method: "GET",
        }
      );

      if (!response.ok) throw new Error("Failed to fetch file");

      // Lấy Blob từ response
      const blob = await response.blob();
      const url = URL.createObjectURL(blob);

      // Tạo thẻ <a> để tải file
      const a = document.createElement("a");
      a.href = url;
      a.download = `students.${format}`;
      document.body.appendChild(a);
      a.click();

      // Cleanup
      URL.revokeObjectURL(url);
      document.body.removeChild(a);
    } catch (error) {
      console.error("Download error:", error);
    } finally {
      setDownloading(false);
    }
  };

  return (
    <DropdownMenuItem onClick={handleDownload} disabled={downloading}>
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
          .map((column) => {
            return (
              <DropdownMenuCheckboxItem
                key={column.id}
                className="capitalize"
                checked={column.getIsVisible()}
                onCheckedChange={(value) => column.toggleVisibility(!!value)}
              >
                {column.id}
              </DropdownMenuCheckboxItem>
            );
          })}
      </DropdownMenuContent>
    </DropdownMenu>
  );
};
