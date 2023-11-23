"use client";

import {
  Card,
  CardHeader,
  CardTitle,
  CardDescription,
  CardContent,
} from "@/components/ui/card";
import useNewRecipeStore from "./state";

import {
  Bar,
  BarChart,
  Pie,
  PieChart,
  ResponsiveContainer,
  XAxis,
  YAxis,
} from "recharts";
import { useEffect } from "react";

export default function IngredientsMacros() {
  const { totalCarbs, totalProtein, totalFat, totalCalories } =
    useNewRecipeStore();

  useEffect(() => {
    console.log(totalCarbs, totalProtein, totalFat, totalCalories);
  }, [totalCarbs, totalProtein, totalFat, totalCalories]);

  return (
    <Card className="col-span-3">
      <CardHeader>
        <CardTitle>Informazioni</CardTitle>
        <CardDescription>
          Informazioni riguardanti gli ingredienti della ricetta
        </CardDescription>
      </CardHeader>
      <CardContent>
        <ResponsiveContainer width={400} height={300}>
          <BarChart
            data={[
              {
                name: "Carboidrati",
                total: totalCarbs,
              },
              {
                name: "Proteine",
                total: totalProtein,
              },
              {
                name: "Grassi",
                total: totalFat,
              },
            ]}
          >
            <XAxis
              dataKey="name"
              stroke="#888888"
              fontSize={12}
              tickLine={false}
              axisLine={false}
            />
            <YAxis
              stroke="#888888"
              fontSize={12}
              tickLine={false}
              axisLine={false}
              tickFormatter={(value) => `${value}`}
            />
            <Bar dataKey="total" fill="#adfa1d" radius={[4, 4, 0, 0]} />
          </BarChart>
        </ResponsiveContainer>
      </CardContent>
    </Card>
  );
}
