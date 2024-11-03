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

function CardCobranca() {
  return (<Card className="relative hover:bg-gray-200 ">
    <CardHeader >
      <CardTitle>
        <div className="flex flex-row flex-nowrap items-center justify-between">
          <h1 className="text-2xl text-bold">pay_lyfihp3qruu851v7</h1>
          <Circle className="absolute bottom-3 right-2 text-red-600"></Circle>
        </div>
      </CardTitle>
    </CardHeader>
    <CardContent>
      <p className="text-l font-semibold">R$1349,90</p>
    </CardContent>
  </Card>)
}

function DialogCardCobranca() {
  return (<Dialog>
    <DialogTrigger>
      <CardCobranca></CardCobranca>
    </DialogTrigger>
    <DialogContent>
      <DialogHeader>
        <DialogTitle>pay_lyfihp3qruu851v7</DialogTitle>
      </DialogHeader>
      <div>
        <h1>R$1349,90</h1>
      </div>
      <Button>
        QRCode
      </Button>
    </DialogContent>
  </Dialog>)
}
export default function Home() {
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
              <Cobranca>
              </Cobranca>
            </div>
          </div>
        </div>
        <div className="grid grid-cols-4 gap-4">
          <DialogCardCobranca />
        </div>
      </div>
    </main>
  );
}
