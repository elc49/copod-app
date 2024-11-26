import { z } from "zod";

export const landDetailsSchema = z.object({
  titleId: z.string(),
  size: z.string(),
  unit: z.string({
    required_error: "Select land measurement unit",
  }).array(),
})
