"use client";
import { useState, useCallback } from "react";
import { TableData } from "@/components/admin/table/table.data";
import { TablePagination } from "@/components/admin/table/table.pagination";
import { TableSortingCol } from "@/components/admin/table/table.action";
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
import {
  Select,
  SelectTrigger,
  SelectContent,
  SelectItem,
  SelectValue,
} from "@/components/ui/select";
import type { Column } from "@tanstack/react-table";

interface Student {
  id: string;
  fullName: string;
  birthDate: string;
  gender: Gender;
  faculty: Faculty;
  course: Course;
  program: Program;
  address: string;
  email: string;
  phone: string;
  status: Status;
}

const GENDERS = ["Male", "Female", "Non-binary"] as const;
const COURSES = ["Computer Science", "Biology", "History"] as const;
const FACULTIES = ["Engineering", "Science", "Arts"] as const;
const PROGRAMS = ["Undergraduate", "Graduate"] as const;
const STATUSES = ["Actived", "Inactive"] as const;

type Gender = (typeof GENDERS)[number];
type Course = (typeof COURSES)[number];
type Faculty = (typeof FACULTIES)[number];
type Program = (typeof PROGRAMS)[number];
type Status = (typeof STATUSES)[number];
// fake data
const students: Student[] = [
  {
    id: "1",
    fullName: "Student 1",
    birthDate: "2001-01-01",
    gender: "Male",
    faculty: "Engineering",
    course: "Computer Science",
    program: "Undergraduate",
    address: "123 Main St",
    email: "student1@example.com",
    phone: "123-456-7890",
    status: "Actived",
  },
  {
    id: "2",
    fullName: "Student 2",
    birthDate: "2002-02-02",
    gender: "Female",
    faculty: "Science",
    course: "Biology",
    program: "Undergraduate",
    address: "456 Elm St",
    email: "student2@example.com",
    phone: "234-567-8901",
    status: "Actived",
  },
  {
    id: "3",
    fullName: "Student 3",
    birthDate: "2003-03-03",
    gender: "Non-binary",
    faculty: "Arts",
    course: "History",
    program: "Undergraduate",
    address: "789 Oak St",
    email: "student3@example.com",
    phone: "345-678-9012",
    status: "Inactive",
  },
  {
    id: "3",
    fullName: "Student 3",
    birthDate: "2003-03-03",
    gender: "Non-binary",
    faculty: "Arts",
    course: "History",
    program: "Undergraduate",
    address: "789 Oak St",
    email: "student3@example.com",
    phone: "345-678-9012",
    status: "Inactive",
  },
  {
    id: "3",
    fullName: "Student 3",
    birthDate: "2003-03-03",
    gender: "Non-binary",
    faculty: "Arts",
    course: "History",
    program: "Undergraduate",
    address: "789 Oak St",
    email: "student3@example.com",
    phone: "345-678-9012",
    status: "Inactive",
  },
  {
    id: "3",
    fullName: "Student 3",
    birthDate: "2003-03-03",
    gender: "Non-binary",
    faculty: "Arts",
    course: "History",
    program: "Undergraduate",
    address: "789 Oak St",
    email: "student3@example.com",
    phone: "345-678-9012",
    status: "Inactive",
  },
  {
    id: "3",
    fullName: "Student 3",
    birthDate: "2003-03-03",
    gender: "Non-binary",
    faculty: "Arts",
    course: "History",
    program: "Undergraduate",
    address: "789 Oak St",
    email: "student3@example.com",
    phone: "345-678-9012",
    status: "Inactive",
  },
  {
    id: "3",
    fullName: "Student 3",
    birthDate: "2003-03-03",
    gender: "Non-binary",
    faculty: "Arts",
    course: "History",
    program: "Undergraduate",
    address: "789 Oak St",
    email: "student3@example.com",
    phone: "345-678-9012",
    status: "Inactive",
  },
  {
    id: "3",
    fullName: "Student 3",
    birthDate: "2003-03-03",
    gender: "Non-binary",
    faculty: "Arts",
    course: "History",
    program: "Undergraduate",
    address: "789 Oak St",
    email: "student3@example.com",
    phone: "345-678-9012",
    status: "Inactive",
  },
  {
    id: "3",
    fullName: "Student 3",
    birthDate: "2003-03-03",
    gender: "Non-binary",
    faculty: "Arts",
    course: "History",
    program: "Undergraduate",
    address: "789 Oak St",
    email: "student3@example.com",
    phone: "345-678-9012",
    status: "Inactive",
  },
  {
    id: "3",
    fullName: "Student 3",
    birthDate: "2003-03-03",
    gender: "Non-binary",
    faculty: "Arts",
    course: "History",
    program: "Undergraduate",
    address: "789 Oak St",
    email: "student3@example.com",
    phone: "345-678-9012",
    status: "Inactive",
  },
  {
    id: "3",
    fullName: "Student 3",
    birthDate: "2003-03-03",
    gender: "Non-binary",
    faculty: "Arts",
    course: "History",
    program: "Undergraduate",
    address: "789 Oak St",
    email: "student3@example.com",
    phone: "345-678-9012",
    status: "Inactive",
  },
  {
    id: "3",
    fullName: "Student 3",
    birthDate: "2003-03-03",
    gender: "Non-binary",
    faculty: "Arts",
    course: "History",
    program: "Undergraduate",
    address: "789 Oak St",
    email: "student3@example.com",
    phone: "345-678-9012",
    status: "Inactive",
  },
  {
    id: "3",
    fullName: "Student 3",
    birthDate: "2003-03-03",
    gender: "Non-binary",
    faculty: "Arts",
    course: "History",
    program: "Undergraduate",
    address: "789 Oak St",
    email: "student3@example.com",
    phone: "345-678-9012",
    status: "Inactive",
  },
  {
    id: "3",
    fullName: "Student 3",
    birthDate: "2003-03-03",
    gender: "Non-binary",
    faculty: "Arts",
    course: "History",
    program: "Undergraduate",
    address: "789 Oak St",
    email: "student3@example.com",
    phone: "345-678-9012",
    status: "Inactive",
  },
  {
    id: "3",
    fullName: "Student 3",
    birthDate: "2003-03-03",
    gender: "Non-binary",
    faculty: "Arts",
    course: "History",
    program: "Undergraduate",
    address: "789 Oak St",
    email: "student3@example.com",
    phone: "345-678-9012",
    status: "Inactive",
  },
  {
    id: "3",
    fullName: "Student 3",
    birthDate: "2003-03-03",
    gender: "Non-binary",
    faculty: "Arts",
    course: "History",
    program: "Undergraduate",
    address: "789 Oak St",
    email: "student3@example.com",
    phone: "345-678-9012",
    status: "Inactive",
  },
];

