import { getAccessToken } from "@auth0/nextjs-auth0";
import { getSession } from "@auth0/nextjs-auth0";
import { redirect } from "next/navigation";
import { Router } from "next/router";

interface FetcherProps {
  url: string;
  method?: RequestInit["method"];
  body?: string;
}

export const fetcher = async ({
  url,
  method,
  body,
}: FetcherProps): Promise<any> => {
  const session = await getSession();

  const res = await fetch(process.env.NEXT_PUBLIC_BACKEND_URL + url, {
    method,
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${session?.idToken}`,
    },
    body,
  });

  if (res.status === 401) {
    // Redirect to the login page.
    return redirect("/api/auth/login");
  }

  const response = await res.json();

  return response;
};
