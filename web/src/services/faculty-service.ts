"use client";
import { appApi } from "@/services/config";
import {
  Student,
  StudentResponseDTO,
  OptionDTO,
  Faculty,
  Program,
  Status,
} from "@/types/student";

const facultyApi = appApi.injectEndpoints({
  overrideExisting: true,
  endpoints: (builder) => ({
    getFaculties: builder.query<Faculty[], void>({
      query: () => ({
        url: "/api/v1/faculties/",
        method: "GET",
      }),
      providesTags: ["Faculty"],
    }),

    createFaculty: builder.mutation<Faculty, Faculty>({
      query: (faculty) => ({
        url: "/api/v1/faculties/",
        method: "POST",
        body: faculty,
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
  useDeleteFacultyMutation,
} = facultyApi;
