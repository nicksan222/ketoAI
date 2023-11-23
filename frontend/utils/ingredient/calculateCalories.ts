import { Ingredient } from "@/types/ingredient";

/**
 *
 * @param ingredient The ingredient to calculate the calories for
 * @param quantity The quantity of the ingredient
 * @returns
 */
export function calculateCaloriesForIngredient(
  ingredient: Ingredient,
  quantity: number
) {
  const calories =
    ingredient.fat * 9 + ingredient.protein * 4 + ingredient.carbs * 4;

  switch (ingredient.quantity_measurement) {
    case "L":
      return calories * (quantity / 1000);
    case "g":
      return calories * (quantity / 100);
    case "pcs":
      return calories * quantity;
    case "qb":
      return 0;
  }
}
