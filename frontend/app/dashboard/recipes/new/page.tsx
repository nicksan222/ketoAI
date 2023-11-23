import { Input } from "@/components/ui/input";
import AddIngredientSelector from "./components/addIngredient";
import { fetchIngredients } from "@/hooks/ingredients/list";
import IngredientList from "./components/ingredientsList";
import IngredientsMacros from "./components/ingredientsMacros";

export default async function Page() {
  const ingredients = await fetchIngredients();

  return (
    <div>
      <Input
        placeholder="Nome (Verrà generato quanto saranno inseriti degli step)"
        type="text"
        className="mt-4"
        disabled
      />

      <div className="mt-4 grid grid-cols-3 gap-4">
        <div className="col-span-2">
          <IngredientList />
        </div>

        <div className="col-span-1">
          <IngredientsMacros />
        </div>
      </div>

      <div className="mt-4">
        <AddIngredientSelector ingredients={ingredients} />
      </div>
    </div>
  );
}
