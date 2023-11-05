"use server";

import { Ingredient } from "@/types/ingredient";
import { fetcher } from "@/utils/fetcher";
import { revalidatePath } from "next/cache";

interface SwapIngredientPreferenceProps {
  liked: boolean;
  ingredient: Ingredient;
}

export async function swapIngredientPreference({
  ingredient,
  liked,
}: SwapIngredientPreferenceProps) {
  if (liked) {
    // Remove this ingredient from the liked ones
    const response = await fetcher({
        url: `/ingredients/preferences/` + ingredient.id,
        method: "DELETE",
    });

    if (response.ok) {
      revalidatePath("/dashboard/ingredients/favorites");
    }
  } else {
    const response = await fetcher({
      url: "/ingredients/preferences",
      method: "POST",
      body: JSON.stringify({
        ingredient_ids: [ingredient.id],
      }),
    });

    if (response.ok) {
      revalidatePath("/dashboard/ingredients/favorites");
    }
  }
}
