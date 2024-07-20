import { relations } from "drizzle-orm";
import { integer, pgTable, serial, text, timestamp } from "drizzle-orm/pg-core";
import { createInsertSchema, createSelectSchema } from "drizzle-typebox";

import { Static } from "elysia";
import { medias } from "~/schemas/media";
import { programs } from ".";

export const programExercises = pgTable("program_exercises", {
  id: serial("id").primaryKey(),
  programId: integer("program_id").references(() => programs.id, {
    onDelete: "cascade",
  }),
  mediaId: integer("media_id").references(() => medias.id, {
    onDelete: "set null",
  }),
  name: text("name").notNull(),
  description: text("description"),
  repetition: integer("repetition"),
  sets: integer("sets"),
  rest: integer("rest"),
  createdAt: timestamp("created_at").defaultNow(),
  updatedAt: timestamp("updated_at").defaultNow(),
});

export const programExercisesRelations = relations(
  programExercises,
  ({ one }) => ({
    program: one(programs, {
      fields: [programExercises.programId],
      references: [programs.id],
    }),
    media: one(medias, {
      fields: [programExercises.mediaId],
      references: [medias.id],
    }),
  })
);

export const InsertProgramExerciseSchema = createInsertSchema(programExercises);
export const SelectProgramExerciseSchema = createSelectSchema(programExercises);
export type ProgramType = Static<typeof SelectProgramExerciseSchema>;
