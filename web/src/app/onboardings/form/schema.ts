import * as yup from "yup";
import { createListCollection } from "@chakra-ui/react";
import { Verification } from "@/graphql/graphql";

export const landDetailsSchema = yup.object({
  titleId: yup.string().required(),
  size: yup.string().required(),
  unit: yup.array().required(),
  status: yup.array().required(),
  registration: yup.date().required(),
})

export const units = createListCollection({
  items: [
    {
      label: "acres",
      value: "acres",
    },
    {
      label: "Hectares",
      value: "ha",
    },
  ],
})

export const status = createListCollection({
  items: [
    {
      label: "Verify",
      value: Verification.Verified,
    },
    {
      label: "Reject",
      value: Verification.Rejected,
    },
  ],
})

export const userDetailsSchema = yup.object({
  firstname: yup.string().required(),
  lastname: yup.string().required(),
  verification: yup.array().required(),
})

export const verifications = createListCollection({
  items: [
    {
      label: "Verify",
      value: Verification.Verified,
    },
    {
      label: "Reject",
      value: Verification.Rejected,
    },
  ],
})

export const displayPictureDetailsSchema = yup.object({
  verification: yup.array().required(),
})
