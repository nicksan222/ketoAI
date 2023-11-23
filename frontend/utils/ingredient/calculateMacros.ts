import { Ingredient } from "@/types/ingredient";

/**
 *
 * @param ingredient The ingredient to calculate the macros for
 * @param quantity The quantity of the ingredient
 * @returns
 */
export function calculateMacrosForIngredient(
  ingredient: Ingredient,
  quantity: number
) {
  const fat = ingredient.fat * quantity;
  const protein = ingredient.protein * quantity;
  const carbs = ingredient.carbs * quantity;

  switch (ingredient.quantity_measurement) {
    case "L":
      return {
        fat: fat / 1000,
        protein: protein / 1000,
        carbs: carbs / 1000,
      };
    case "g":
      return {
        fat: fat / 100,
        protein: protein / 100,
        carbs: carbs / 100,
      };
    case "pcs":
      return {
        fat,
        protein,
        carbs,
      };
    case "qb":
      return {
        fat: 0,
        protein: 0,
        carbs: 0,
      };
  }
}
