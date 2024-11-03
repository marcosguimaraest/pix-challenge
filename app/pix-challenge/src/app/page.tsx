import Cobranca from "@/components/ui/dialog_cobranca";
import Image from "next/image";

export default function Home() {
  return (
    <main>
      <div className="container mx-auto h-full">
                <div className="flex flex-col xl:flex-row items-center justify-between xl:pt-8 xl:pb-24 xl:pl-20 xl:pr-20">
                    <div className="text-center xl:text-left">
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
    </main>
  );
}
