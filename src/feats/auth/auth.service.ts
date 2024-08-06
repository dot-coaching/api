import bearer from "@elysiajs/bearer";
import jwt from "@elysiajs/jwt";
import Elysia from "elysia";
import { AuthenticationError } from "../../core/errors";
import { env } from "../../utils/env";
import { User } from "../user/user.schema";
import { AuthJWT } from "./auth.schema";

export const AuthService = new Elysia()
  .use(
    jwt({
      name: "dotJWT",
      secret: env.JWT_SECRET,
      exp: "7d",
    })
  )
  .use(bearer())
  .derive({ as: "global" }, ({ dotJWT, bearer, cookie: { auth } }) => ({
    async verifyJWT() {
      const verify = await dotJWT.verify(bearer);
      if (!verify) throw new AuthenticationError("Invalid token");
      return verify as AuthJWT;
    },
    async rotateJWT(user: User) {
      const token = await dotJWT.sign({
        id: user.id,
        email: user.email,
        name: user.name,
        image: user.image,
        exp: Math.floor(Date.now() / 1000) + 7 * 86400, // 7 days
        iat: Math.floor(Date.now() / 1000),
      });
      auth.set({
        value: token,
        httpOnly: true,
        maxAge: 7 * 86400,
      });
      return token;
    },
  }));
