import { fetcher } from "@/utils/fetcher";

// Define a type for the response to ensure type safety.
type IngredientsResponse = {
  ingredients: string[];
};

export const fetchFavoriteIngredients = async (): Promise<string[]> => {
  try {
    // Use explicit type assertion for better type checking.
    const response = await fetcher<IngredientsResponse>({
      url: "/ingredients/preferences/list",
    });

    // Directly return the ingredients array, ensuring it falls back to an empty array if undefined.
    return response.ingredients ?? [];
  } catch (error) {
    // Handle errors gracefully.
    console.error("Error fetching favorite ingredients:", error);

    // Return an empty array in case of an error to maintain consistent function return type.
    return [];
  }
};
