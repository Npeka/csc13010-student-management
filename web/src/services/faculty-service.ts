"use client";
import { appApi } from "@/services/config";
import { Faculty } from "@/types/student";
import { Response } from "@/types/response";

const facultyApi = appApi.injectEndpoints({
  overrideExisting: true,
  endpoints: (builder) => ({
    getFaculties: builder.query<Response<Faculty[]>, void>({
      query: () => ({
        url: "/api/v1/faculties/",
        method: "GET",
      }),
      providesTags: ["Faculty"],
    }),

    createFaculty: builder.mutation<Response<Faculty>, Faculty>({
      query: (faculty) => ({
        url: "/api/v1/faculties/",
        method: "POST",
        body: faculty,
      }),
      invalidatesTags: ["Faculty"],
    }),

    updateFaculty: builder.mutation<Response<Faculty>, Faculty>({
      query: (faculty) => ({
        url: `/api/v1/faculties/${faculty.id}`,
        method: "PUT",
        body: { name: faculty.name },
      }),
      invalidatesTags: ["Faculty"],
    }),

    deleteFaculty: builder.mutation<void, string>({
      query: (faculty_id) => ({
        url: `/api/v1/faculties/${faculty_id}`,
        method: "DELETE",
      }),
      invalidatesTags: ["Faculty"],
    }),
  }),
});

export const {
  useGetFacultiesQuery,
  useCreateFacultyMutation,
  useUpdateFacultyMutation,
  useDeleteFacultyMutation,
} = facultyApi;
