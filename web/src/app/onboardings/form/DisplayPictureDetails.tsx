"use client";

import { Controller, useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { Stack } from "@chakra-ui/react";
import { Field } from "@/components/ui/field";
import { Button } from "@/components/ui/button";
import {
  SelectRoot,
  SelectContent,
  SelectTrigger,
  SelectValueText,
  SelectItem,
} from "@/components/ui/select";
import { displayPictureDetailsSchema, verifications } from "./schema";

interface Props {
  saveDp: (values: any) => void
  savingDp: boolean
}

export default function DisplayPictureDetails({ saveDp, savingDp }: Props) {
  const { control, handleSubmit } = useForm({
    resolver: yupResolver(displayPictureDetailsSchema),
    defaultValues: {
      verification: undefined,
    },
  })

  const onSubmit = (values: any) => {
    saveDp(values)
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <Stack gap="4" align="flex-start" maxW="sm">
        <Field
          required
          label="verification"
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
        <Button loading={savingDp} type="submit">
          Save
        </Button>
      </Stack>
    </form>
  )
}
