BEGIN;

CREATE TABLE "restricted_foods" (
    "id" BIGSERIAL PRIMARY KEY,
    "diet_id" BIGINT NOT NULL,
    "name" VARCHAR(100) NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX "uidx__restricted_foods__diet_id" ON restricted_foods ("diet_id");

ALTER TABLE "restricted_foods"
ADD CONSTRAINT "fk__diet_id__diets.id"
FOREIGN KEY ("diet_id")
REFERENCES "diets" ("id")
ON DELETE CASCADE ON UPDATE CASCADE;

COMMIT;
