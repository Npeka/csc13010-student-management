"use client";
import { appApi } from "@/services/config";
import { Program } from "@/types/student";
import { Response } from "@/types/response";

const programApi = appApi.injectEndpoints({
  overrideExisting: true,
  endpoints: (builder) => ({
    getPrograms: builder.query<Response<Program[]>, void>({
      query: () => ({
        url: "/api/v1/programs/",
        method: "GET",
      }),
      providesTags: ["Program"],
    }),

    createProgram: builder.mutation<Response<Program>, Program>({
      query: (program) => ({
        url: "/api/v1/programs/",
        method: "POST",
        body: program,
      }),
      invalidatesTags: ["Program"],
    }),

    updateProgram: builder.mutation<Response<Program>, Program>({
      query: (program) => ({
        url: `/api/v1/programs/${program.id}`,
        method: "PUT",
        body: { name: program.name },
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
  useUpdateProgramMutation,
  useDeleteProgramMutation,
} = programApi;
