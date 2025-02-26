import { FormStudentData } from "./form-student-data";
import { PageTitle } from "../page-title";

export default async function StudentPage({
  params,
}: {
  params: Promise<{ student_id: string }>;
}) {
  const student_id = await params;
  return (
    <div>
      <PageTitle title={`Student ${student_id.student_id}`} />
      <FormStudentData student_id={student_id.student_id} />
    </div>
  );
}
