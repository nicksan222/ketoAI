import * as fs from "fs";
import * as path from "path";
import chalk from "chalk";
import { Ingredient } from "./model";
import { getDbClient } from "../db";

let ingredients: Ingredient[] = [];

// Does ingredients.json exist?
const ingredientsPath = path.join(import.meta.dir, "/ingredients.json");

if (!fs.existsSync(ingredientsPath)) {
  console.log(
    chalk.red(
      ingredientsPath +
        ` does not exist. Please run seed/ingredients/process_raw.ts first.`
    )
  );
  process.exit(1);
}

ingredients = JSON.parse(fs.readFileSync(ingredientsPath, "utf8"))[
  "ingredients"
];

if (!ingredients.length) {
  console.log(
    chalk.red(
      `ingredients.json is empty. Please run seed/ingredients/process_raw.ts first.`
    )
  );
  process.exit(1);
}

const db = await getDbClient();

let skipped = 0;

for (const ingredient of ingredients) {
  // Does this exact name exist?
  const existingingredient = await db.collection("ingredients").findOne({
    name: ingredient.name,
  });

  if (existingingredient) {
    skipped++;
    continue;
  }

  const result = await db
    .collection<Ingredient>("ingredients")
    .insertOne(ingredient);

  if (!result.acknowledged) {
    console.log(
      chalk.red(
        `Failed to insert ingredient ${ingredient.name} into ingredients.`
      )
    );
    process.exit(1);
  }
}

console.log(
  chalk.green(
    `Inserted ${
      ingredients.length - skipped
    } ingredients into ingredients collection.`
  )
);

process.exit(0);
