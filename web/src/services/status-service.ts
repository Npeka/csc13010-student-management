"use client";
import { appApi } from "@/services/config";
import { Status } from "@/types/student";

const statusApi = appApi.injectEndpoints({
  overrideExisting: true,
  endpoints: (builder) => ({
    getStatuses: builder.query<Status[], void>({
      query: () => ({
        url: "/api/v1/statuses/",
        method: "GET",
      }),
      providesTags: ["Status"],
    }),

    createStatus: builder.mutation<Status, Status>({
      query: (status) => ({
        url: "/api/v1/statuses/",
        method: "POST",
        body: status,
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
  useDeleteStatusMutation,
} = statusApi;
