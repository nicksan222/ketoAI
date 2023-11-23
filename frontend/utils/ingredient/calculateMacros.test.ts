import { Ingredient } from "@/types/ingredient";
import { calculateMacrosForIngredient } from "./calculateMacros";

describe("Calcolo dei macronutrienti", () => {
  test("Calcolo dei macronutrienti in grammi", () => {
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

    const result = calculateMacrosForIngredient(ingredient, quantity);

    expect(result).toEqual({
      fat: 1,
      protein: 2,
      carbs: 3,
    });
  });

  test("Calcolo dei macronutrienti in litri", () => {
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

    const result = calculateMacrosForIngredient(ingredient, quantity);

    expect(result).toEqual({
      fat: 1,
      protein: 2,
      carbs: 3,
    });
  });

  test("Calcolo dei macronutrienti in pezzi", () => {
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

    const result = calculateMacrosForIngredient(ingredient, quantity);

    expect(result).toEqual({
      fat: 1,
      protein: 2,
      carbs: 3,
    });
  });
});
