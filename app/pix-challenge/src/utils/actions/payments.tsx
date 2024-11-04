import { get } from "http"
import { DefaultGetRequest, DefaultRequest, DefaultUrl, ResolveResponse } from "../http"

export interface PaymentsList {
  payments: Payment[]
}

export interface Payment {
  id: string
  customer: string
  billingType: string
  value: number
  dueDate: string
  status: string
}

  
export default async function GetPayments()
{
    var paymentsList: PaymentsList
    var res = await DefaultGetRequest("GET", DefaultUrl("payment"))
    var body = await ResolveResponse(res)
    paymentsList = JSON.parse(body)
    return (paymentsList)
}