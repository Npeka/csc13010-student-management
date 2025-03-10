"use client";

import { appApi } from "@/services/config";
import { Config } from "@/types/config";
import { Response } from "@/types/response";

const configApi = appApi.injectEndpoints({
  overrideExisting: true,
  endpoints: (builder) => ({
    getConfig: builder.query<Response<Config>, void>({
      query: () => ({
        url: "/api/v1/config/",
        method: "GET",
      }),
      providesTags: ["Config"],
    }),

    updateConfig: builder.mutation<Response<Config>, Config>({
      query: (config) => ({
        url: "/api/v1/config/",
        method: "PUT",
        body: config,
      }),
      invalidatesTags: ["Config"],
    }),
  }),
});

export const { useGetConfigQuery, useUpdateConfigMutation } = configApi;
