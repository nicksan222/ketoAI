import { Input } from "@/components/ui/input";
import AddIngredientSelector from "./components/addIngredient";
import { fetchIngredients } from "@/hooks/ingredients/list";
import IngredientList from "./components/ingredientsList";
import IngredientsMacros from "./components/ingredientsMacros";
import AddStep from "./components/addStep";
import StepsList from "./components/stepsList";
import InputName from "./components/inputName";

export default async function Page() {
  const ingredients = await fetchIngredients();

  return (
    <div>
      <InputName />

      <div className="mt-4 grid md:grid-cols-5 gap-4">
        <div className="col-span-3">
          <IngredientList />
          <div className="mt-4">
            <AddIngredientSelector ingredients={ingredients} />
          </div>
          <div className="mt-4">
            <StepsList />
          </div>
          <div className="my-4">
            <AddStep />
          </div>
        </div>

        <div className="col-span-2">
          <IngredientsMacros />
        </div>
      </div>

      <div></div>
    </div>
  );
}
