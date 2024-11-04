"use server"

import { get } from "http"
import { DefaultGetRequest, DefaultRequest, DefaultUrl, ResolveResponse } from "../http"

export interface CustomerList {
    customers: Customer[]
  }
  
  export interface Customer {
    id: string
    name: string
    cpfCnpj: string
  }
  
export default async function GetCustomers()
{
    var listOfCostumer: CustomerList
    var res = await DefaultGetRequest("GET", DefaultUrl("customer"))
    var body = await ResolveResponse(res)
    console.log(body)
    listOfCostumer = JSON.parse(body)
    console.log(listOfCostumer)
    return (listOfCostumer)
}