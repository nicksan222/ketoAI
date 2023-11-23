import { Ingredient } from "@/types/ingredient";
import { NewRecipeInterface } from "@/types/newRecipe";
import { newRecipeStep } from "@/types/newRecipe/step";
import { create } from "zustand";

interface NewRecipeStore extends NewRecipeInterface {
  addIngredient: (ingredient: Ingredient, quantity: number) => void;
  addStep: (step: newRecipeStep) => void;
  setDescription: (description: string) => void;
  removeIngredient: (index: number) => void;
  removeStep: (index: number) => void;
  updateStep: (index: number, step: newRecipeStep) => void;

  totalCarbs: number;
  totalFat: number;
  totalProtein: number;
  totalCalories: number;
}

const useNewRecipeStore = create<NewRecipeStore>((set) => ({
  description: "",
  ingredients: [],
  steps: [],

  addIngredient: (ingredient: Ingredient, quantity: number) => {
    set((state) => ({
      ingredients: [
        ...state.ingredients,
        {
          ingredient,
          quantity,
        },
      ],
    }));
  },

  addStep: (step: newRecipeStep) =>
    set((state) => ({
      steps: [...state.steps, step],
    })),

  setDescription: (description: string) => set({ description }),

  removeIngredient: (index: number) =>
    set((state) => ({
      ingredients: state.ingredients.filter((_, i) => i !== index),
    })),

  removeStep: (index: number) =>
    set((state) => ({
      steps: state.steps.filter((_, i) => i !== index),
    })),

  updateStep: (index: number, step: newRecipeStep) =>
    set((state) => {
      const steps = [...state.steps];
      steps[index] = step;
      return { steps };
    }),

  totalCalories: 0,
  totalCarbs: 0,
  totalFat: 0,
  totalProtein: 0,
}));

export default useNewRecipeStore;
