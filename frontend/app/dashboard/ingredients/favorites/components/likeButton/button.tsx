"use client";

import { Button } from "@/components/ui/button";
import { Ingredient } from "@/types/ingredient";
import { useTransition } from "react";
import { swapIngredientPreference } from "./actions";
import { Skeleton } from "@/components/ui/skeleton";

interface Props {
  ingredient: Ingredient;
  liked: boolean;
}

const LikeIngredientButton = ({ ingredient, liked }: Props) => {
  let [isPending, startTransition] = useTransition();

  if (isPending) {
    return <Skeleton className="flex items-center justify-center w-40  h-12" />;
  }

  if (liked) {
    return (
      <Button
        onClick={() =>
          startTransition(() => swapIngredientPreference({ ingredient, liked }))
        }
        className=" items-center justify-center bg-green-500 font-bold w-40 h-12"
      >
        Mi piace
      </Button>
    );
  }

  return (
    <Button
      onClick={() =>
        startTransition(() => swapIngredientPreference({ ingredient, liked }))
      }
      className=" items-center justify-center bg-red-500 font-bold w-40 h-12"
    >
      Non mi piace
    </Button>
  );
};

export default LikeIngredientButton;
