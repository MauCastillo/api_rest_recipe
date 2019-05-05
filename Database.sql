cockroach sql --insecure
create database company_db;
set database = company_db;


CREATE TABLE "recipe_object" (
    "id" SERIAL,
    "title" STRING(100),
    "ingredients" STRING(4096),
    "designation" STRING(4096),
    "preparation" STRING(4096),
    "updated_at" TIMESTAMPTZ,
    PRIMARY KEY ("id")
);