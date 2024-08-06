import { TypeCompiler } from "@sinclair/typebox/compiler";
import { t } from "elysia";

const EnvSchema = t.Object({
  APP_PORT: t.String({ default: 3000 }),
  SALT_ROUND: t.String({ default: 10 }),
  JWT_SECRET: t.String({ minLength: 10 }),
  DATABASE_URL: t.String({ error: "DATABASE_URL didn't provided" }),
  SUPABASE_URL: t.String({ error: "SUPABASE_URL didn't provided" }),
  SUPABASE_ANON_KEY: t.String({ error: "SUPABASE_ANON_KEY didn't provided" }),
});

const createEnv = (payload: Record<string, string | undefined>) => {
  const EnvCompiler = TypeCompiler.Compile(EnvSchema);

  const err = EnvCompiler.Errors(payload);
  const errs = [...err];
  if (errs.length > 0) {
    errs.forEach((err) => {
      console.error({
        path: err.path,
        message: err.message,
        actualValue: err.value,
      });
    });
    return process.exit(1);
  }

  return EnvCompiler.Decode(payload);
};

const env = createEnv({
  APP_PORT: process.env.APP_PORT,
  SALT_ROUND: process.env.SALT_ROUND,
  JWT_SECRET: process.env.JWT_SECRET,
  DATABASE_URL: process.env.DATABASE_URL,
  SUPABASE_URL: process.env.SUPABASE_URL,
  SUPABASE_ANON_KEY: process.env.SUPABASE_ANON_KEY,
});

export { env };
