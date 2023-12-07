import { Ingredient } from "@/types/ingredient";
import { fetcher } from "@/utils/fetcher";
import getSessionToken from "@/utils/getSessionToken";
import api from "@api/backend";

export const fetchIngredients = async (): Promise<Ingredient[]> => {
  const sessionToken = await getSessionToken();
  api.auth(sessionToken);
  api.server(process.env.NEXT_PUBLIC_BACKEND_URL ?? "")

  console.info(api)

  try {
    // Use the fetcher with a specified response type
    const response = await api.getIngredients();

    // Directly return the ingredients array, ensuring it falls back to an empty array if undefined
    return []
  } catch (error) {
    // Handle or log the error
    console.error("Error fetching ingredients:", error);

    // Return an empty array in case of an error to maintain a consistent return type
    return [];
  }
};
