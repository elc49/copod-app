"use client";

import { useState } from "react";
import { Input, HStack } from "@chakra-ui/react";
import { Button } from "@/components/ui/button";
import { Field } from "@/components/ui/field";
import { toaster } from "@/components/ui/toaster";
import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { object, string } from "yup";
import { sendGAEvent } from "@next/third-parties/google";

const emailFormSchema = object({
  email: string().email("Invalid email").required("Email is required"),
})

export const EmailForm = () => {
  const [submitting, setSubmitting] = useState(false)
  const { register, handleSubmit, formState: { errors } } = useForm({
    resolver: yupResolver(emailFormSchema),
    defaultValues: {
      email: "",
    },
  })

  const onSubmit = async (values: any) => {
    if (!submitting) {
      setSubmitting(true)
      try {
        const res = await fetch(`${process.env.NEXT_PUBLIC_LOCAL_API}/early`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(values.email),
        })
        if (res.status === 200) {
          toaster.create({
            type: "success",
            duration: 4000,
            title: "Email received",
          })
          sendGAEvent("early_signup", "formSubmission")
        } else {
          toaster.create({
            type: "error",
            duration: 4000,
            title: "Something went wrong",
          })
        }
      } catch (e) {
        console.error(e)
        toaster.create({
          type: "error",
          duration: 4000,
          title: `${(e as Error).message}`,
        })
      } finally {
        setSubmitting(false)
      }
    }
  }

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
        <Button loading={submitting} type="submit">Get Early Access</Button>
      </HStack>
    </form>
  )
}
