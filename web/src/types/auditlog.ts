export interface AuditLog {
  id: number;
  table_name: string;
  record_id: number;
  action: Action;
  changed_fields: string;
  changed_by: Role;
  created_at: Date;
}

export type Action = "CREATE" | "UPDATE" | "DELETE";

export type Role = "ADMIN" | "STUDENT";
