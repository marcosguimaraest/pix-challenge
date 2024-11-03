import { DefaultRequest, DefaultUrl, ResolveResponse } from "../utils/http"
import { FormDataToPaymentObject } from "../utils/parse"

const paymentEndpoint = "payment"

export async function PaymentAction(formData: FormData) {
    "use server"
    let paymentObject = FormDataToPaymentObject(formData)
    //console.log(JSON.parse(paymentObject))
    const res = await DefaultRequest("POST", DefaultUrl(paymentEndpoint), paymentObject)
    ResolveResponse(res)
    //console.log(JSON.parse(body))
}