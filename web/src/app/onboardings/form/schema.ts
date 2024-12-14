import * as yup from "yup";
import { createListCollection } from "@chakra-ui/react";
import { Verification } from "@/graphql/graphql";

export const landDetailsSchema = yup.object({
  titleId: yup.string().required(),
  size: yup.string().required(),
  unit: yup.array().required(),
  status: yup.array().required(),
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


