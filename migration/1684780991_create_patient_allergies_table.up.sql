BEGIN;

CREATE TABLE "patient_allergies" (
  "id" BIGSERIAL PRIMARY KEY,
  "patient_id" BIGINT NOT NULL,
  "name" VARCHAR(100) NOT NULL
);

CREATE INDEX "idx__allergies__patient_id" ON "patient_allergies" ("patient_id");

CREATE UNIQUE INDEX "uidx__allergies__name"
ON "patient_allergies" ("name", "patient_id");

ALTER TABLE "patient_allergies"
ADD CONSTRAINT "fk__patient_id__patients.id"
FOREIGN KEY ("patient_id")
REFERENCES "patient_patients" ("id")
ON DELETE CASCADE ON UPDATE CASCADE;


COMMIT;
