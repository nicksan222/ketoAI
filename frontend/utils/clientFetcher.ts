"use client";

import { UserProfile } from "@auth0/nextjs-auth0/client";

interface FetcherProps {
  url: string;
  method?: RequestInit["method"];
  body?: string;
  user?: UserProfile;
}

export const clientFetcher = async ({
  url,
  method,
  body,
  user,
}: FetcherProps): Promise<any | Error> => {
  const res = await fetch(process.env.NEXT_PUBLIC_BACKEND_URL + url, {
    method,
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${user?.sub}`,
    },
    body,
  });

  if (res.status > 400) {
    // Redirect to the login page.
    return Error("Something went wrong");
  }

  const response = await res.json();

  return response;
};
