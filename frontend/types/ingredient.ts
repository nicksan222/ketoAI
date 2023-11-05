export interface Ingredient {
  id?: string;
  name: string;
  quantityMeasurement: "g" | "L" | "pcs" | "qb";
  fat: number;
  protein: number;
  carbs: number;
  approved: boolean;
  is_main_ingredient: boolean;
}
