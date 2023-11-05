import { ObjectId } from "mongodb";

export interface Recipe {
  _id?: ObjectId;
  title: string;
  steps: string[];
  createdBy?: string;
  createdAt?: Date;
  updatedAt?: Date;
  approved?: boolean;
  tags: string[];
  ingredients: {
    id: string;
    quantity: number;
  }[];
}