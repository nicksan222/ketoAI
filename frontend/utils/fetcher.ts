import { getAccessToken } from "@auth0/nextjs-auth0";
import { getSession } from "@auth0/nextjs-auth0";
import { redirect } from "next/navigation";
import { Router } from "next/router";

export const fetcher = async (
  url: string,
  method: RequestInit["method"] = "GET"
): Promise<any> => {
  const session = await getSession();

  const res = await fetch(process.env.BACKEND_URL + url, {
    method,
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${session?.idToken}`,
    },
  });

  if (res.status === 401) {
    // Redirect to the login page.
    return redirect("/api/auth/login");
  }

  const response = await res.json();

  return response;
};
