import { getSession } from "@auth0/nextjs-auth0";
import { redirect } from "next/navigation";

export default async function getSessionToken() {
  const session = await getSession();

  // Check if session exists
  if (!session) {
    // Redirect to login if there's no session
    redirect("/api/auth/login");
  }

  return session.idToken ?? "";
}
