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

export default function AdminCategoriesPage({
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

  return (
    <>
      <h1>AdminCategories</h1>
    </>
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
