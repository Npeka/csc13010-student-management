export interface AuditLog {
  id: number;
  table_name: string;
  record_id: number;
  action: Action;
  field_changes: string;
  old_record: string;
  new_record: string;
  lsn: number;
  created_at: Date;
}

export type Action = "CREATE" | "UPDATE" | "DELETE";

export type Role = "ADMIN" | "STUDENT";
