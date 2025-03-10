"use client";
import { appApi } from "@/services/config";
import { AuditLog } from "@/types/auditlog";
import { Response } from "@/types/response";

const auditlogApi = appApi.injectEndpoints({
  overrideExisting: true,
  endpoints: (builder) => ({
    getAuditLogs: builder.query<Response<AuditLog[]>, void>({
      query: () => "/api/v1/auditlogs/",
    }),

    getModelAuditLogs: builder.query<
      AuditLog[],
      { tableName: string; recordId: string }
    >({
      query: ({ tableName, recordId }) =>
        `/api/v1/auditlogs/${tableName}/${recordId}`,
    }),
  }),
});

export const { useGetAuditLogsQuery, useGetModelAuditLogsQuery } = auditlogApi;
