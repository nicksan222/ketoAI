"use client";

import {
  Card,
  CardHeader,
  CardTitle,
  CardDescription,
  CardContent,
} from "@/components/ui/card";
import { Avatar, AvatarImage, AvatarFallback } from "@radix-ui/react-avatar";
import useNewRecipeStore from "./state";
import { Badge } from "@/components/ui/badge";

export default function StepsList() {
  const { steps, removeStep } = useNewRecipeStore();

  return (
    <Card className="col-span-3">
      <CardHeader>
        <CardTitle>Steps</CardTitle>
        <CardDescription>
          Gli steps per la preparazione della ricetta
        </CardDescription>
      </CardHeader>
      <CardContent>
        {steps.map((step, index) => (
          <div key={index} className="flex items-center mb-4">
            <div className="flex items-center mb-0">
              <Avatar className="h-9 w-9">
                <AvatarImage
                  src={"https://ui-avatars.com/api/?name=" + index.toString()}
                  alt="Avatar"
                />
              </Avatar>
              <div className="ml-4 space-y-1">
                <p className="text-sm font-medium leading-none">{step}</p>
              </div>
            </div>
            <div className="ml-auto font-medium ">
              <Badge
                className="rounded-full
              hover:cursor-pointer
          bg-red-500 text-white hover:bg-red-600"
                onClick={() => removeStep(index)}
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
