"use client";
import { appApi } from "@/services/config";
import { Program } from "@/types/student";

const programApi = appApi.injectEndpoints({
  overrideExisting: true,
  endpoints: (builder) => ({
    getPrograms: builder.query<Program[], void>({
      query: () => ({
        url: "/api/v1/programs/",
        method: "GET",
      }),
      providesTags: ["Program"],
    }),

    createProgram: builder.mutation<Program, Program>({
      query: (program) => ({
        url: "/api/v1/programs/",
        method: "POST",
        body: program,
      }),
      invalidatesTags: ["Program"],
    }),

    deleteProgram: builder.mutation<void, string>({
      query: (program_id) => ({
        url: `/api/v1/programs/${program_id}`,
        method: "DELETE",
      }),
      invalidatesTags: ["Program"],
    }),
  }),
});

export const {
  useGetProgramsQuery,
  useCreateProgramMutation,
  useDeleteProgramMutation,
} = programApi;
