"use client";

import { Avatar, AvatarImage } from "@/components/ui/avatar";
import { CardDescription } from "@/components/ui/card";

interface StepRowProps {
  step: string;
  index: number;
}

export default function StepRow({ index, step }: StepRowProps) {
  return (
    <div className="flex items-center mb-0 text-sm text-gray-800">
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
  );
}
