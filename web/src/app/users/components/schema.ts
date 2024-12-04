import * as yup from "yup";
import { Verification } from "@/graphql/graphql";
import { createListCollection } from "@chakra-ui/react";

export const userDetailsSchema = yup.object({
  govtid: yup.string().required(),
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
