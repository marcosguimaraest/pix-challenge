import Header from "@/app/components/ui/header";
import { Button } from "@/components/ui/button";

export default function Dev() {
    return (
        <section className="h-full">
            <div className="container mx-auto h-full">
                <div className="flex flex-col xl:flex-row items-center justify-between xl:pt-8 xl:pb-24 xl:pl-20 xl:pr-20">
                    <div className="text-center xl:text-left">
                        <h1 className="">
                            _ULTIMAS TRANSAÇÕES
                        </h1>
                    </div>
                    <div className="flex flex-row flex-wrap xl:flex-row items-center justify-center">
                        <Button>
                            GERAR COBRANÇA
                        </Button>
                        <Button>
                            REALIZAR SAQUE
                        </Button>
                    </div>
                </div>
            </div>
        </section>
    );
}