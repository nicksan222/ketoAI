import { ObjectId } from "mongodb";

export const INGREDIENT_COLLECTION = "ingredients";

export interface Ingredient {
  id?: ObjectId;
  name: string;
  quantity_measurement: string;
  calories: number;
  protein: number;
  fat: number;
  carbs: number;
  is_main_ingredient: boolean;
  category: string;
}
