DO $$ BEGIN
 CREATE TYPE "public"."question_type" AS ENUM('multipleChoice', 'trueFalse', 'shortAnswer', 'essay');
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 CREATE TYPE "public"."sport_type" AS ENUM('volleyBall', 'basketBall', 'soccer');
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 CREATE TYPE "public"."media_parent" AS ENUM('club', 'program', 'exercise', 'exam', 'question', 'user');
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 CREATE TYPE "public"."media_type" AS ENUM('image/png', 'image/jpeg', 'image/jpg', 'image/gif', 'image/svg+xml', 'application/pdf', 'application/msword', 'application/vnd.openxmlformats-word', 'application/vnd.ms-excel', 'application/vnd.openxmlformats-excel', 'text/plain', 'video/mp4', 'video/avi', 'video/mpeg', 'video/quicktime', 'audio/mpeg', 'audio/wav', 'audio/ogg', 'application/zip', 'application/x-rar-compressed', 'application/octet-stream');
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 CREATE TYPE "public"."user_role" AS ENUM('coach', 'athlete');
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
CREATE TABLE IF NOT EXISTS "exams" (
	"id" serial PRIMARY KEY NOT NULL,
	"club_id" integer,
	"image_id" integer,
	"title" text NOT NULL,
	"description" text,
	"due_at" timestamp DEFAULT now() + interval '1 day',
	"created_at" timestamp DEFAULT now(),
	"updated_at" timestamp DEFAULT now()
);
--> statement-breakpoint
CREATE TABLE IF NOT EXISTS "exam_questions" (
	"id" serial PRIMARY KEY NOT NULL,
	"exam_id" integer,
	"media_id" integer,
	"type" "question_type" NOT NULL,
	"content" text NOT NULL,
	"answer" text NOT NULL,
	"created_at" timestamp DEFAULT now(),
	"updated_at" timestamp DEFAULT now()
);
--> statement-breakpoint
CREATE TABLE IF NOT EXISTS "clubs" (
	"id" serial PRIMARY KEY NOT NULL,
	"creator_id" integer,
	"image_id" integer,
	"name" text NOT NULL,
	"description" text NOT NULL,
	"type" "sport_type" NOT NULL,
	"created_at" timestamp DEFAULT now(),
	"updated_at" timestamp DEFAULT now()
);
--> statement-breakpoint
CREATE TABLE IF NOT EXISTS "program_exercises" (
	"id" serial PRIMARY KEY NOT NULL,
	"program_id" integer,
	"media_id" integer,
	"name" text NOT NULL,
	"description" text,
	"repetition" integer,
	"sets" integer,
	"rest" integer,
	"created_at" timestamp DEFAULT now(),
	"updated_at" timestamp DEFAULT now()
);
--> statement-breakpoint
CREATE TABLE IF NOT EXISTS "programs" (
	"id" serial PRIMARY KEY NOT NULL,
	"club_id" integer,
	"image_id" integer,
	"name" text NOT NULL,
	"start_date" timestamp DEFAULT now(),
	"end_date" timestamp DEFAULT now(),
	"created_at" timestamp DEFAULT now(),
	"updated_at" timestamp DEFAULT now()
);
--> statement-breakpoint
CREATE TABLE IF NOT EXISTS "tacticals" (
	"id" serial PRIMARY KEY NOT NULL,
	"club_id" integer,
	"image_id" integer,
	"sport_type" text NOT NULL,
	"name" text NOT NULL,
	"description" text,
	"content" json,
	"created_at" timestamp DEFAULT now(),
	"updated_at" timestamp DEFAULT now()
);
--> statement-breakpoint
CREATE TABLE IF NOT EXISTS "medias" (
	"id" serial PRIMARY KEY NOT NULL,
	"creator_id" integer,
	"name" text NOT NULL,
	"description" text,
	"file_size" integer,
	"path" text,
	"type" "media_type" NOT NULL,
	"parent" "media_parent" NOT NULL,
	"url" text NOT NULL,
	"created_at" timestamp DEFAULT now(),
	"updated_at" timestamp DEFAULT now()
);
--> statement-breakpoint
CREATE TABLE IF NOT EXISTS "users" (
	"id" serial PRIMARY KEY NOT NULL,
	"name" text NOT NULL,
	"email" text NOT NULL,
	"image" text DEFAULT 'https://api.dicebear.com/9.x/adventurer/png' NOT NULL,
	"password" text NOT NULL,
	"phone" text,
	"role" "user_role" DEFAULT 'athlete' NOT NULL,
	"expertise" text,
	"created_at" timestamp DEFAULT now(),
	"updated_at" timestamp DEFAULT now(),
	CONSTRAINT "users_email_unique" UNIQUE("email")
);
--> statement-breakpoint
CREATE TABLE IF NOT EXISTS "users_to_clubs" (
	"user_id" serial NOT NULL,
	"club_id" serial NOT NULL,
	"role" "user_role" DEFAULT 'athlete' NOT NULL,
	"created_at" timestamp DEFAULT now(),
	CONSTRAINT "users_to_clubs_user_id_club_id_pk" PRIMARY KEY("user_id","club_id")
);
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "exams" ADD CONSTRAINT "exams_club_id_clubs_id_fk" FOREIGN KEY ("club_id") REFERENCES "public"."clubs"("id") ON DELETE cascade ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "exams" ADD CONSTRAINT "exams_image_id_medias_id_fk" FOREIGN KEY ("image_id") REFERENCES "public"."medias"("id") ON DELETE set null ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "exam_questions" ADD CONSTRAINT "exam_questions_exam_id_clubs_id_fk" FOREIGN KEY ("exam_id") REFERENCES "public"."clubs"("id") ON DELETE cascade ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "exam_questions" ADD CONSTRAINT "exam_questions_media_id_medias_id_fk" FOREIGN KEY ("media_id") REFERENCES "public"."medias"("id") ON DELETE set null ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "clubs" ADD CONSTRAINT "clubs_creator_id_users_id_fk" FOREIGN KEY ("creator_id") REFERENCES "public"."users"("id") ON DELETE set null ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "clubs" ADD CONSTRAINT "clubs_image_id_medias_id_fk" FOREIGN KEY ("image_id") REFERENCES "public"."medias"("id") ON DELETE set null ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "program_exercises" ADD CONSTRAINT "program_exercises_program_id_programs_id_fk" FOREIGN KEY ("program_id") REFERENCES "public"."programs"("id") ON DELETE cascade ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "program_exercises" ADD CONSTRAINT "program_exercises_media_id_medias_id_fk" FOREIGN KEY ("media_id") REFERENCES "public"."medias"("id") ON DELETE set null ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "programs" ADD CONSTRAINT "programs_club_id_clubs_id_fk" FOREIGN KEY ("club_id") REFERENCES "public"."clubs"("id") ON DELETE cascade ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "programs" ADD CONSTRAINT "programs_image_id_medias_id_fk" FOREIGN KEY ("image_id") REFERENCES "public"."medias"("id") ON DELETE set null ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "tacticals" ADD CONSTRAINT "tacticals_club_id_clubs_id_fk" FOREIGN KEY ("club_id") REFERENCES "public"."clubs"("id") ON DELETE cascade ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "tacticals" ADD CONSTRAINT "tacticals_image_id_medias_id_fk" FOREIGN KEY ("image_id") REFERENCES "public"."medias"("id") ON DELETE set null ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "medias" ADD CONSTRAINT "medias_creator_id_users_id_fk" FOREIGN KEY ("creator_id") REFERENCES "public"."users"("id") ON DELETE set null ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "users_to_clubs" ADD CONSTRAINT "users_to_clubs_user_id_users_id_fk" FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE cascade ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "users_to_clubs" ADD CONSTRAINT "users_to_clubs_club_id_clubs_id_fk" FOREIGN KEY ("club_id") REFERENCES "public"."clubs"("id") ON DELETE cascade ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
