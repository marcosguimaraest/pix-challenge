"use server"
import {
    Card,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
} from "@/components/ui/dialog"

import { Circle } from 'lucide-react';
import { Button } from "@/components/ui/button";
import { Payment } from "@/utils/actions/payments";
import { Customer } from "@/utils/actions/customers";
import GetQrCode from "@/utils/actions/qrcode";
import Image from "next/image";

function CircleColor({status}: {status: string})
{
    if (status == "OVERDUE"){
        return(<Circle className="absolute bottom-3 right-2 text-red-600"></Circle>)
    }
    if (status == "PENDING"){
        return(<Circle className="absolute bottom-3 right-2 text-orange-400"></Circle>)
    }
    return(<Circle className="absolute bottom-3 right-2 text-green-600"></Circle>)
}

function CardCobranca({payment, customer}: { payment: Payment, customer: string}) {
    return (<Card className="relative hover:bg-gray-200 ">
        <CardHeader >
            <CardTitle>
                <div className="flex flex-row flex-nowrap items-center justify-between">
                    <h1 className="text-2xl text-bold">{customer}</h1>
                    <CircleColor status={payment.status}></CircleColor>
                </div>
            </CardTitle>
        </CardHeader>
        <CardContent>
            <p className="text-l font-semibold">{Intl.NumberFormat('pt-BR', {style: "currency", currency: "BSD"}).format(payment.value)}</p>
        </CardContent>
    </Card>)
}

export default async function DialogCardCobranca({payment, customer}: { payment: Payment, customer: string}) {
    var qrcode = await GetQrCode(payment.id)

    return (<Dialog>
        <DialogTrigger>
            <CardCobranca key={payment.id} payment={payment} customer={customer} ></CardCobranca>
        </DialogTrigger>
        <DialogContent>
            <DialogHeader>
                <DialogTitle>{customer}</DialogTitle>
            </DialogHeader>
            <div>
                <h1>{Intl.NumberFormat('pt-BR', {style: "currency", currency: "BSD"}).format(payment.value)}</h1>
            </div>
            <Image className="self-center" src={`data:image/jpeg;base64, ${qrcode.encodedImage}`} alt="QRCODE" width={200} height={200}/>
        </DialogContent>
    </Dialog>)
}