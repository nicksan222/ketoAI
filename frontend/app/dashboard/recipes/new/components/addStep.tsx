"use client";

import { Input } from "@/components/ui/input";
import useNewRecipeStore from "./state";
import { Button } from "@/components/ui/button";
import { useState } from "react";

export default function AddStep() {
  const { addStep } = useNewRecipeStore();
  const [step, setStep] = useState<string>("");

  return (
    <div className="grid md:grid-cols-3 gap-x-2">
      <div className="md:col-span-2">
        <Input
          className="md:mt-0 mt-4"
          placeholder="Inserisci un nuovo step"
          type="text"
          maxLength={255}
          minLength={8}
          value={step}
          onChange={(e) => setStep(e.target.value)}
        />
      </div>

      <Button
        className="md:mt-0 mt-4 w-full"
        onClick={() => {
          if (step) {
            addStep(step);
            setStep("");
          }
        }}
      >
        Aggiungi
      </Button>
    </div>
  );
}
