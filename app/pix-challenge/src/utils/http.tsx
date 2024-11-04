import { STATUS_CODES } from "http"
import { HTTP_METHODS } from "next/dist/server/web/http"

export function DefaultUrl(endpoint: string)
{
    return "http://localhost:8080/" + endpoint
}
function DefaultHeaders()
{
    return {
        "Content-Type": "application/json",
    }
}

export async function DefaultGetRequest(method: string, url: string)
{
    const res = await fetch(
        url,
        {
            next: {
                revalidate: 10
            },
            method: method,
            headers: DefaultHeaders()
        })
    return (res)
}

export async function DefaultRequest(method: string, url: string, body: string)
{
    const res = await fetch(
        url,
        {
            method: method,
            body: body,
            headers: DefaultHeaders()
        })
    return (res)
}

export async function ResolveResponse(res: Response)
{
    const body = await res.text()
    return (body)
}