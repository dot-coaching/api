import { eq } from "drizzle-orm";
import { MEDIA_QUERY_WITH } from "~/helper/query";
import { db } from "~/lib";
import { InsertProgramSchema, programs } from "~/schemas/clubs/programs";
import { APIResponse } from "~/types";
import { ServerType } from "..";

export function createProgramRouter(app: ServerType) {
  app.get(
    "/",
    async ({ query: { cursor = "0", limit = "10", clubId }, error }) => {
      if (!clubId) {
        return error(400, {
          error: "Club id is required",
        } satisfies APIResponse);
      }

      const find = await db.query.clubs.findFirst({
        where(fields, { eq }) {
          return eq(fields.id, parseInt(clubId));
        },
      });

      if (!find) {
        return error(404, {
          error: `Club with id ${clubId} not found`,
        } satisfies APIResponse);
      }

      const res = await db.query.programs.findMany({
        where(fields, { eq, gt, and }) {
          return and(
            gt(fields.id, parseInt(cursor || "0")),
            eq(fields.clubId, parseInt(clubId || "0"))
          );
        },
        with: {
          image: {
            columns: MEDIA_QUERY_WITH,
          },
        },
        limit: parseInt(limit),
      });

      if (res.length == 0) {
        return error(404, {
          error: "No programs found",
        } satisfies APIResponse);
      }

      for (const program of res as any) {
        program.media = program.image;
        delete program.image;
      }

      return {
        message: "Programs found",
        data: res,
      } satisfies APIResponse;
    }
  );
  app.get("/:id", async ({ params: { id }, error }) => {
    const res = await db
      .select()
      .from(programs)
      .where(eq(programs.id, parseInt(id)))
      .limit(1);

    if (res.length == 0) {
      return error(404, {
        error: `Program with id ${id} not found`,
      } satisfies APIResponse);
    }

    for (const program of res as any) {
      program.media = program.image;
      delete program.image;
    }

    return {
      message: `Program with id ${id} found`,
      data: res[0],
    } satisfies APIResponse;
  });
  app.post(
    "/",
    async ({ body, error }) => {
      const { clubId } = body;
      if (!clubId) {
        return error(400, {
          error: "Club id is required",
        } satisfies APIResponse);
      }

      const find = await db.query.clubs.findFirst({
        where(fields, { eq }) {
          return eq(fields.id, clubId);
        },
      });

      if (!find) {
        return error(404, {
          error: `Club with id ${clubId} not found`,
        } satisfies APIResponse);
      }

      const res = await db
        .insert(programs)
        .values({
          ...body,
          endDate: new Date(body.endDate || new Date()),
          startDate: new Date(body.startDate || new Date()),
        })
        .returning();
      if (res.length == 0) {
        return error(500, {
          error: `Failed to create ${body.name} program`,
        } satisfies APIResponse);
      }

      return {
        message: "Program inserted",
        data: res[0],
      } satisfies APIResponse;
    },
    {
      body: InsertProgramSchema,
    }
  );
  app.put(
    "/:id",
    async ({ params: { id }, body, error }) => {
      const { clubId } = body;
      if (!clubId) {
        return error(400, {
          error: "Club id is required",
        } satisfies APIResponse);
      }

      const find = await db.query.clubs.findFirst({
        where(fields, { eq }) {
          return eq(fields.id, clubId);
        },
      });

      if (!find) {
        return error(404, {
          error: `Club with id ${clubId} not found`,
        } satisfies APIResponse);
      }

      const res = await db
        .update(programs)
        .set({
          ...body,
          endDate: new Date(body.endDate || new Date()),
          startDate: new Date(body.startDate || new Date()),
        })
        .where(eq(programs.id, parseInt(id)))
        .returning();

      if (res.length == 0) {
        return error(500, {
          error: `Failed to update program with id ${id}`,
        } satisfies APIResponse);
      }

      return {
        message: `Program with id ${id} updated`,
        data: res[0],
      } satisfies APIResponse;
    },
    {
      body: InsertProgramSchema,
    }
  );
  app.delete("/:id", async ({ params: { id }, error }) => {
    const res = await db
      .delete(programs)
      .where(eq(programs.id, parseInt(id)))
      .returning();

    if (res.length == 0) {
      return error(500, {
        error: `Failed to delete program with id ${id}`,
      } satisfies APIResponse);
    }

    return {
      message: `Program with id ${id} deleted`,
      data: res[0],
    } satisfies APIResponse;
  });

  return app;
}
