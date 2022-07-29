CREATE TABLE "Layers" (
  "id" bigserial PRIMARY KEY,
  "user_id" int NOT NULL,
  "date" timestamp NOT NULL,
  "layer" varchar NOT NULL
);
