"use client";

import { Controller, useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { Input, Stack } from "@chakra-ui/react";
import { Field } from "@/components/ui/field";
import { Button } from "@/components/ui/button";
import {
  SelectRoot,
  SelectContent,
  SelectTrigger,
  SelectValueText,
  SelectItem,
} from "@/components/ui/select";
import { userDetailsSchema, verifications } from "./schema";

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
      verification: undefined,
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
        <Field
          required
          label="Verification"
          invalid={!!errors.verification}
          errorText={errors.verification?.message}
        >
          <Controller
            control={control}
            name="verification"
            render={({ field }) => (
              <SelectRoot
                name={field.name}
                value={field.value}
                onValueChange={({ value }) => field.onChange(value)}
                onInteractOutside={() => field.onBlur()}
                collection={verifications}
              >
                <SelectTrigger>
                  <SelectValueText placeholder="Verification status" />
                </SelectTrigger>
                <SelectContent>
                  {verifications.items.map((item) => (
                    <SelectItem item={item} key={item.value}>{item.label}</SelectItem>
                  ))}
                </SelectContent>
              </SelectRoot>
            )}
          />
        </Field>
        <Button loading={updating} type="submit">
          Save
        </Button>
      </Stack>
    </form>
  )
}
