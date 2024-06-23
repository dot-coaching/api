import { getTableName } from "drizzle-orm";
import { exit } from "process";
import { db } from "~/db";
import { exams } from "~/schemas/exam";
import { questions } from "~/schemas/exam/question";
import { programs } from "~/schemas/programs";
import { tacticals } from "~/schemas/tacticals";
import { users } from "~/schemas/users";

const tables = [questions, exams, tacticals, programs, users] as const;

console.log("Dropping the entire database");

for (const table of tables) {
  const name = getTableName(table);
  console.log(`Dropping ${name}`);
  await db.delete(table);
  console.log(`Dropped ${name}`);
  const tableResult = await db.select().from(table);
  console.log(`${name} result: `, tableResult);
}

console.log("Database dropped");

exit(0);