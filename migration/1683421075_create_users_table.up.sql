BEGIN;

CREATE TABLE users (
    "id" BIGSERIAL PRIMARY KEY,
    "name" VARCHAR(100) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "password" VARCHAR(32) NOT NULL
);

CREATE UNIQUE INDEX "uidx__users__email" ON users ("email");

COMMIT;
