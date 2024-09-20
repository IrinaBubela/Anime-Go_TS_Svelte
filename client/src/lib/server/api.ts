export default async function api<T, D = T>(
    url: string,
    {
        method = "GET",
        body = undefined,
    }: {
        method?: "GET" | "POST" | "PUT" | "DELETE" | "PATCH";
        body?: D;
    } = {},
): Promise<{ success: true; data: T } | { success: false; error: string }> {
    try {
        const res = await fetch(url, {
            method,
            headers: {
                "content-type": "application/json",
            },
            body: body ? JSON.stringify(body) : null,
        });

        // check if invalid response
        if (!res.ok) {
            return { success: false, error: await res.text() };
        }
        // check if empty response
        if (res.status === 204) {
            return { success: true, data: {} as T };
        }
        // check if invalid response
        if (!res.headers.get("content-type")?.includes("application/json")) {
            return { success: false, error: "Response was not JSON" };
        }
        const data = (await res.json()) as T;
        return { success: true, data };
    } catch (error) {
        console.error(error);
        return { success: false, error: "Unknown error" };
    }
}
