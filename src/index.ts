import cors from "@elysiajs/cors";
import jwt from "@elysiajs/jwt";
import { swagger } from "@elysiajs/swagger";
import { config } from "dotenv";
import { Elysia } from "elysia";
import logixlysia from "logixlysia";
import {
  AuthenticationError,
  AuthorizationError,
  BadRequestError,
  ERROR_CODE_STATUS_MAP,
} from "~/errors";
import { description, version } from "../package.json";
import { createRouter } from "./router";

config({ path: ".env" });

function setupApp<T>() {
  const app = new Elysia()
    .error({
      AUTHENTICATION: AuthenticationError,
      AUTHORIZATION: AuthorizationError,
      BAD_REQUEST: BadRequestError,
    })
    .onError(({ error, code, set }) => {
      set.status = ERROR_CODE_STATUS_MAP.get(code);
      const errorType = "type" in error ? error.type : "internal";
      return { errors: { [errorType]: error.message } };
    })
    .use(
      logixlysia({
        config: {
          ip: true,
          logFilePath: "logs/app.log",
          customLogFormat:
            "🦊 {now} {level} {duration} {method} {pathname} {status} {message} {ip}",
          // logFilter: {
          //   level: ["ERROR", "WARNING"],
          //   status: [500, 404],
          //   method: "GET",
          // },
        },
      })
    )
    .use(cors())
    .use(
      swagger({
        documentation: {
          info: { title: "DOT Coaching API", version, description },
        },
      })
    )
    .use(
      jwt({
        name: "jwt",
        secret: process.env.JWT_SECRET ?? "secret",
      })
    );
  return app;
}

export type ServerType = ReturnType<typeof setupApp>;

function main() {
  const app = setupApp();
  createRouter(app);

  app.listen(3000, () => {
    console.log(`Server started successfully 🚀`);
  });
}

main();