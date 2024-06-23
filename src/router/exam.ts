import { ServerType } from "..";

export function createExamRouter(app: ServerType) {
  app.get("/", ({ query: { cursor, limit } }) => {});
  app.get("/:id", ({ params: { id } }) => {});
  app.post("/", () => {});
  app.put("/:id", ({ params: { id } }) => {});
  app.delete("/:id", ({ params: { id } }) => {});
  return app;
}
