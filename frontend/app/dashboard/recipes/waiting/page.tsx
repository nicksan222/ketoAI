import { Avatar, AvatarImage } from "@/components/ui/avatar";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { fetchIngredients } from "@/hooks/ingredients/list";
import getToApproveRecipes from "@/hooks/recipes/getToApprove";
import StepRow from "./components/stepRow";

export default async function Page() {
  const [recipes] = await Promise.all([getToApproveRecipes()]);

  return (
    <div>
      {recipes.map((recipe) => (
        <Card className=" mt-3" key={recipe._id}>
          <CardHeader>
            <CardTitle>{recipe.title}</CardTitle>
            <CardDescription className="text-sm text-gray-500">
              In attesa di approvazione
            </CardDescription>
          </CardHeader>
          <CardContent>
            <div className="grid grid-cols-2 gap-4">
              <div>
                <p className="text-lg text-gray-500">Ingredienti</p>
                <div className="grid grid-cols-2 gap-2">
                  {recipe.ingredients.map((ingredient, index) => (
                    <Card key={index} className="mt-3">
                      <CardHeader>
                        <CardTitle className="text-md">
                          {ingredient.ingredient.name}
                        </CardTitle>
                        <CardDescription className="text-sm text-gray-500">
                          {ingredient.quantity} {ingredient.unit}
                        </CardDescription>
                      </CardHeader>
                    </Card>
                  ))}
                </div>
              </div>
              <div>
                <p className="text-lg text-gray-500">Preparazione</p>
                <div className="mt-3">
                  <Card>
                    <CardHeader>
                      {recipe.steps.map((step, index) => (
                        <StepRow key={index} index={index} step={step} />
                      ))}
                    </CardHeader>
                  </Card>
                </div>
              </div>
            </div>
          </CardContent>
        </Card>
      ))}
    </div>
  );
}
