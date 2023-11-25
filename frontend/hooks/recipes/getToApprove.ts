import { NewRecipeInterface } from "@/types/newRecipe";
import { ToApproveRecipeInterface } from "@/types/toApproveRecipe";
import { fetcher } from "@/utils/fetcher";

interface RecipesToApproveResponse {
  recipes: ToApproveRecipeInterface[];
}

export default async function getToApproveRecipes() {
  try {
    // Use the fetcher with a specified response type
    const response = await fetcher<RecipesToApproveResponse>({
      url: "/recipes",
    });

    // Directly return the ingredients array, ensuring it falls back to an empty array if undefined
    return response.recipes ?? [];
  } catch (error) {
    // Handle or log the error
    console.error("Error fetching ingredients:", error);

    // Return an empty array in case of an error to maintain a consistent return type
    return [];
  }
}
