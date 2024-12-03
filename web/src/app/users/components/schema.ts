import * as yup from "yup";

export const userDetailsSchema = yup.object({
  govtid: yup.string().required(),
  firstname: yup.string().required(),
  lastname: yup.string().required(),
})
