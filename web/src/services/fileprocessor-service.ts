"use client";
import { appApi } from "@/services/config";

const fileprocessorApi = appApi.injectEndpoints({
  overrideExisting: true,
  endpoints: (builder) => ({
    importFile: builder.mutation<
      void,
      { file: File; format: string; module: string }
    >({
      query: ({ file, format, module }) => {
        const formData = new FormData();
        formData.append("file", file);

        return {
          url: `/api/v1/fileprocessor/import?format=${format}&module=${module}`,
          method: "POST",
          body: formData,
          headers: {},
        };
      },
      invalidatesTags: ["FileProcessor"],
    }),

    exportFile: builder.query<Blob, { format: string; module: string }>({
      query: ({ format, module }) => ({
        url: `/api/v1/fileprocessor/export?format=${format}&module=${module}`,
        method: "GET",
        responseHandler: async (response) => {
          const blob = await response.blob();
          return blob;
        },
      }),
    }),
  }),
});

export const {
  useImportFileMutation,
  useExportFileQuery,
  useLazyExportFileQuery,
} = fileprocessorApi;
