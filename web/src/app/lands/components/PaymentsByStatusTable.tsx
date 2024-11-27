import { useMemo } from "react";
import { Payment } from "@/graphql/graphql";
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
import { Tag } from "@/components/ui/tag";

interface Props {
  payments: Payment[]
}

const columnHelper = createColumnHelper<Payment>()


export default function PaymentsByStatusTable(props: Props) {
  const { payments } = props
  const columns = useMemo(() => {
    return [
      columnHelper.accessor("reference_id", {
        cell: info => (
          <div>{info.getValue()}</div>
        ),
        header: () => <span>#</span>
      }),
      columnHelper.accessor("status", {
        cell: info => <Tag textTransform="capitalize">{info.getValue()}</Tag>,
        header: () => <span>Payment</span>
      }),
      columnHelper.accessor("title.verified", {
        cell: info => <Tag>{info.getValue()}</Tag>,
        header: () => <span>Verification</span>
      }),
    ]
  }, [])
  const table = useReactTable({
    data: payments,
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
               onClick={() => router.push(`payment/${row.original.id}`)}
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
