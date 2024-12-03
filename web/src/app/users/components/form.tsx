"use client";

import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { Input, Stack } from "@chakra-ui/react";
import { Field } from "@/components/ui/field";
import { Button } from "@/components/ui/button";
import { userDetailsSchema } from "./schema";

interface Props {
  updating: boolean
  saveDetails: (values: any) => void
}

export default function UserDetailsForm({ updating, saveDetails }: Props) {
  const {
    control,
    register,
    handleSubmit,
    formState: { errors },
  } = useForm({
    resolver: yupResolver(userDetailsSchema),
    defaultValues: {
      govtid: "",
      firstname: "",
      lastname: "",
    },
  })

  const onSubmit = (values: any) => {
    if (!updating) {
      try {
        saveDetails(values)
      } catch (e) {
        console.error(e)
      }
    }
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <Stack gap="4" align="flex-start" maxW="sm">
        <Field
          required
          label="Govt Issued ID"
          invalid={!!errors.govtid}
          errorText={errors.govtid?.message}
        >
          <Input
            {...register("govtid", { required: "Govt id required" })}
          />
        </Field>
        <Field
          required
          label="First name"
          invalid={!!errors.firstname}
          errorText={errors.firstname?.message}
        >
          <Input
            {...register("firstname", { required: "First name required" })}
          />
        </Field>
        <Field
          required
          label="Last name"
          invalid={!!errors.lastname}
          errorText={errors.lastname?.message}
        >
          <Input
            {...register("lastname", { required: "Last name required" })}
          />
        </Field>
        <Button loading={updating} type="submit">
          Save
        </Button>
      </Stack>
    </form>
  )
}
