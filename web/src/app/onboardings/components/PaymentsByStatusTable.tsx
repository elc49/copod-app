import { useMemo } from "react";
import NextImage from "next/image";
import { Payment } from "@/graphql/graphql";
import {
  HStack,
  IconButton,
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
import { DoneIcon, WaitingIcon, FailIcon, ViewIcon } from "@/components/icons";

interface Props {
  payments: Payment[]
}

const columnHelper = createColumnHelper<Payment>()


export default function PaymentsByStatusTable(props: Props) {
  const { payments } = props
  const router = useRouter()
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
  const renderPaymentStatus = (status: string) => {
    switch (status) {
      case "success":
        return <DoneIcon />
      case "failed":
        return <FailIcon />
      default:
        return <Tag>{status}</Tag>
    }
  }
  const renderDocImage = (url: string) => {
    return (
      <NextImage
        src={url}
        alt={"Land title"}
        priority={true}
        width={24}
        height={24}
      />
    )
  }
  const shouldWeDoSomething = (status: string) => {
    switch(status) {
      case "VERIFIED":
        return false
      default:
        return true
    }
  }
  const columns = useMemo(() => {
    return [
      columnHelper.accessor("reference_id", {
        cell: info => (
          <div>{info.getValue()}</div>
        ),
        header: () => <span># Reference</span>
      }),
      columnHelper.accessor("status", {
        cell: info => renderPaymentStatus(info.getValue()),
        header: () => <span>Status</span>
      }),
      columnHelper.accessor("onboarding.title.url", {
        cell: info => (
          <HStack>
            {shouldWeDoSomething(info.row.original.onboarding?.title.verified || "") ? (
              <IconButton
                size="xs"
                aria-label="Go back"
                onClick={() => router.push(`onboardings/document/${info.row.original?.onboarding?.title.id}/${info.row.original?.onboarding?.title.__typename?.toLowerCase()}`)}
            >
                <ViewIcon />
              </IconButton>
            ) : (<DoneIcon />)}
            {renderDocImage(info.getValue())}
          </HStack>
        ),
        header: () => <span>Title</span>,
      }),
      columnHelper.accessor("onboarding.supportingDoc.url", {
        cell: info => (
          <HStack>
            {shouldWeDoSomething(info.row.original.onboarding?.supportingDoc.verified || "") ? (
              <IconButton
                size="xs"
                aria-label="Go back"
                onClick={() => router.push(`onboardings/document/${info.row.original?.onboarding?.supportingDoc.id}/${info.row.original?.onboarding?.supportingDoc.__typename?.toLowerCase()}`)}
              >
                <ViewIcon />
              </IconButton>
            ) : (<DoneIcon />)}
            {renderDocImage(info.getValue())}
          </HStack>
        ),
        header: () => <span>ID</span>
      }),
      columnHelper.accessor("onboarding.displayPicture.url", {
        cell: info => (
          <HStack>
            {shouldWeDoSomething(info.row.original.onboarding?.displayPicture.verified || "") ? (
              <IconButton
                size="xs"
                aria-label="Go back"
                onClick={() => router.push(`onboardings/document/${info.row.original?.onboarding?.displayPicture.id}/${info.row.original?.onboarding?.displayPicture.__typename?.toLowerCase()}`)}
              >
                <ViewIcon />
              </IconButton>
            ) : (<DoneIcon />)}
            {renderDocImage(info.getValue())}
          </HStack>
        ),
        header: () => <span>Profile picture</span>
      }),
      columnHelper.accessor("onboarding.verification", {
        cell: info => renderStatusColumn(info.getValue()),
        header: () => <span>Verification</span>
      }),
    ]
  }, [])
  const table = useReactTable({
    data: payments,
    columns,
    getCoreRowModel: getCoreRowModel(),
  })

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
