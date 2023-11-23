"use client";

import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import useNewRecipeStore from "./state";
import { useEffect } from "react";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";

export default function IngredientList() {
  const { ingredients, removeIngredient } = useNewRecipeStore();

  useEffect(() => {
    console.log(ingredients);
  }, [ingredients]);

  return (
    <Card className="col-span-3">
      <CardHeader>
        <CardTitle>Ingredienti</CardTitle>
        <CardDescription>Gli ingredienti della ricetta</CardDescription>
      </CardHeader>
      <CardContent>
        {ingredients.map((ingredient, index) => (
          <div
            className="flex items-center my-2"
            key={ingredient.ingredient.id}
          >
            <Avatar className="h-9 w-9">
              <AvatarImage
                src={
                  "https://ui-avatars.com/api/?name=" +
                  ingredient.ingredient.name +
                  "?length=" +
                  ingredient.ingredient.name.length
                }
                alt="Avatar"
              />
              <AvatarFallback>
                {ingredient.ingredient.name.slice(0, 2)}
              </AvatarFallback>
            </Avatar>
            <div className="ml-4 space-y-1">
              <p className="text-sm font-medium leading-none">
                {ingredient.ingredient.name}
              </p>
            </div>
            <div className="ml-auto font-medium ">
              <Badge className="bg-gray-500 text-white mr-2">
                {ingredient.quantity}{" "}
                {ingredient.ingredient.quantity_measurement}
              </Badge>
              <Badge
                className="rounded-full
                    hover:cursor-pointer
                bg-red-500 text-white hover:bg-red-600"
                onClick={() => removeIngredient(index)}
              >
                Rimuovi
              </Badge>
            </div>
          </div>
        ))}
      </CardContent>
    </Card>
  );
}
