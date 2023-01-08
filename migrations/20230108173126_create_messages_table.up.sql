CREATE TABLE "messages" (
    "id" BIGSERIAL PRIMARY KEY,
    "pid" VARCHAR(12) NOT NULL UNIQUE,
    "text" TEXT NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP NOT NULL
);
