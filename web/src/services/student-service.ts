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
import { fetchBaseQuery } from "@reduxjs/toolkit/query";

const studentApi = appApi.injectEndpoints({
  overrideExisting: true,
  endpoints: (builder) => ({
    getAllStudents: builder.query<StudentResponseDTO[], void>({
      query: () => ({
        url: "/api/v1/students/",
        method: "GET",
      }),
      providesTags: ["Student"],
    }),

    getStudentById: builder.query<StudentResponseDTO, string>({
      query: (id) => ({
        url: `/api/v1/students/${id}`,
        method: "GET",
      }),
      providesTags: ["Student"],
    }),

    createStudent: builder.mutation<StudentResponseDTO, Student>({
      query: (student) => ({
        url: "/api/v1/students",
        method: "POST",
        body: student,
      }),
      invalidatesTags: ["Student"],
    }),

    updateStudent: builder.mutation<
      StudentResponseDTO,
      { id: string; student: Student }
    >({
      query: ({ id, student }) => ({
        url: `/api/v1/students/${id}`,
        method: "PATCH",
        body: student,
      }),
      invalidatesTags: ["Student"],
    }),

    deleteStudent: builder.mutation<void, string>({
      query: (student_id) => ({
        url: `/api/v1/students/${student_id}`,
        method: "DELETE",
      }),
      invalidatesTags: ["Student"],
    }),

    getStudentOptions: builder.query<OptionDTO, void>({
      query: () => ({
        url: "/api/v1/students/options",
        method: "GET",
      }),
      providesTags: ["Student"],
    }),

    getFaculties: builder.query<Faculty[], void>({
      query: () => ({
        url: "/api/v1/students/faculties",
        method: "GET",
      }),
      providesTags: ["Faculty"],
    }),

    getPrograms: builder.query<Program[], void>({
      query: () => ({
        url: "/api/v1/students/programs",
        method: "GET",
      }),
      providesTags: ["Program"],
    }),

    getStatuses: builder.query<Status[], void>({
      query: () => ({
        url: "/api/v1/students/statuses",
        method: "GET",
      }),
      providesTags: ["Status"],
    }),

    createFaculty: builder.mutation<Faculty, Faculty>({
      query: (faculty) => ({
        url: "/api/v1/students/faculties",
        method: "POST",
        body: faculty,
      }),
      invalidatesTags: ["Faculty"],
    }),

    createProgram: builder.mutation<Program, Program>({
      query: (program) => ({
        url: "/api/v1/students/programs",
        method: "POST",
        body: program,
      }),
      invalidatesTags: ["Program"],
    }),

    createStatus: builder.mutation<Status, Status>({
      query: (status) => ({
        url: "/api/v1/students/statuses",
        method: "POST",
        body: status,
      }),
      invalidatesTags: ["Status"],
    }),

    deleteFaculty: builder.mutation<void, string>({
      query: (faculty_id) => ({
        url: `/api/v1/students/faculties/${faculty_id}`,
        method: "DELETE",
      }),
      invalidatesTags: ["Faculty"],
    }),

    deleteProgram: builder.mutation<void, string>({
      query: (program_id) => ({
        url: `/api/v1/students/programs/${program_id}`,
        method: "DELETE",
      }),
      invalidatesTags: ["Program"],
    }),

    deleteStatus: builder.mutation<void, string>({
      query: (status_id) => ({
        url: `/api/v1/students/statuses/${status_id}`,
        method: "DELETE",
      }),
      invalidatesTags: ["Status"],
    }),
  }),
});

export const {
  useGetAllStudentsQuery,
  useGetStudentByIdQuery,
  useCreateStudentMutation,
  useUpdateStudentMutation,
  useDeleteStudentMutation,
  useGetStudentOptionsQuery,

  useGetFacultiesQuery,
  useGetProgramsQuery,
  useGetStatusesQuery,

  useCreateFacultyMutation,
  useCreateProgramMutation,
  useCreateStatusMutation,

  useDeleteFacultyMutation,
  useDeleteProgramMutation,
  useDeleteStatusMutation,
} = studentApi;