import { Button } from "@/components/ui/button";
export default function AdminStudentsPage({
  searchParams,
}: {
  searchParams: Promise<{
    page: string;
    limit: string;
    sort: string;
    search: string;
  }>;
}) {
  // const page = parseInt((await searchParams).page || "1", 10);
  // const limit = parseInt((await searchParams).limit || "5", 10);
  // const search = (await searchParams).search || "";
  // const sort = (await searchParams).sort || "desc";
  // const { data: products, total } = await getProducts({
  //   page,
  //   limit,
  //   search,
  //   sort,
  // });
  // fake data

  const columns: ColumnDef<Student>[] = [
    {
      header: "ID",
      accessorKey: "id",
    },
    { header: "Full Name", accessorKey: "fullName" },
    { header: "Birth Date", accessorKey: "birthDate" },
    {
      header: ({ column }) => {
        return ColumnFilter({
          header: "Gender",
          column,
          options: GENDERS.slice(),
        });
      },
      accessorKey: "gender",
    },
    {
      header: ({ column }) => {
        return ColumnFilter({
          header: "Faculty",
          column,
          options: FACULTIES.slice(),
        });
      },
      accessorKey: "faculty",
    },
    {
      header: ({ column }) => {
        return ColumnFilter({
          header: "Course",
          column,
          options: COURSES.slice(),
        });
      },
      accessorKey: "course",
    },
    {
      header: ({ column }) => {
        return ColumnFilter({
          header: "Program",
          column,
          options: PROGRAMS.slice(),
        });
      },
      accessorKey: "program",
    },
    { header: "Address", accessorKey: "address" },
    { header: "Email", accessorKey: "email" },
    { header: "Phone", accessorKey: "phone" },
    {
      header: ({ column }) => {
        return ColumnFilter({
          header: "Status",
          column,
          options: STATUSES.slice(),
        });
      },
      accessorKey: "status",
    },
  ];

  const [sorting, setSorting] = useState<SortingState>([]);
  const [columnFilters, setColumnFilters] = useState<ColumnFiltersState>([]);
  const [selectedProduct, setSelectedProduct] = useState<Student | null>(null);

  // const newColumns: ColumnDef<Product>[] = useMemo(() => {
  //   return columns.map((column) => {
  //     if (column.header === "Action") {
  //       return {
  //         ...column,
  //         cell: ({ row }) => {
  //           const product = row.original;
  //           return (
  //             <ProductAction
  //               onEdit={() => openAlert("edit", product)}
  //               onDelete={() => openAlert("delete", product)}
  //             />
  //           );
  //         },
  //       };
  //     }
  //     return column;
  //   });
  // }, [sorting, columnFilters]);

  const [studentss, setStudents] = useState<Student[]>(students);
  const handleImport = (newData: Student[]) => {
    setStudents([...studentss, ...newData]); // Cập nhật danh sách sinh viên
  };

  const table = useReactTable({
    data: studentss,
    columns: columns,
    getCoreRowModel: getCoreRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
    onSortingChange: setSorting,
    getSortedRowModel: getSortedRowModel(),
    onColumnFiltersChange: setColumnFilters,
    getFilteredRowModel: getFilteredRowModel(),
    state: {
      sorting,
      columnFilters,
    },
  });

  return (
    <>
      <h1 className="mb-4 text-2xl font-bold w-full">Students</h1>
      <CreateStudentList onImport={handleImport} />
      <div className="flex justify-between items-center gap-2">
        <TableFilterRow table={table} by="fullName" />
        <TablePagination table={table} />
      </div>
      <TableData table={table} />
    </>
  );
}

