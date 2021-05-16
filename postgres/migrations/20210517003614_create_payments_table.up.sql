BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DO $$ BEGIN
    CREATE TYPE "payment_status" AS ENUM (
        'ACTIVE',
        'INACTIVE',
        'EXPIRED'
    );
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;


CREATE TABLE IF NOT EXISTS "payment_codes" (
    "id" VARCHAR(255) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "payment_code" VARCHAR(255) NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "status" payment_status,
    "expiration_date" TIMESTAMP NOT NULL DEFAULT ((now() AT TIME ZONE 'utc') + interval '50 years'),
    "created" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    "updated" TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc')
);

COMMIT;