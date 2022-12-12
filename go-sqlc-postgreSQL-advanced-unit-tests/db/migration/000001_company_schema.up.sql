CREATE TABLE "companies" (
  "id" serial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "headquarters" varchar NOT NULL,
  "founded" bigint NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX ON "companies" ("owner");
