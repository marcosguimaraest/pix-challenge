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
import Cobranca from "@/components/ui/dialog_cobranca";
import Image from "next/image";
import { Button } from "@/components/ui/button";
import GetCustomers from "@/utils/actions/customers";
import GetPayments, { Payment } from "@/utils/actions/payments";
import DialogCardCobranca from "@/components/ui/card_cobranca";

var customersList = await GetCustomers()
var paymentsList = await GetPayments()

export default async function Home() {
  return (
    <main>
      <div className="container mx-auto h-full">
        <div className="flex flex-wrap flex-col xl:flex-row items-center justify-center xl:pt-8 xl:pb-24 xl:pl-20 xl:pr-20">
          <div className="flex flex-row xl:flex-row items-center justify-between gap-60">
            <div className="text-center">
              <h1 className="text-2xl font-bold">
                _SUAS TRANSAÇÕES
              </h1>
            </div>
            <div className="flex flex-row flex-wrap xl:flex-row items-center justify-center">
              <Cobranca customersList={customersList}>
              </Cobranca>
            </div>
          </div>
        </div>
        <div className="grid grid-cols-4 gap-4 px-20 pb-16">
          {paymentsList.payments.map((payment) => {
            let customerName: string = ""
            customersList.customers.forEach((customer) => {
              if (customer.id == payment.customer) {
                customerName = customer.name
              }
            })
            const dialogProps = {
              payment: payment,
              customer: customerName
            }
            return <DialogCardCobranca key={payment.id} {...dialogProps} ></DialogCardCobranca>
          })}
        </div>
      </div>
    </main>
  );
}
