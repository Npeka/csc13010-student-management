"use client";

import { useEffect } from "react";
import { type ColumnDef } from "@tanstack/react-table";
import {
  Select,
  SelectTrigger,
  SelectContent,
  SelectItem,
  SelectValue,
} from "@/components/ui/select";
import type { Column } from "@tanstack/react-table";
import { Option, OptionDTO } from "@/types/student";
import { StudentResponseDTO } from "@/types/student";
import { MoreHorizontal } from "lucide-react";

import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { useDeleteStudentMutation } from "@/services/student-service";

export const StudentColumns = ({ options }: { options?: OptionDTO }) => {
  const columns: ColumnDef<StudentResponseDTO>[] = [
    { header: "ID", accessorKey: "id" },
    { header: "Full Name", accessorKey: "full_name" },
    {
      header: "Birth Date",
      accessorKey: "birth_date",
      cell: ({ row }) => {
        return new Date(row.original.birth_date).toLocaleDateString();
      },
    },
    {
      header: ({ column }) => {
        return ColumnFilter({
          header: "Gender",
          column,
          options: options?.genders.slice() ?? [],
        });
      },
      accessorKey: "gender",
      filterFn: (row, columnId, filterValue) => {
        if (!filterValue) return true;
        return row.getValue(columnId) === filterValue;
      },
    },
    {
      header: ({ column }) => {
        return ColumnFilter({
          header: "Faculty",
          column,
          options: options?.faculties.slice() ?? [],
        });
      },
      accessorKey: "faculty",
      filterFn: (row, columnId, filterValue) => {
        if (!filterValue) return true;
        return row.getValue(columnId) === filterValue;
      },
    },
    {
      header: ({ column }) => {
        return ColumnFilter({
          header: "Course",
          column,
          options: options?.courses.slice() ?? [],
        });
      },
      accessorKey: "course",
      filterFn: (row, columnId, filterValue) => {
        if (!filterValue) return true;
        return row.getValue(columnId) === filterValue;
      },
    },
    {
      header: ({ column }) => {
        return ColumnFilter({
          header: "Program",
          column,
          options: options?.programs.slice() ?? [],
        });
      },
      accessorKey: "program",
      filterFn: (row, columnId, filterValue) => {
        if (!filterValue) return true;
        return row.getValue(columnId) === filterValue;
      },
    },
    { header: "Address", accessorKey: "address" },
    {
      header: "Email",
      accessorKey: "email",
      meta: { isHidden: true },
    },
    { header: "Phone", accessorKey: "phone" },
    {
      header: ({ column }) => {
        return ColumnFilter({
          header: "Status",
          column,
          options: options?.statuses.slice() ?? [],
        });
      },
      accessorKey: "status",
      filterFn: (row, columnId, filterValue) => {
        if (!filterValue) return true;
        return row.getValue(columnId) === filterValue;
      },
    },
    {
      header: "Actions",
      id: "actions",
      cell: ({ row }) => {
        const student = row.original;
        return (
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button variant="ghost" className="h-8 w-8 p-0">
                <span className="sr-only">Open menu</span>
                <MoreHorizontal className="h-4 w-4" />
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end">
              <DropdownMenuLabel>Actions</DropdownMenuLabel>
              <ActionEditDropdown student_id={student.student_id} />
              <ActionDeleteDropdown student_id={student.student_id} />
              <ActionExportCertificateDropdown
                student_id={student.student_id}
                format="pdf"
                text="PDF Cert"
              />
              <ActionExportCertificateDropdown
                student_id={student.student_id}
                format="docx"
                text="DOCX Cert"
              />
            </DropdownMenuContent>
          </DropdownMenu>
        );
      },
    },
  ];

  return columns;
};

const ActionExportCertificateDropdown = ({
  student_id,
  format,
  text,
}: {
  student_id: string;
  format: "pdf" | "docx";
  text?: string;
}) => {
  const handleExport = async (format: "pdf" | "docx") => {
    try {
      const response = await fetch(
        `${process.env.NEXT_PUBLIC_BACKEND_BASE_URL}/api/v1/students/certificate/${student_id}?format=${format}`
      );

      if (!response.ok) throw new Error("Failed to fetch file");

      const blob = await response.blob();
      const url = URL.createObjectURL(blob);

      const a = document.createElement("a");
      a.href = url;
      a.download = `certificate_${student_id}.${format}`;
      document.body.appendChild(a);
      a.click();

      URL.revokeObjectURL(url);
      document.body.removeChild(a);
    } catch (error) {
      console.error("Download error:", error);
    }
  };

  return (
    <DropdownMenuItem onClick={() => handleExport(format)}>
      {text || `Download as ${format.toUpperCase()}`}
    </DropdownMenuItem>
  );
};

import { useRouter } from "next/navigation";
const ActionEditDropdown = ({ student_id }: { student_id: string }) => {
  const router = useRouter();
  return (
    <DropdownMenuItem
      onClick={() => {
        router.push(`/${student_id}`);
      }}
    >
      Edit
    </DropdownMenuItem>
  );
};

const ActionDeleteDropdown = ({ student_id }: { student_id: string }) => {
  const [deleteStudent, { isLoading }] = useDeleteStudentMutation();
  return (
    <DropdownMenuItem
      onClick={() => {
        deleteStudent(student_id);
      }}
      disabled={isLoading}
    >
      Delete
    </DropdownMenuItem>
  );
};

const ColumnFilter = ({
  header,
  column,
  options,
}: {
  header: string;
  column: Column<any, unknown>;
  options: Option[];
}) => {
  const columnFilterValue = column.getFilterValue() as string | undefined;
  return (
    <Select
      value={columnFilterValue ?? "All"}
      onValueChange={(value) => {
        column.setFilterValue(value === "All" ? undefined : value);
      }}
    >
      <SelectTrigger>
        <SelectValue placeholder={`All ${header}`}>
          {columnFilterValue ? columnFilterValue : `All ${header}`}
        </SelectValue>
      </SelectTrigger>
      <SelectContent>
        <SelectItem value="All">All</SelectItem>
        {options.map((option) => (
          <SelectItem key={option.name} value={option.name}>
            {option.name}
          </SelectItem>
        ))}
      </SelectContent>
    </Select>
  );
};
