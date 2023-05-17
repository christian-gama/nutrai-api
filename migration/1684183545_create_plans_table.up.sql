BEGIN;

CREATE TABLE "plans" (
    "id" BIGSERIAL PRIMARY KEY,
    "diet_id" BIGINT NOT NULL,
    "text" VARCHAR(1000) NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX "uidx__plans__diet_id" ON plans ("diet_id");

ALTER TABLE "plans"
ADD CONSTRAINT "fk__diet_id__diets.id"
FOREIGN KEY ("diet_id")
REFERENCES "diets" ("id")
ON DELETE CASCADE ON UPDATE CASCADE;

COMMIT;
