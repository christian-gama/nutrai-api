BEGIN;

CREATE TABLE "diet_plans" (
    "id" BIGSERIAL PRIMARY KEY,
    "diet_id" BIGINT NOT NULL,
    "text" VARCHAR(1000) NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX "uidx__plans__diet_id" ON "diet_plans" ("diet_id");

ALTER TABLE "diet_plans"
ADD CONSTRAINT "fk__diet_id__diets.id"
FOREIGN KEY ("diet_id")
REFERENCES "diet_diets" ("id")
ON DELETE CASCADE ON UPDATE CASCADE;

COMMIT;
