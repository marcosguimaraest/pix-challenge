import { json } from "stream/consumers"

export default function Generate()
{
    


    return (
        <main className="flex flex-col gap-8 row-start-2 items-center sm:items-start">
            <form action={async (formData: FormData) => {
                "use server"
                let data = formData.get("value")
                let dataString = data?.toString()
                let float
                if (dataString != undefined)
                {  
                    float = parseFloat(dataString)
                    console.log(float)
                }
                let stringForm = JSON.stringify({
                    customer: formData.get("customer"),
                    billingType: formData.get("billingType"),
                    value: float,
                    dueDate: formData.get("dueDate"),
                })
                console.log(JSON.parse(stringForm))
                const res = await fetch(
                    "http://localhost:8080/payment",
                    {
                        method: "POST",
                        body: stringForm,
                        headers: {
                            "Content-Type": "application/json",
                        }
                    })
                let body = await res.text()
                console.log(JSON.parse(body))
            }}>
                <input
                className="rounded-full border border-solid border-black/[.08] dark:border-white/[.145] transition-colors flex items-center justify-center hover:bg-[#f2f2f2] dark:hover:bg-[#1a1a1a] hover:border-transparent text-sm sm:text-base h-10 sm:h-12 px-4 sm:px-5 sm:min-w-44"
                name="customer"
                type="name"
                required
                />
                <input
                name="billingType"
                type="name"
                required
                />
                <input
                name="value"
                type="number"
                step="any"
                required
                />
                <input
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