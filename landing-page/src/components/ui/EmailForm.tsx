"use client";

import { Button, Input, HStack } from "@chakra-ui/react";
import { Field } from "@/components/ui/field";
import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { object, string } from "yup";

const emailFormSchema = object({
  email: string().email("Invalid email").required("Email is required"),
})

export const EmailForm = () => {
  const { register, handleSubmit, formState: { errors } } = useForm({
    resolver: yupResolver(emailFormSchema),
    defaultValues: {
      email: "",
    },
  })

  const onSubmit = (values: any) => console.log(values)

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <HStack gap="12">
        <Field
          invalid={!!errors.email}
          errorText={errors.email?.message}
        >
          <Input
            {...register("email")}
            placeholder="Enter email address"
          />
        </Field>
        <Button type="submit">Get Early Access</Button>
      </HStack>
    </form>
  )
}
