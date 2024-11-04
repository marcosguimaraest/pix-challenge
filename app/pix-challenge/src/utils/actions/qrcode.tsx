import { DefaultGetRequest, DefaultUrl, ResolveResponse } from "../http"

export interface QrCode {
    encodedImage: string
    expirationDate: string
    payload: string
  }
  
export default async function GetQrCode(id: string)
{
    var qrCode: QrCode
    var res = await DefaultGetRequest("GET", DefaultUrl("qrcode?id=" + id))
    var body = await ResolveResponse(res)
    qrCode = JSON.parse(body)
    return (qrCode)
}