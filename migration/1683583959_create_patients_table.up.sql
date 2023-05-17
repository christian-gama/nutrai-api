BEGIN;

CREATE TABLE "patients" (
    "id" BIGSERIAL PRIMARY KEY,
    "weight_kg" FLOAT NOT NULL,
    "height_m" FLOAT NOT NULL,
    "age" INTEGER NOT NULL
);

CREATE UNIQUE INDEX "uidx__patients__id" ON patients ("id");

ALTER TABLE "patients"
ADD CONSTRAINT "fk__id__users.id"
FOREIGN KEY ("id") 
REFERENCES "users" ("id")
ON DELETE CASCADE ON UPDATE CASCADE;

-- The BMI is a calculated field that depends on the weight and height. The code below
-- creates a column that is calculated on the fly and stored in the table. The column
-- will be updated automatically whenever the weight or height is updated by the user, 
-- using a trigger (bmi_update_trigger). The trigger calls the update_bmi() function.
-- The function is written in PL/pgSQL, which is the PostgreSQL procedural language.
-- Refer to BMI calculation: https://www.cdc.gov/nccdphp/dnpao/growthcharts/training/bmiage/page5_1.html
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
