"use client";

import { Button } from "@/components/ui/button";
import { fetcher } from "@/utils/fetcher";
import { useTransition } from "react";
import useNewRecipeStore from "./state";
import { createRecipe } from "./createAction";
import { Toaster } from "@/components/ui/toaster";
import { useToast } from "@/components/ui/use-toast";

export default function CreateButton() {
  const [isPending, startTransition] = useTransition();
  const { ingredients, steps, title, tags, resetValues } = useNewRecipeStore();
  const { toast } = useToast();

  return (
    <>
      <Toaster />
      <Button
        disabled={isPending}
        className="w-56"
        onClick={() => {
          if (
            ingredients.length === 0 ||
            steps.length === 0 ||
            title.length === 0
          )
            return toast({
              title: "Errore",
              description: "Inserisci tutti i campi",
            });

          startTransition(() =>
            createRecipe({ title, ingredients, steps, tags })
              .then(() => {
                toast({
                  title: "Ricetta creata",
                  description: "La tua ricetta è stata creata",
                });
                resetValues();
              })
              .catch(() => {
                toast({
                  title: "Errore",
                  description: "Si è verificato un errore",
                });
              })
          );
        }}
      >
        Crea
      </Button>
    </>
  );
}
