"use client";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { fetchIngredients } from "@/hooks/ingredients/list";
import { Ingredient } from "@/types/ingredient";
import useNewRecipeStore from "./state";
import { useState } from "react";
import { Toast, ToastProvider } from "@/components/ui/toast";
import { useToast } from "@/components/ui/use-toast";
import { Toaster } from "@/components/ui/toaster";

interface Props {
  ingredients: Ingredient[];
}

export default function AddIngredientSelector({ ingredients }: Props) {
  const [ingredient, setIngredient] = useState<Ingredient | undefined>(
    undefined
  );
  const [quantity, setQuantity] = useState<number>(0);

  const { addIngredient } = useNewRecipeStore();
  const { toast } = useToast();

  return (
    <>
      <Toaster />

      <div className="grid md:grid-cols-3 gap-x-2">
        <Select
          onValueChange={(value) => {
            const ingredient = ingredients.find(
              (ingredient) => ingredient.id === value
            );
            setIngredient(ingredient);
          }}
        >
          <SelectTrigger className="w-full">
            <SelectValue placeholder="Seleziona un alimento" />
          </SelectTrigger>
          <SelectContent>
            <SelectGroup className="max-h-52">
              <SelectLabel>Ingredienti</SelectLabel>
              {ingredients.map((ingredient) => {
                return (
                  <SelectItem key={ingredient.id} value={ingredient.id}>
                    <div className="w-full flex flex-row justify-between">
                      <div className="text-left">{ingredient.name}</div>
                      <div className="text-right" style={{ color: "#9CA3AF" }}>
                        &nbsp;{ingredient.quantity_measurement}
                      </div>
                    </div>
                  </SelectItem>
                );
              })}
            </SelectGroup>
          </SelectContent>
        </Select>

        <Input
          placeholder="0"
          type="number"
          value={quantity}
          onChange={(e) => setQuantity(parseInt(e.target.value))}
        />

        <Button
          onClick={() => {
            if (ingredient) {
              addIngredient(ingredient, quantity);

              setIngredient(undefined);
              setQuantity(0);
            } else {
              toast({
                title: "Attenzione",
                description:
                  "Seleziona un ingrediente e inserisci una quantità",
                duration: 3000,
              });
            }
          }}
        >
          Aggiungi
        </Button>
      </div>
    </>
  );
}
