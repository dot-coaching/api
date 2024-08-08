import { and, eq } from "drizzle-orm";
import {
  BadRequestError,
  NoContentError,
  ServerError,
} from "../../core/errors";
import { BaseRepo } from "../../core/repo";
import { DrizzlePostgres } from "../../core/services/db";
import { MediaModel } from "./media.model";
import { InsertMedia, Media, MediaParent } from "./media.schema";

abstract class MediaRepo extends BaseRepo<Media> {
  abstract findByURL(url: string): Promise<Media>;
}

export class MediaRepoImpl extends MediaRepo {
  private db: DrizzlePostgres;

  constructor(db: DrizzlePostgres) {
    super();
    this.db = db;
  }

  async create(data: InsertMedia): Promise<Media> {
    const medias = await this.db.insert(MediaModel).values(data).returning();

    if (medias.length === 0) throw new ServerError("Failed to create media");

    return medias[0];
  }

  async update(data: InsertMedia): Promise<Media> {
    if (!data.id) throw new BadRequestError("Media id is required");

    const medias = await this.db
      .update(MediaModel)
      .set(data)
      .where(eq(MediaModel.id, data.id))
      .returning();

    if (medias.length === 0) throw new ServerError("Failed to update media");

    return medias[0];
  }

  async delete(id: number): Promise<Media> {
    const medias = await this.db
      .delete(MediaModel)
      .where(eq(MediaModel.id, id))
      .returning();

    if (medias.length === 0) throw new ServerError("Failed to delete media");

    return medias[0];
  }

  async find(id: number): Promise<Media> {
    const medias = await this.db
      .select()
      .from(MediaModel)
      .where(eq(MediaModel.id, id));

    if (medias.length === 0) throw new NoContentError("Media not found");

    return medias[0];
  }

  async list({
    clubId,
    parent,
  }: {
    clubId: number;
    parent: MediaParent;
  }): Promise<Media[]> {
    const medias = await this.db
      .select()
      .from(MediaModel)
      .where(and(eq(MediaModel.clubId, clubId), eq(MediaModel.parent, parent)));

    if (medias.length === 0) throw new NoContentError("No media found");

    return medias;
  }

  async findByURL(url: string): Promise<Media> {
    const medias = await this.db
      .select()
      .from(MediaModel)
      .where(eq(MediaModel.url, url));

    if (medias.length === 0) throw new NoContentError("Media not found");

    return medias[0];
  }
}
