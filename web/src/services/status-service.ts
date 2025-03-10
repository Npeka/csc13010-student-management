"use client";
import { appApi } from "@/services/config";
import { Status } from "@/types/student";
import { Response } from "@/types/response";

const statusApi = appApi.injectEndpoints({
  overrideExisting: true,
  endpoints: (builder) => ({
    getStatuses: builder.query<Response<Status[]>, void>({
      query: () => ({
        url: "/api/v1/statuses/",
        method: "GET",
      }),
      providesTags: ["Status"],
    }),

    createStatus: builder.mutation<Response<Status>, Status>({
      query: (status) => ({
        url: "/api/v1/statuses/",
        method: "POST",
        body: status,
      }),
      invalidatesTags: ["Status"],
    }),

    updateStatus: builder.mutation<Response<Status>, Status>({
      query: (status) => ({
        url: `/api/v1/statuses/${status.id}`,
        method: "PUT",
        body: { name: status.name },
      }),
      invalidatesTags: ["Status"],
    }),

    deleteStatus: builder.mutation<void, string>({
      query: (status_id) => ({
        url: `/api/v1/statuses/${status_id}`,
        method: "DELETE",
      }),
      invalidatesTags: ["Status"],
    }),
  }),
});

export const {
  useGetStatusesQuery,
  useCreateStatusMutation,
  useUpdateStatusMutation,
  useDeleteStatusMutation,
} = statusApi;
