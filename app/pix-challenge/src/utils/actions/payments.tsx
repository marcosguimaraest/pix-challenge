"use server"

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

export interface PaymentRequest {
  customer: string
  billingType: string
  value: number
  dueDate: string
}

  
export default async function GetPayments()
{
    var paymentsList: PaymentsList
    var res = await DefaultGetRequest("GET", DefaultUrl("payment"))
    var body = await ResolveResponse(res)
    paymentsList = JSON.parse(body)
    return (paymentsList)
}

export async function PostPayment(payment: PaymentRequest)
{
  var stringfiedPayment = JSON.stringify(payment)
  console.log(stringfiedPayment)
  var res = await DefaultRequest("POST", DefaultUrl("payment"), stringfiedPayment)
  var body = await ResolveResponse(res);
}