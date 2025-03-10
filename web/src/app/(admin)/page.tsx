"use client";

import { PageTitle } from "./page-title";
import {
  useGetAllStudentsQuery,
  useGetStudentOptionsQuery,
} from "@/services/student-service";
import { StudentTable } from "../../components/student/table/table-student";
import { OptionDTO } from "@/types/student";

export default function AdminStudentsPage() {
  const { data: studentsResponse, isLoading: isLoadingStudents } =
    useGetAllStudentsQuery();
  const { data: optionsResponse, isLoading: isLoadingOptions } =
    useGetStudentOptionsQuery();
  const students = studentsResponse?.data;
  const options: OptionDTO = optionsResponse?.data ?? {
    genders: [],
    faculties: [],
    courses: [],
    programs: [],
    statuses: [],
  };

  if (isLoadingStudents || isLoadingOptions) {
    return <div>Loading...</div>;
  }

  return (
    <div className="space-y-4">
      <PageTitle title="Students" />
      <StudentTable data={students ?? []} options={options ?? {}} />
    </div>
  );
}
