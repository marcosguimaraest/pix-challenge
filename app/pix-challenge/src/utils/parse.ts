export function FormDataToPaymentObject(formData: FormData)
{
    let data = formData.get("value")
    let dataString = data?.toString()
    let float
    if (dataString != undefined)
    {  
        float = parseFloat(dataString)
        console.log(float)
    }
    let paymentObject = JSON.stringify({
        customer: formData.get("customer"),
        billingType: formData.get("billingType"),
        value: float,
        dueDate: formData.get("dueDate"),
    })
    return (paymentObject)
}

export function ResponseToPaymentResponseObject(body: string)
{
    return (body);
}