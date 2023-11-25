"use client"

import { Input } from "@/components/ui/input";
import useNewRecipeStore from "./state";

export default function InputName() {
  const { setTitle } = useNewRecipeStore();

  return (
    <Input
      placeholder="Nome (Verrà generato quanto saranno inseriti degli step)"
      type="text"
      className="mt-4"
      onChange={(e) => setTitle(e.target.value)}
    />
  );
}
