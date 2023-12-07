import { Ingredient } from "@/types/ingredient";
import { fetcher } from "@/utils/fetcher";
import api from "@api/backend"

interface IngredientsResponse {
  ingredients: Ingredient[];
}

export const fetchIngredients = async (): Promise<Ingredient[]> => {
  try {
    // Use the fetcher with a specified response type
    const response = await fetcher<IngredientsResponse>({
      url: "/ingredients",
    });

    // Directly return the ingredients array, ensuring it falls back to an empty array if undefined
    return response.ingredients ?? [];
  } catch (error) {
    // Handle or log the error
    console.error("Error fetching ingredients:", error);

    // Return an empty array in case of an error to maintain a consistent return type
    return [];
  }
};
