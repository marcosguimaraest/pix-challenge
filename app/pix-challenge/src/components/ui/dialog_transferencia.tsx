"use client"

import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from "@/components/ui/select"

import { zodResolver } from "@hookform/resolvers/zod"
import { UseFormReturn, useForm } from "react-hook-form"
import { date, z } from "zod"
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
} from "@/components/ui/dialog"
import {
    Form,
    FormControl,
    FormDescription,
    FormField,
    FormItem,
    FormLabel,
    FormMessage,
} from "@/components/ui/form"
import Image from "next/image";
import { Button } from "@/components/ui/button";
import { useState } from "react";
import { useToast } from "@/hooks/use-toast";
import { Loader2 } from 'lucide-react';
import GetCustomers, { CustomerList } from "@/utils/actions/customers"
import { DefaultRequest, DefaultUrl, ResolveResponse } from "@/utils/http"
import { Payment, PaymentRequest, PostPayment } from "@/utils/actions/payments"
import { useRouter } from "next/navigation"

const formSchema = z.object({
    customer: z.string().min(2).max(50),
    billingType: z.string().length(3),
    value: z.string().min(1).max(50),
    dueDate: z.string().min(2).max(50)
})

export default function Cobranca({ customersList }: { customersList: CustomerList }) {
    const router = useRouter()
    const { toast } = useToast()
    const form = useForm<z.infer<typeof formSchema>>({
        resolver: zodResolver(formSchema),
        defaultValues: {
            customer: "",
            billingType: "PIX",
            value: "",
            dueDate: ""
        },
    })
    function onSubmit(values: z.infer<typeof formSchema>) {
        var payment: PaymentRequest = {
            customer: values.customer,
            billingType: values.billingType,
            value: parseFloat(values.value),
            dueDate: values.dueDate
        }
        PostPayment(payment)
        form.reset()
        setDialogOpen(false)
        toast({
            variant: "default",
            title: "Cobrança criada!",
            description: "Esperando pagamento..."
        })
        router.refresh()
    }
    const [isDialogOpen, setDialogOpen] = useState(false)
    return (
        <div className="container mx-auto xl:p-14">
            <Dialog open={isDialogOpen} onOpenChange={setDialogOpen}>
                <DialogTrigger asChild>
                    <Button
                        className="flex gap-2"
                        >
                        Gerar Cobrança
                    </Button>
                </DialogTrigger>
                <DialogContent>
                    <DialogHeader>
                        <DialogTitle className="mb-10">Informe os dados da cobrança.</DialogTitle>
                        <DialogDescription>

                        </DialogDescription>
                    </DialogHeader>
                    <Form {...form}>
                        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
                            <FormField
                                control={form.control}
                                name="customer"
                                render={({ }) => (
                                    <FormItem>
                                        <FormLabel>Customer</FormLabel>
                                        <FormControl>
                                            <Select onValueChange={
                                                (value) => {
                                                    form.setValue("customer", value)
                                                }
                                            }>
                                                <SelectTrigger className="w-[300px]">
                                                    <SelectValue placeholder="Clientes" />
                                                </SelectTrigger>
                                                <SelectContent>
                                                    {customersList.customers.map((customer) => {
                                                        return (<SelectItem
                                                            key={customer.id} value={customer.id}> {customer.name}</SelectItem>)
                                                    })}
                                                </SelectContent>
                                            </Select>
                                        </FormControl>
                                        <FormMessage />
                                    </FormItem>
                                )}
                            />
                            <FormField
                                control={form.control}
                                name="value"
                                render={({ field }) => (
                                    <FormItem>
                                        <FormLabel>Valor</FormLabel>
                                        <FormControl>
                                            <Input type="number" step="0.01" min={"0.01"} placeholder="10.2" {...field} />
                                        </FormControl>
                                        <FormMessage />
                                    </FormItem>
                                )}
                            />
                            <FormField
                                control={form.control}
                                name="dueDate"
                                render={({ field }) => (
                                    <FormItem>
                                        <FormLabel>Valor</FormLabel>
                                        <FormControl>
                                            <Input type="date" {...field} />
                                        </FormControl>
                                        <FormMessage />
                                    </FormItem>
                                )}
                            />
                            <Button type="submit" disabled={form.formState.isSubmitting}>
                            {form.formState.isSubmitting && (<Loader2 className="mr-2 h-4 w-4 animate-spin"></Loader2>)}
                                Submit
                                </Button>
                        </form>
                    </Form>
                </DialogContent>
            </Dialog>
        </div>
    )
}