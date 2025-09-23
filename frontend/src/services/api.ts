// export async function apiRequest<T>(
//   endpoint: string,
//   options: RequestInit = {}
// ): Promise<T> {
//   const response = await fetch(endpoint, {
//     headers: {
//       "Content-Type": "application/json",
//       ...options.headers,
//     },
//     ...options,
//   });

//   if (!response.ok) {
//     const message = await response.text();
//     throw new Error(`API Error: ${message || response.statusText}`);
//   }

//   return response.json() as Promise<T>;
// }

export async function apiRequest<T>(
  endpoint: string,
  options: RequestInit = {}
): Promise<T> {
  const baseUrl = import.meta.env.VITE_APP_API_URL;
  const url = `${baseUrl}${endpoint}`;
    console.log("API Request URL:", url); // Debug log
    console.log("API URL:", import.meta.env.VITE_APP_API_URL);
  const res = await fetch(url, {
    headers: {
      "Content-Type": "application/json",
      ...(options.headers || {}),
    },
    ...options,
  });

  if (!res.ok) {
    throw new Error(`API error: ${res.status} ${res.statusText}`);
  }

  return res.json() as Promise<T>;
}