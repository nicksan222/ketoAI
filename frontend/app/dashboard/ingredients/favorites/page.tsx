import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { fetcher } from "@/utils/fetcher";
import { NextRequest } from "next/server";
import IngredientsTable from "./components/table";
import { Ingredient } from "@/types/ingredient";

export default async function Page() {
  const ingredients: Ingredient[] = (await fetcher("/api/v1/ingredients"))[
    "ingredients"
  ];

  return (
    <div className="grid gap-4 w-full py-10">
      <IngredientsTable ingredients={ingredients} />
    </div>
  );
}
