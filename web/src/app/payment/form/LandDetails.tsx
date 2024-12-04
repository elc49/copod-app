"use client";

import { Controller, useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { Input, Stack } from "@chakra-ui/react";
import {
  SelectContent,
  SelectItem,
  SelectRoot,
  SelectTrigger,
  SelectValueText,
} from "@/components/ui/select";
import { landDetailsSchema, units, status } from "./schema";
import { Button } from "@/components/ui/button";
import { Field } from "@/components/ui/field";

interface Props {
  registerLand: (titleId: string, size: number, unit: string, status: string) => void
  registering: boolean
}

function LandDetails({ registerLand, registering }: Props) {
  const {
    control,
    register,
    handleSubmit,
    formState: { errors },
  } = useForm({
    resolver: yupResolver(landDetailsSchema),
    defaultValues: {
      titleId: "",
      size: "",
      unit: undefined,
      status: undefined,
    },
  })

  const onSubmit = (values: any) => {
    if (!registering) {
      try {
        registerLand(values.titleId, values.size, values.unit[0], values.status[0])
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
          label="Title number"
          invalid={!!errors.titleId}
          errorText={errors.titleId?.message}
        >
          <Input
            {...register("titleId")}
          />
        </Field>
        <Field
          required
          label="Size"
          invalid={!!errors.size}
          errorText={errors.size?.message}
        >
          <Input {...register("size")}
          />
        </Field>
        <Field
          required
          label="Unit"
          invalid={!!errors.unit}
          errorText={errors.unit?.message}
        >
          <Controller
            control={control}
            name="unit"
            render={({ field }) => (
              <SelectRoot
                name={field.name}
                value={field.value}
                onValueChange={({ value }) => field.onChange(value)}
                onInteractOutside={() => field.onBlur()}
                collection={units}
              >
                <SelectTrigger>
                  <SelectValueText placeholder="Select land size unit" />
                </SelectTrigger>
                <SelectContent>
                  {units.items.map((item) => (
                    <SelectItem item={item} key={item.value}>{item.label}</SelectItem>
                  ))}
                </SelectContent>
              </SelectRoot>
            )}
          />
        </Field>
        <Field
          required
          label="Status"
          invalid={!!errors.status}
          errorText={errors.status?.message}
        >
          <Controller
            control={control}
            name="status"
            render={({ field }) => (
              <SelectRoot
                name={field.name}
                value={field.value}
                onValueChange={({ value }) => field.onChange(value)}
                onInteractOutside={() => field.onBlur()}
                collection={status}
              >
                <SelectTrigger>
                  <SelectValueText placeholder="Verification status" />
                </SelectTrigger>
                <SelectContent>
                  {status.items.map((item) => (
                    <SelectItem item={item} key={item.value}>{item.label}</SelectItem>
                  ))}
                </SelectContent>
              </SelectRoot>
            )}
          />
        </Field>
        <Button loading={registering} type="submit">
          Send
        </Button>
      </Stack>
    </form>
  )
}

export default LandDetails
