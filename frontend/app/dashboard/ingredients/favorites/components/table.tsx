import { Button } from "@/components/ui/button";
import {
  TableCaption,
  TableHeader,
  TableRow,
  TableHead,
  TableBody,
  TableCell,
  Table,
} from "@/components/ui/table";
import { Ingredient } from "@/types/ingredient";

import { FiArrowUp } from "react-icons/fi";
import LikeIngredientButton from "./likeButton/button";

interface Props {
  ingredients: Ingredient[];
  favoriteIngredients: string[];
}

const IngredientsTable = ({ ingredients, favoriteIngredients }: Props) => {
  // Calculating medians for each macronutrient
  const carbsMedian =
    ingredients.reduce((acc, ingredient) => {
      return acc + ingredient.carbs;
    }, 0) / ingredients.length;

  const fatMedian =
    ingredients.reduce((acc, ingredient) => {
      return acc + ingredient.fat;
    }, 0) / ingredients.length;

  const proteinMedian =
    ingredients.reduce((acc, ingredient) => {
      return acc + ingredient.protein;
    }, 0) / ingredients.length;

  const isHighCarb = (ingredient: Ingredient) => {
    // Is this higher than the median? (10 g of difference)
    return ingredient.carbs > carbsMedian + 5;
  };

  const isHighFat = (ingredient: Ingredient) => {
    // Is this higher than the median? (10 g of difference)
    return ingredient.fat > fatMedian + 5;
  };

  const isHighProtein = (ingredient: Ingredient) => {
    // Is this higher than the median? (10 g of difference)
    return ingredient.protein > proteinMedian + 5;
  };

  return (
    <Table>
      <TableCaption>
        Seleziona le tue preferenze riguardo agli ingredienti.
      </TableCaption>
      <TableHeader>
        <TableRow>
          <TableHead>Ingrediente</TableHead>
          <TableHead>Preferiti</TableHead>
          <TableHead>Carboidrati</TableHead>
          <TableHead>Grassi</TableHead>
          <TableHead>Proteine</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {ingredients.map((ingredient) => (
          <TableRow key={ingredient.id}>
            <TableCell className="font-medium">{ingredient.name}</TableCell>
            <TableCell>
              <LikeIngredientButton
                ingredient={ingredient}
                liked={(favoriteIngredients ?? []).includes(
                  ingredient.id ?? ""
                )}
              />
            </TableCell>
            <TableCell>
              <div className="flex flex-row">
                {ingredient.carbs} g
                {isHighCarb(ingredient) && (
                  <FiArrowUp className="text-red-500 font-bold ml-4" />
                )}
              </div>
            </TableCell>
            <TableCell>
              <div className="flex flex-row">
                {ingredient.fat} g
                {isHighFat(ingredient) && (
                  <FiArrowUp className="text-green-500 font-bold ml-4" />
                )}
              </div>
            </TableCell>
            <TableCell>
              <div className="flex flex-row">
                {ingredient.protein} g
                {isHighProtein(ingredient) && (
                  <FiArrowUp className="text-green-500 font-bold ml-4" />
                )}
              </div>
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );
};

export default IngredientsTable;
