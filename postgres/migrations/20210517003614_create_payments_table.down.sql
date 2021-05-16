BEGIN;

DROP TABLE IF EXISTS "payment_codes";

DROP TYPE "payment_status" CASCADE;
DROP EXTENSION "uuid-ossp" CASCADE;

COMMIT;