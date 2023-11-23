import { getSession } from "@auth0/nextjs-auth0";
import { redirect } from "next/navigation";

interface FetcherProps {
  url: string;
  method?: RequestInit["method"];
  body?: string;
}

// Use a generic type T for better type safety.
export const fetcher = async <T = any>({
  url,
  method = "GET", // Default method to GET if not provided
  body,
}: FetcherProps): Promise<T> => {
  const session = await getSession();

  // Check if session exists
  if (!session) {
    // Redirect to login if there's no session
    redirect("/api/auth/login");
    return Promise.reject("No session found");
  }

  const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}${url}`, {
    method,
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${session.idToken}`,
    },
    body,
  });

  if (!res.ok) {
    if (res.status === 401) {
      // Redirect to login for unauthorized requests
      redirect("/api/auth/login");
    }

    // Handle other HTTP errors
    const errorDetail = await res.text();
    return Promise.reject(
      new Error(`HTTP error ${res.status}: ${errorDetail}`)
    );
  }

  const response = await res.json();

  return response as T;
};
