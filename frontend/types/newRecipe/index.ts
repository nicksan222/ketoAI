import { Ingredient } from "../ingredient";
import { newRecipeStep } from "./step";

export interface NewRecipeInterface {
  ingredients: {
    ingredient: Ingredient;
    quantity: number;
  }[];
  steps: newRecipeStep[];
  description?: string;
}
