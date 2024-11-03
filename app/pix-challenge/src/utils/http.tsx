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
    if (res.status == 200)
    {
        const body = await res.text()
        const object = JSON.parse(body);
    }
}