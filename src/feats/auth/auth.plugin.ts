import Elysia, { t } from "elysia";
import { AuthenticationError, BadRequestError } from "../../core/errors";
import { APIResponseSchema } from "../../core/response";
import { InsertUserSchema, SelectUserSchema } from "../user/user.schema";
import { Dependency } from "./auth.dependency";
import { AuthSignInSchema } from "./auth.schema";
import { AuthService } from "./auth.service";

export const AuthPlugin = new Elysia()
  .use(Dependency)
  .use(AuthService)
  .post(
    "/sign-in",
    async ({ rotateJWT, body, authRepo }) => {
      const identifier = body.identifier;
      if (identifier?.startsWith("0")) {
        body.identifier = `62${identifier}`;
      }
      if (identifier?.startsWith("8")) {
        body.identifier = `62${identifier}`;
      }
      if (identifier?.startsWith("+62")) {
        body.identifier = identifier.replace("+", "");
      }

      const user = await authRepo.signIn(body);
      const token = await rotateJWT(user);
      return {
        message: "Sign in success, welcome back!",
        data: {
          ...user,
          token,
        },
      };
    },
    {
      detail: {
        tags: ["AUTH"],
      },
      body: AuthSignInSchema,
      response: APIResponseSchema(
        t.Composite([
          t.Omit(SelectUserSchema, ["password", "fcmToken"]),
          t.Object({
            token: t.String(),
          }),
        ])
      ),
    }
  )
  .post(
    "/sign-up",
    async ({ body, authRepo }) => {
      const strPhone = body.phone.toString();

      if (strPhone.startsWith("8")) {
        body.phone = parseInt(`62${strPhone}`);
      } else {
        throw new BadRequestError("Invalid phone number");
      }

      const user = await authRepo.signUp(body);

      return {
        message: "Sign up success, welcome!",
        data: user,
      };
    },
    {
      detail: {
        tags: ["AUTH"],
      },
      body: InsertUserSchema,
      response: APIResponseSchema(
        t.Omit(SelectUserSchema, ["password", "fcmToken"])
      ),
    }
  )
  .get(
    "/me",
    async ({ verifyJWT, rotateJWT, userRepo }) => {
      const user = await verifyJWT();
      if (!user) throw new AuthenticationError("Unauthorized");
      const userInDB = await userRepo.find(user.id);
      const token = await rotateJWT(userInDB);
      return {
        data: {
          ...userInDB,
          token,
        },
      };
    },
    {
      detail: {
        tags: ["AUTH"],
      },
      response: APIResponseSchema(
        t.Composite([
          t.Omit(SelectUserSchema, ["password", "fcmToken"]),
          t.Object({
            token: t.String(),
          }),
        ])
      ),
    }
  )
  .get(
    "/logout",
    () => {
      return {
        message: "Logout success",
      };
    },
    {
      detail: {
        tags: ["AUTH"],
      },
      response: APIResponseSchema(t.Null()),
    }
  );