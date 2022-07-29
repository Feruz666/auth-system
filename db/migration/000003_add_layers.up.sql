CREATE TABLE "layers" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "date" varchar NOT NULL,
  "layer" varchar NOT NULL
);


ALTER TABLE "users" ADD COLUMN "username" VARCHAR NOT NULL;