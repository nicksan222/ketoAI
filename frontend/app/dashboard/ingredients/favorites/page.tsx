import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { fetcher } from "@/utils/fetcher";
import { NextRequest } from "next/server";
import IngredientsTable from "./components/table";
import { Ingredient } from "@/types/ingredient";
import { fetchIngredients } from "@/hooks/ingredients/list";
import { fetchFavoriteIngredients } from "@/hooks/ingredients/getFavorites";

export default async function Page() {
  const [ingredients, favoriteIngredients] = await Promise.all([
    fetchIngredients(),
    fetchFavoriteIngredients(),
  ]);

  const mainIngredients = ingredients.filter((ingredient) => {
    return ingredient.is_main_ingredient;
  });

  return (
    <div className="grid gap-4 w-full py-10">
      <IngredientsTable
        ingredients={mainIngredients}
        favoriteIngredients={favoriteIngredients}
      />
    </div>
  );
}
