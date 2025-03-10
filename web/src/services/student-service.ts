"use client";
import { appApi } from "@/services/config";
import { Student, StudentResponseDTO, OptionDTO } from "@/types/student";
import { Response } from "@/types/response";

const studentApi = appApi.injectEndpoints({
  overrideExisting: true,
  endpoints: (builder) => ({
    getAllStudents: builder.query<Response<Student[]>, void>({
      query: () => ({
        url: "/api/v1/students/",
        method: "GET",
      }),
      providesTags: ["Student", "FileProcessor"],
    }),

    getStudentById: builder.query<Response<Student>, string>({
      query: (id) => ({
        url: `/api/v1/students/full/${id}`,
        method: "GET",
      }),
      providesTags: ["Student"],
    }),

    createStudent: builder.mutation<Response<Student>, Student>({
      query: (student) => ({
        url: "/api/v1/students",
        method: "POST",
        body: student,
      }),
      invalidatesTags: ["Student"],
    }),

    updateStudent: builder.mutation<
      Response<Student>,
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

    getStudentOptions: builder.query<Response<OptionDTO>, void>({
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
