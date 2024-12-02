"use client";

import { useMemo } from "react";
import { Controller, useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { createListCollection, Input, Stack } from "@chakra-ui/react";
import {
  SelectContent,
  SelectItem,
  SelectRoot,
  SelectTrigger,
  SelectValueText,
} from "@/components/ui/select";
import { landDetailsSchema } from "./schema";
import { Button } from "@/components/ui/button";
import { Field } from "@/components/ui/field";

interface Props {
  registerLand: (titleId: string, size: number, unit: string) => void
  registering: boolean
}

function LandDetails({ registerLand, registering }: Props) {
  const {
    control,
    register,
    handleSubmit,
    formState: { errors },
  }= useForm({
    resolver: yupResolver(landDetailsSchema),
    defaultValues: {
      titleId: "",
      size: "",
      unit: [],
    },
  })
  const sizeUnits = useMemo(() => {
    return createListCollection({
      items: [
        {
          label: "acres",
          value: "acres",
        },
        {
          label: "Hectares",
          value: "ha",
        },
      ]
    })
  }, [])

  const onSubmit = (values: any) => {
    try {
      registerLand(values.titleId, values.size, values.unit[0])
    } catch (e) {
      console.error(e)
    }
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <Stack gap="4" align="flex-start" maxW="sm">
        <Field
          label="Title number"
          invalid={!!errors.titleId}
          errorText={errors.titleId?.message}
        >
          <Input
            {...register("titleId", { required: "Land title required" })}
          />
        </Field>
        <Field
          label="Size"
          invalid={!!errors.size}
          errorText={errors.size?.message}
        >
          <Input {...register("size", { required: "Land size required" })}
          />
        </Field>
        <Field
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
                collection={sizeUnits}
              >
                <SelectTrigger>
                  <SelectValueText placeholder="Select land size unit" />
                </SelectTrigger>
                <SelectContent>
                  {sizeUnits.items.map((item) => (
                    <SelectItem item={item.value} key={item.value}>{item.label}</SelectItem>
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
