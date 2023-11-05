import chalk from "chalk";
import { MongoClient, Db } from "mongodb";

let clientInstance: Db | null = null;

// Check if MONGO_CONNECTION_STRING is set
if (!process.env.MONGO_CONNECTION_STRING) {
  console.log(
    chalk.red(`MONGO_CONNECTION_STRING is not set. Please set it in .env file.`)
  );
  process.exit(1);
}

if (!process.env.MONGO_DB_NAME) {
  console.log(
    chalk.red(`MONGO_DB_NAME is not set. Please set it in .env file.`)
  );
  process.exit(1);
}

async function connect(): Promise<Db> {
  const MONGO_CONNECTION_STRING = process.env.MONGO_CONNECTION_STRING;
  if (!MONGO_CONNECTION_STRING) {
    throw new Error("MONGO_CONNECTION_STRING not set");
  }

  const client = new MongoClient(MONGO_CONNECTION_STRING, {});
  await client.connect();

  const db = client.db(process.env.MONGO_DB_NAME);

  return db;
}

export async function getDbClient(): Promise<Db> {
  if (clientInstance) {
    return clientInstance;
  }

  clientInstance = await connect();

  return clientInstance;
}
