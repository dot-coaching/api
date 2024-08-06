import Elysia from "elysia";
import { db } from "../../core/services/db";
import { TacticalRepoImpl } from "./tactical.repo";

export const Dependency = () => {
  return new Elysia().derive({ as: "scoped" }, () => {
    return {
      get tacticalRepo() {
        return new TacticalRepoImpl(db);
      },
    };
  });
};
