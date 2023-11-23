import { Ingredient } from "@/types/ingredient";
import { calculateCaloriesForIngredient } from "./calculateCalories";

describe("Calcolo delle calorie", () => {
  it("Calcolo delle calorie in grammi", () => {
    const ingredient: Ingredient = {
      id: "1",
      name: "Pasta",
      quantity_measurement: "g",
      fat: 1,
      protein: 2,
      carbs: 3,
      approved: true,
      is_main_ingredient: true,
    };

    const quantity = 100;
    const result = calculateCaloriesForIngredient(ingredient, quantity);
    expect(result).toEqual(29);
  });

  it("Calcolo delle calorie in litri", () => {
    const ingredient: Ingredient = {
      id: "1",
      name: "Pasta",
      quantity_measurement: "L",
      fat: 1,
      protein: 2,
      carbs: 3,
      approved: true,
      is_main_ingredient: true,
    };

    const quantity = 1000;
    const result = calculateCaloriesForIngredient(ingredient, quantity);
    expect(result).toEqual(29);
  });

  it("Calcolo delle calorie in pezzi", () => {
    const ingredient: Ingredient = {
      id: "1",
      name: "Pasta",
      quantity_measurement: "pcs",
      fat: 1,
      protein: 2,
      carbs: 3,
      approved: true,
      is_main_ingredient: true,
    };

    const quantity = 1;
    const result = calculateCaloriesForIngredient(ingredient, quantity);
    expect(result).toEqual(29);
  });
});
