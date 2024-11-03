import { json } from "stream/consumers"
import { FormDataToPaymentObject } from "../../utils/parse"
import { ResolveResponse } from "@/utils/http"
import { PaymentAction } from "@/actions/payment"

export default function Generate()
{
    return (
        <main className="flex flex-col gap-8 row-start-2 items-center sm:items-start">
            <form action={PaymentAction.bind(null)}>
                <input
                className="rounded-full border border-solid border-black/[.08] dark:border-white/[.145] transition-colors flex items-center justify-center hover:bg-[#f2f2f2] dark:hover:bg-[#1a1a1a] hover:border-transparent text-sm sm:text-base h-10 sm:h-12 px-4 sm:px-5 sm:min-w-44"
                name="customer"
                type="name"
                required
                />
                <input
                className="rounded-full border border-solid border-black/[.08] dark:border-white/[.145] transition-colors flex items-center justify-center hover:bg-[#f2f2f2] dark:hover:bg-[#1a1a1a] hover:border-transparent text-sm sm:text-base h-10 sm:h-12 px-4 sm:px-5 sm:min-w-44"
                name="billingType"
                type="name"
                required
                />
                <input
                className="rounded-full border border-solid border-black/[.08] dark:border-white/[.145] transition-colors flex items-center justify-center hover:bg-[#f2f2f2] dark:hover:bg-[#1a1a1a] hover:border-transparent text-sm sm:text-base h-10 sm:h-12 px-4 sm:px-5 sm:min-w-44"
                name="value"
                type="number"
                step="any"
                required
                />
                <input
                className="rounded-full border border-solid border-black/[.08] dark:border-white/[.145] transition-colors flex items-center justify-center hover:bg-[#f2f2f2] dark:hover:bg-[#1a1a1a] hover:border-transparent text-sm sm:text-base h-10 sm:h-12 px-4 sm:px-5 sm:min-w-44"
                name="dueDate"
                type="date"
                required
                />
                <button>
                    SUBMIT
                </button>
            </form>
        </main>
    )
}