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

interface Props {
  ingredients: Ingredient[];
}

const IngredientsTable = ({ ingredients }: Props) => {
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
    return ingredient.carbs > carbsMedian + 10;
  };

  const isHighFat = (ingredient: Ingredient) => {
    // Is this higher than the median? (10 g of difference)
    return ingredient.fat > fatMedian + 10;
  };

  const isHighProtein = (ingredient: Ingredient) => {
    // Is this higher than the median? (10 g of difference)
    return ingredient.protein > proteinMedian + 10;
  };

  return (
    <Table>
      <TableCaption>
        Seleziona le tue preferenze riguardo agli ingredienti.
      </TableCaption>
      <TableHeader>
        <TableRow>
          <TableHead>Ingrediente</TableHead>
          <TableHead>Carboidrati</TableHead>
          <TableHead>Grassi</TableHead>
          <TableHead>Proteine</TableHead>
          <TableHead className="text-right">Preferiti</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {ingredients.map((ingredient) => (
          <TableRow key={ingredient.id}>
            <TableCell className="font-medium">{ingredient.name}</TableCell>
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
            <TableCell>$250.00</TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );
};

export default IngredientsTable;
