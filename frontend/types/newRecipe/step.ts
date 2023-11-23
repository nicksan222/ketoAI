import { Ingredient } from "../ingredient";

export type newRecipeStepElement = string | Ingredient;

export interface newRecipeStep {
  elements: newRecipeStepElement[];
}
