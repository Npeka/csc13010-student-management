import { FormFaculty } from "@/components/student/form/form-faculty";
import { FormProgram } from "@/components/student/form/form-program";
import { FormStatus } from "@/components/student/form/form-status";
import { TableFaculty } from "@/components/student/table/table-faculty";
import { TableProgram } from "@/components/student/table/table-program";
import { TableStatus } from "@/components/student/table/table-status";
import { PageTitle } from "../page-title";

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";

export default function AdminCategoriesPage() {
  return (
    <>
      <PageTitle title="Categories" />
      <div className="grid grid-cols-3 gap-4">
        <div className="flex flex-col gap-4">
          <CardLayout
            title="Faculty"
            form={<FormFaculty />}
            table={<TableFaculty />}
          />
        </div>
        <div className="flex flex-col gap-4">
          <CardLayout
            title="Program"
            form={<FormProgram />}
            table={<TableProgram />}
          />
        </div>
        <div className="flex flex-col gap-4">
          <CardLayout
            title="Status"
            form={<FormStatus />}
            table={<TableStatus />}
          />
        </div>
      </div>
    </>
  );
}

const CardLayout = ({
  title,
  form,
  table,
}: {
  title: string;
  form: React.ReactNode;
  table: React.ReactNode;
}) => {
  return (
    <Card>
      <CardHeader>
        <CardTitle>
          <h2 className="text-xl font-semibold mb-3">{title}</h2>
          {form}
        </CardTitle>
      </CardHeader>
      <CardContent>{table}</CardContent>
    </Card>
  );
};
