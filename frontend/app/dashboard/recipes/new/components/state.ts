import { Ingredient } from "@/types/ingredient";
import { NewRecipeInterface } from "@/types/newRecipe";
import { calculateCaloriesForIngredient } from "@/utils/ingredient/calculateCalories";
import { calculateMacrosForIngredient } from "@/utils/ingredient/calculateMacros";
import { create } from "zustand";

interface NewRecipeStore extends NewRecipeInterface {
  addIngredient: (ingredient: Ingredient, quantity: number) => void;
  addStep: (step: string) => void;
  setDescription: (description: string) => void;
  removeIngredient: (index: number) => void;
  removeStep: (index: number) => void;
  updateStep: (index: number, step: string) => void;
  setTitle: (title: string) => void;
  addTag: (tag: string) => void;
  removeTag: (index: number) => void;
  resetValues: () => void;

  totalCarbs: number;
  totalFat: number;
  totalProtein: number;
  totalCalories: number;
}

const useNewRecipeStore = create<NewRecipeStore>((set) => ({
  description: "",
  ingredients: [],
  steps: [],
  title: "",
  tags: [],

  totalCalories: 0,
  totalCarbs: 0,
  totalFat: 0,
  totalProtein: 0,

  addIngredient: (ingredient: Ingredient, quantity: number) => {
    set((state) => {
      const existingIngredientIndex = state.ingredients.findIndex(
        (i) => i.ingredient.id === ingredient.id
      );

      if (existingIngredientIndex !== -1) {
        // Update the quantity of the existing ingredient
        const updatedIngredients = [...state.ingredients];
        updatedIngredients[existingIngredientIndex] = {
          ...updatedIngredients[existingIngredientIndex],
          quantity:
            updatedIngredients[existingIngredientIndex].quantity + quantity,
        };

        // Recalculate macros and calories for the updated ingredient
        const newMacros = calculateMacrosForIngredient(
          ingredient,
          updatedIngredients[existingIngredientIndex].quantity
        );
        const newCalories = calculateCaloriesForIngredient(
          ingredient,
          updatedIngredients[existingIngredientIndex].quantity
        );

        return {
          ...state,
          ingredients: updatedIngredients,
          totalCalories:
            state.totalCalories +
            newCalories -
            calculateCaloriesForIngredient(ingredient, quantity),
          totalCarbs:
            state.totalCarbs +
            newMacros.carbs -
            calculateMacrosForIngredient(ingredient, quantity).carbs,
          totalFat:
            state.totalFat +
            newMacros.fat -
            calculateMacrosForIngredient(ingredient, quantity).fat,
          totalProtein:
            state.totalProtein +
            newMacros.protein -
            calculateMacrosForIngredient(ingredient, quantity).protein,
        };
      } else {
        // Add the new ingredient
        const macros = calculateMacrosForIngredient(ingredient, quantity);
        const calories = calculateCaloriesForIngredient(ingredient, quantity);

        return {
          ...state,
          ingredients: [
            ...state.ingredients,
            { ingredient, quantity, unit: ingredient.quantity_measurement },
          ],
          totalCalories: state.totalCalories + calories,
          totalCarbs: state.totalCarbs + macros.carbs,
          totalFat: state.totalFat + macros.fat,
          totalProtein: state.totalProtein + macros.protein,
        };
      }
    });
  },

  addStep: (step: string) =>
    set((state) => ({
      steps: [...state.steps, step],
    })),

  setDescription: (description: string) => set({ description }),

  removeIngredient: (index: number) => {
    set((state) => {
      const ingredient = state.ingredients[index];
      const macros = calculateMacrosForIngredient(
        ingredient.ingredient,
        ingredient.quantity
      );
      const calories = calculateCaloriesForIngredient(
        ingredient.ingredient,
        ingredient.quantity
      );

      return {
        ingredients: state.ingredients.filter((_, i) => i !== index),
        totalCalories: state.totalCalories - calories,
        totalCarbs: state.totalCarbs - macros.carbs,
        totalFat: state.totalFat - macros.fat,
        totalProtein: state.totalProtein - macros.protein,
      };
    });
  },

  removeStep: (index: number) =>
    set((state) => ({
      steps: state.steps.filter((_, i) => i !== index),
    })),

  updateStep: (index: number, step: string) =>
    set((state) => {
      const steps = [...state.steps];
      steps[index] = step;
      return { steps };
    }),

  setTitle: (title: string) => set({ title }),

  addTag: (tag: string) => set((state) => ({ tags: [...state.tags, tag] })),

  removeTag: (index: number) =>
    set((state) => ({ tags: state.tags.filter((_, i) => i !== index) })),

  resetValues: () =>
    set({
      description: "",
      ingredients: [],
      steps: [],
      title: "",
      tags: [],

      totalCalories: 0,
      totalCarbs: 0,
      totalFat: 0,
      totalProtein: 0,
    }),
}));

export default useNewRecipeStore;
