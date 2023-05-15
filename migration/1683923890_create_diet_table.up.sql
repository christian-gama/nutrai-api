BEGIN;

CREATE TABLE "diets" (
    "id" BIGSERIAL PRIMARY KEY,
    "patient_id" BIGINT NOT NULL,
    "name" VARCHAR(100) NOT NULL,
    "description" VARCHAR(255) NOT NULL,
    "duration_in_weeks" INTEGER NOT NULL,
    "goal" VARCHAR(100) NOT NULL,
    "meal_plan" VARCHAR(100) NOT NULL,
    "monthly_cost_usd" DECIMAL(10, 2) NOT NULL,
    "restricted_food" VARCHAR(100)[] NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX "uidx__diets__patient_id" ON diets ("patient_id");


ALTER TABLE "diets"
ADD CONSTRAINT "fk__patient_id__patients.id"
FOREIGN KEY ("patient_id")
REFERENCES "patients" ("id")
ON DELETE CASCADE ON UPDATE CASCADE;

COMMIT;
