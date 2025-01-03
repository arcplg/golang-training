export interface RequestOptions extends RequestInit {
  headers?: HeadersInit
}

export async function httpRequest<T>(
  url: string,
  options: RequestOptions = {},
): Promise<T> {
  const defaultHeaders = {
    "Content-Type": "application/json",
  }

  const mergedOptions: RequestOptions = {
    ...options,
    headers: {
      ...defaultHeaders,
      ...options.headers,
    },
  }

  const response = await $fetch(url, mergedOptions)

  if (!response.ok) {
    const error = await response.json()
    throw new Error(error.message || "An error occurred while fetching data")
  }

  return response.json()
}
