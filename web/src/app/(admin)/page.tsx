"use client";

import { PageTitle } from "./page-title";
import {
  useGetAllStudentsQuery,
  useGetStudentOptionsQuery,
} from "@/services/student-service";
import { StudentTable } from "../../components/student/table/table-student";

export default function AdminStudentsPage() {
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

  const { data: students } = useGetAllStudentsQuery();
  const { data: options } = useGetStudentOptionsQuery();

  if (!students || !options) {
    return <div>Loading...</div>;
  }

  const data = students.map((student) => ({
    ...student,
    birth_date: new Date(student.birth_date).toLocaleDateString(),
    gender:
      options?.genders.find((gender) => gender.id === student.gender_id)
        ?.name ?? "Unknown",
    faculty:
      options?.faculties.find((faculty) => faculty.id === student.faculty_id)
        ?.name ?? "Unknown",
    course:
      options?.courses.find((course) => course.id === student.course_id)
        ?.name ?? "Unknown",
    program:
      options?.programs.find((program) => program.id === student.program_id)
        ?.name ?? "Unknown",
    status:
      options?.statuses.find((status) => status.id === student.status_id)
        ?.name ?? "Unknown",
  }));

  return (
    <div className="space-y-4">
      <PageTitle title="Students" />
      <StudentTable data={data ?? []} options={options ?? []} />
    </div>
  );
}
