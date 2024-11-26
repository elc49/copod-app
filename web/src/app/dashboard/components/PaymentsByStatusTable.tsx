import { useMemo } from "react";
import { Payment } from "@/graphql/graphql";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import {
  useReactTable,
  flexRender,
  createColumnHelper,
  getCoreRowModel,
} from "@tanstack/react-table";
import { Badge } from "@/components/ui/badge";
import { useRouter } from "next/navigation";

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
        cell: info => (
          <div className="capitalize">
            <Badge>
              {info.getValue()}
            </Badge>
          </div>
        ),
        header: () => <span>Payment</span>
      }),
      columnHelper.accessor("title.verified", {
        cell: info => (
          <div>
            <Badge
              variant={
                info.getValue() === "ONBOARDING"
                  ? `secondary`
                  : info.getValue() === "VERIFIED"
                    ? `default`
                    : info.getValue() === "REJECTED"
                      ? `destructive`
                      : `outline`
              }
            >
              {info.getValue()}
            </Badge>
          </div>
        ),
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
    <div className="w-full p-4">
      <div className="rounded-md border">
        <Table>
          <TableHeader>
            {table.getHeaderGroups().map((headerGroup) => (
              <TableRow key={headerGroup.id}>
                {headerGroup.headers.map((header) => (
                  <TableHead key={header.id}>
                    {header.isPlaceholder
                      ? null
                      : flexRender(
                          header.column.columnDef.header,
                          header.getContext(),
                        )}
                  </TableHead>
                ))}
              </TableRow>
            ))}
          </TableHeader>
          <TableBody>
            {table.getRowModel().rows?.length ? (
              table.getRowModel().rows.map((row) => (
                <TableRow
                 key={row.id}
                 onClick={() => router.push(`payment/${row.original.id}`)}
                >
                  {row.getVisibleCells().map((cell) => (
                    <TableCell key={cell.id}>
                      {flexRender(
                        cell.column.columnDef.cell,
                        cell.getContext(),
                      )}
                    </TableCell>
                  ))}
                </TableRow>))
            ) : (
              <TableRow>
                <TableCell colSpan={columns.length} className="h-24 text-center">No results.</TableCell>
              </TableRow>
            )}
          </TableBody>
        </Table>
      </div>
    </div>
  )
}
