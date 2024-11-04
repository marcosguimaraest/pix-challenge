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
            <p className="text-l font-semibold">{"R$" + payment.value + ",00"}</p>
        </CardContent>
    </Card>)
}

export default function DialogCardCobranca({payment, customer}: { payment: Payment, customer: string}) {
    return (<Dialog>
        <DialogTrigger>
            <CardCobranca key={payment.id} payment={payment} customer={customer} ></CardCobranca>
        </DialogTrigger>
        <DialogContent>
            <DialogHeader>
                <DialogTitle>{payment.customer}</DialogTitle>
            </DialogHeader>
            <div>
                <h1>{"R$" + payment.value + ",00"}</h1>
            </div>
            <Button>
                QRCode
            </Button>
        </DialogContent>
    </Dialog>)
}