import { Input } from "@/components/ui/input";
import Papa from "papaparse";

function CreateStudentList({
  onImport,
}: {
  onImport: (data: Student[]) => void;
}) {
  const handleFileUpload = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];
    if (!file) return;

    const reader = new FileReader();
    reader.onload = (e) => {
      const content = e.target?.result as string;

      if (file.type === "application/json") {
        try {
          const jsonData: Student[] = JSON.parse(content);
          onImport(jsonData);
        } catch (error) {
          console.error("Invalid JSON file", error);
        }
      } else if (file.type === "text/csv") {
        Papa.parse<Student>(content, {
          header: true,
          skipEmptyLines: true,
          complete: (result) => onImport(result.data),
        });
      } else {
        console.error("Unsupported file format");
      }
    };

    reader.readAsText(file);
  };

  return (
    <div className="flex justify-end items-center gap-2">
      <Input
        type="file"
        accept=".json,.csv"
        onChange={handleFileUpload}
        className="hidden"
        id="file-upload"
      />
      <label htmlFor="file-upload">
        <Button variant="outline" onClick={() => console.log("Add Student")}>
          Import Students
        </Button>
      </label>
      <Button>Add Student</Button>
    </div>
  );
}

const ColumnFilter = ({
  header,
  column,
  options,
}: {
  header: string;
  column: Column<any, unknown>;
  options: string[];
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
          <SelectItem key={option} value={option}>
            {option}
          </SelectItem>
        ))}
      </SelectContent>
    </Select>
  );
};
