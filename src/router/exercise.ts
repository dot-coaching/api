import { eq } from "drizzle-orm";
import { t } from "elysia";
import { MEDIA_QUERY_WITH } from "~/helper/query";
import { db } from "~/lib";
import {
  InsertProgramExerciseSchema,
  programExercises,
} from "~/schemas/clubs/programs/exercise";
import { APIResponse } from "~/types";
import { ServerType } from "..";

export function createExerciseRouter(app: ServerType) {
  app.get(
    "/",
    async ({
      query: { cursor = "0", limit = "10", programId = "0" },
      error,
    }) => {
      if (!programId) {
        return error(400, {
          error: "Program id is required",
        } satisfies APIResponse);
      }

      const find = await db.query.programs.findFirst({
        where(fields, { eq }) {
          return eq(fields.id, parseInt(programId));
        },
      });

      if (!find) {
        return error(404, {
          error: `Program with id ${programId} not found`,
        } satisfies APIResponse);
      }

      const res = await db.query.programExercises.findMany({
        where(fields, { and, gt, eq }) {
          return and(
            gt(fields.id, parseInt(cursor)),
            eq(fields.programId, parseInt(programId))
          );
        },
        limit: parseInt(limit),
        with: {
          media: {
            columns: MEDIA_QUERY_WITH,
          },
        },
      });
      if (res.length == 0) {
        return error(404, {
          error: "No Exercises found",
        } satisfies APIResponse);
      }
      return {
        message: "Exercises found",
        data: res,
      } satisfies APIResponse;
    }
  );
  app.get("/:id", async ({ params: { id }, error }) => {
    const res = await db
      .select()
      .from(programExercises)
      .where(eq(programExercises.id, parseInt(id)))
      .limit(1);

    if (res.length == 0) {
      return error(404, {
        error: `Exercise with id ${id} not found`,
      } satisfies APIResponse);
    }

    return {
      message: `Exercise with id ${id} found`,
      data: res[0],
    } satisfies APIResponse;
  });
  app.post(
    "/",
    async ({ body, error }) => {
      let programIds: number[] = [];
      for (const exercise of body) {
        const { programId } = exercise;
        if (!programId) {
          return error(400, {
            error: "Program id is required",
          } satisfies APIResponse);
        }

        if (!programIds.includes(programId)) {
          programIds.push(programId);
        }
      }

      const finds = await db.query.programs.findMany({
        where(fields, { inArray }) {
          return inArray(fields.id, programIds);
        },
      });

      for (const programId of programIds) {
        const find = finds.find((f) => f.id === programId);
        if (!find) {
          return error(404, {
            error: `Program with id ${programId} not found`,
          } satisfies APIResponse);
        }
      }

      const res = await db.insert(programExercises).values(body).returning();
      if (res.length == 0) {
        return error(500, {
          error: `Failed to create exercise`,
        } satisfies APIResponse);
      }

      return {
        message: "Exercise inserted",
        data: res,
      } satisfies APIResponse;
    },
    {
      body: t.Array(InsertProgramExerciseSchema),
    }
  );
  app.put(
    "/:id",
    async ({ params: { id }, body, error }) => {
      const { programId } = body;
      if (!programId) {
        return error(400, {
          error: "Club id is required",
        } satisfies APIResponse);
      }

      const find = await db.query.programs.findFirst({
        where(fields, { eq }) {
          return eq(fields.id, programId);
        },
      });

      if (!find) {
        return error(404, {
          error: `Exercise with id ${body.programId} not found`,
        } satisfies APIResponse);
      }

      const res = await db
        .update(programExercises)
        .set(body)
        .where(eq(programExercises.id, parseInt(id)))
        .returning();

      if (res.length == 0) {
        return error(500, {
          error: `Failed to update exercise with id ${id}`,
        } satisfies APIResponse);
      }

      return {
        message: `Exercise with id ${id} updated`,
        data: res[0],
      } satisfies APIResponse;
    },
    {
      body: InsertProgramExerciseSchema,
    }
  );
  app.delete("/:id", async ({ params: { id }, error }) => {
    const res = await db
      .delete(programExercises)
      .where(eq(programExercises.id, parseInt(id)))
      .returning();

    if (res.length == 0) {
      return error(500, {
        error: `Failed to delete exercise with id ${id}`,
      } satisfies APIResponse);
    }

    return {
      message: `Exercise with id ${id} deleted`,
      data: res[0],
    } satisfies APIResponse;
  });

  return app;
}