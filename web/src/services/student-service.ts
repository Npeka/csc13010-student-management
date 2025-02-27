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
  }),
});

export const {
  useGetAllStudentsQuery,
  useGetStudentByIdQuery,
  useCreateStudentMutation,
  useUpdateStudentMutation,
  useDeleteStudentMutation,
  useGetStudentOptionsQuery,
} = studentApi;
