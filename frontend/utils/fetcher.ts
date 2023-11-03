import { getAccessToken } from "@auth0/nextjs-auth0";
import { getSession } from "@auth0/nextjs-auth0";

export const fetcher = async (
  url: string,
  method: RequestInit["method"] = "GET"
) => {
  const session = await getSession();

  const res = await fetch(process.env.BACKEND_URL + url, {
    method,
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${session?.idToken}`,
    },
  });

  return await res.json();
};
