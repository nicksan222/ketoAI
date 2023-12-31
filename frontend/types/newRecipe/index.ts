import { Ingredient } from "../ingredient";

export interface RecipeIngredient {
  ingredient: Ingredient;
  quantity: number;
  unit: string;
}

export interface NewRecipeInterface {
  _id?: string;
  title: string;
  ingredients: RecipeIngredient[];
  steps: string[];
  tags: string[];
  description?: string;
}