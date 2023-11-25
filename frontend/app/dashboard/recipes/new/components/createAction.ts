"use server";

import { Ingredient } from "@/types/ingredient";
import { NewRecipeInterface } from "@/types/newRecipe";
import { fetcher } from "@/utils/fetcher";
import { revalidatePath } from "next/cache";

export async function createRecipe({
  title,
  ingredients,
  steps,
  tags,
}: NewRecipeInterface) {
  "use server";

  const ingredientsToSend: {
    ingredient: Ingredient;
    quantity: number;
    unit: string;
  }[] = ingredients.map((ingredient) => ({
    ingredient: ingredient.ingredient,
    quantity: ingredient.quantity,
    unit: ingredient.unit,
  }));

  fetcher({
    url: "/recipes",
    method: "POST",
    body: JSON.stringify({
      title,
      ingredients: ingredientsToSend,
      steps,
      tags,
    }),
  }).then((res) => {
    if (res.ok) {
      revalidatePath("/dashboard/recipes/waiting");
    }
  });
}
