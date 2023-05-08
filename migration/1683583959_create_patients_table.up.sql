BEGIN;

CREATE TABLE "patients" (
    "id" BIGSERIAL PRIMARY KEY,
    "user_id" BIGINT NOT NULL,
    "weight_kg" FLOAT NOT NULL,
    "height_m" FLOAT NOT NULL,
    "age" INTEGER NOT NULL
);

CREATE UNIQUE INDEX "uidx__patients__user_id" ON patients ("user_id");

ALTER TABLE "patients"
ADD CONSTRAINT "fk__user_id__users.id"
FOREIGN KEY ("user_id") 
REFERENCES "users" ("id")
ON DELETE CASCADE ON UPDATE CASCADE;

COMMIT;
