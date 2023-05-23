BEGIN;

CREATE TABLE "exception_exceptions" (
    "id" BIGSERIAL PRIMARY KEY,
    "message" TEXT NOT NULL,
    "stack" TEXT NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

COMMIT;
