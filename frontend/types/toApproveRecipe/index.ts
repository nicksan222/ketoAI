import { Ingredient } from "../ingredient";

export interface ToApproveRecipeIngredient {
  ingredient: Ingredient;
  quantity: number;
  unit: string;
}

export interface ToApproveRecipeInterface {
  _id?: string;
  title: string;
  ingredients: ToApproveRecipeIngredient[];
  steps: string[];
  tags: string[];
  description?: string;
}
