"use client";
import { FormStudent } from "./form-student";
import {
  useGetStudentByIdQuery,
  useGetStudentOptionsQuery,
} from "@/services/student-service";

export const FormStudentData = ({ student_id }: { student_id: string }) => {
  const { data: student } = useGetStudentByIdQuery(student_id);
  const { data: options } = useGetStudentOptionsQuery();

  if (!student || !options) {
    return <div>Loading...</div>;
  }

  return (
    <div>
      <FormStudent student={student.data} options={options.data} />
    </div>
  );
};
