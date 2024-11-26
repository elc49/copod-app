"use client";

import { useMemo } from "react";
import { Controller, useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
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

function LandDetails() {
  const {
    control,
    register,
    handleSubmit,
    formState: { errors },
  }= useForm<z.infer<typeof landDetailsSchema>>({
    resolver: zodResolver(landDetailsSchema),
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

  const onSubmit = (values: z.infer<typeof landDetailsSchema>) => {
    console.log(values)
  }

  return (
    <form onSubmit={onSubmit}>
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
        <Button onClick={() => {}}>
          Send
        </Button>
      </Stack>
    </form>
  )
}

export default LandDetails
