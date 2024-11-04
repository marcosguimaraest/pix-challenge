"use server"

import GetCustomers from "@/utils/actions/customers"
import GetPayments from "@/utils/actions/payments"
import DialogCardCobranca from "./card_cobranca"

var customersList = await GetCustomers()
var paymentsList = await GetPayments()

export async function Cobrancas() {
    return (paymentsList.payments.map((payment) => {
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
    }))
}