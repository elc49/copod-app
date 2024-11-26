import * as yup from "yup";

export const landDetailsSchema = yup.object({
  titleId: yup.string().required(),
  size: yup.string().required(),
  unit: yup.array().required(),
})
