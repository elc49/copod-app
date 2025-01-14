"use client";

import { useState } from "react";
import { Input, HStack } from "@chakra-ui/react";
import { Button } from "@/components/ui/button";
import { Field } from "@/components/ui/field";
import { toaster } from "@/components/ui/toaster";
import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { object, string } from "yup";

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

  const onSubmit = async (values) => {
    if (!submitting) {
      setSubmitting(true)
      try {
        await fetch("https://boss-freely-koi.ngrok-free.app/api/early", {
          method: "POST",
          mode: "no-cors",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(values.email),
        })
        toaster.create({
          type: "success",
          duration: 4000,
          title: "Email received",
        })
      } catch (e) {
        console.error(e)
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
