import { useMemo } from "react";
import { SupportingDoc } from "@/graphql/graphql";
import {
  Table,
} from "@chakra-ui/react";
import {
  useReactTable,
  flexRender,
  createColumnHelper,
  getCoreRowModel,
} from "@tanstack/react-table";
import { useRouter } from "next/navigation";
import { DoneIcon, WaitingIcon, FailIcon } from "@/components/icons";

interface Props {
  supportingDocs: SupportingDoc[]
}

const columnHelper = createColumnHelper<SupportingDoc>()

export default function SupportingDocsVerificationTable(props: Props) {
  const { supportingDocs } = props
  const renderStatusColumn = (status: string) => {
    switch (status) {
      case "ONBOARDING":
        return <WaitingIcon />
      case "VERIFIED":
        return <DoneIcon />
      case "REJECTED":
        return <FailIcon />
    }
  }
  const columns = useMemo(() => {
    return [
      columnHelper.accessor("id", {
        cell: info => (
          <div>{info.getValue()}</div>
        ),
        header: () => <span>#</span>
      }),
      columnHelper.accessor("verified", {
        cell: info => renderStatusColumn(info.getValue()),
        header: () => <span>Verification</span>
      }),
    ]
  }, [])
  const table = useReactTable({
    data: supportingDocs,
    columns,
    getCoreRowModel: getCoreRowModel(),
  })
  const router = useRouter()

  return (
    <Table.ScrollArea height="100%" rounded="md" borderWidth="1px">
      <Table.Root variant="outline" size="lg" stickyHeader interactive>
        <Table.Header>
          {table.getHeaderGroups().map((headerGroup) => (
            <Table.Row key={headerGroup.id}>
              {headerGroup.headers.map((header) => (
                <Table.ColumnHeader key={header.id}>
                  {header.isPlaceholder
                    ? null
                    : flexRender(
                        header.column.columnDef.header,
                        header.getContext(),
                      )}
                </Table.ColumnHeader>
              ))}
            </Table.Row>
          ))}
        </Table.Header>
        <Table.Body>
          {table.getRowModel().rows?.length ? (
            table.getRowModel().rows.map((row) => (
              <Table.Row
               key={row.id}
               onClick={() => router.push(`users/${row.original.id}`)}
              >
                {row.getVisibleCells().map((cell) => (
                  <Table.Cell key={cell.id}>
                    {flexRender(
                      cell.column.columnDef.cell,
                      cell.getContext(),
                    )}
                  </Table.Cell>
                ))}
              </Table.Row>))
          ) : (
            <Table.Row>
              <Table.Cell colSpan={columns.length} className="h-24 text-center">No results.</Table.Cell>
            </Table.Row>
          )}
        </Table.Body>
      </Table.Root>
    </Table.ScrollArea>
  )
}
