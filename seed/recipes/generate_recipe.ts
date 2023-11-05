import * as fs from "fs";
import path from "path";
import { getDbClient } from "../db";
import openai from "../AI";
import { Ingredient } from "../ingredients/model";
import chalk from "chalk";

const db = await getDbClient();

// Get ingredients from db
const _mainIngredients = await db
  .collection<Ingredient>("ingredients")
  .find({
    is_main_ingredient: true,
  })
  .toArray();

const mainIngredients =
  _mainIngredients[Math.floor(Math.random() * _mainIngredients.length)];

const _sideIngredients = await db
  .collection<Ingredient>("ingredients")
  .find({
    is_main_ingredient: false,
  })
  .toArray();

const sideIngredients = _sideIngredients
  .sort(() => Math.random() - Math.random())
  .slice(0, Math.floor(Math.random() * 3));

// Generating a random recipe using OpenAI's GPT-4
const response = await openai.chat.completions.create({
  model: "gpt-4",
  messages: [
    {
      content: `
        Ti darò un json di vari alimenti
        Genera un json in questo formato

        {
        title: string;
        steps: string[];
        approved: boolean;
        tags: string[];
        ingredients: {
            name: string;
            quantity: number;
        }[];
        }
        Tutto sarà in lingua italiana
        Usa un titolo accattivante
        `,
      role: "system",
    },
    {
      content: JSON.stringify([mainIngredients, ...sideIngredients]),
      role: "user",
    },
  ],
});

console.log(response.choices[0].message.content);

// Append to recipes.json
const recipesPath = path.join(import.meta.dir, "/recipes.json");

if (!fs.existsSync(recipesPath)) {
  console.error(chalk.red("File recipes.json not found"));
  process.exit(1);
}

const recipes = {
  recipes: JSON.parse(fs.readFileSync(recipesPath, "utf8"))["recipes"],
};

recipes.recipes.push(response.choices[0].message.content);

// Save to file
fs.writeFileSync(recipesPath, JSON.stringify(recipes, null, 2), "utf8");



process.exit(0);
