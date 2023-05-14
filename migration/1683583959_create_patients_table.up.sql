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

ALTER TABLE patients
ADD COLUMN bmi FLOAT GENERATED ALWAYS AS (weight_kg / (height_m * height_m)) STORED;

CREATE OR REPLACE FUNCTION update_bmi()
RETURNS TRIGGER AS $$
BEGIN
    NEW.bmi = NEW.weight_kg / (NEW.height_m * NEW.height_m);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER bmi_update_trigger
BEFORE INSERT OR UPDATE OF weight_kg, height_m
ON patients
FOR EACH ROW
EXECUTE FUNCTION update_bmi();

COMMIT;
