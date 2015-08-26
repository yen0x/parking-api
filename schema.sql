CREATE USER parking WITH UNENCRYPTED PASSWORD 'parking';
CREATE DATABASE "parking";
GRANT ALL ON DATABASE "parking" TO "parking";

-- Switch to the parking db as the parking user.
\connect "parking";
set role "parking";

CREATE TABLE "user" (
    "uid" text NOT NULL,
    "email" text NOT NULL,
    "firstname" text NOT NULL,
    "lastname" text NOT NULL,
    "gender" text DEFAULT 'UNKNOWN',
    "phone" text DEFAULT '',
    "address" text DEFAULT '',
    "creation_time" timestamp with time zone DEFAULT now(),
    "last_update" timestamp with time zone
);

CREATE UNIQUE INDEX ON "user" ("uid");
CREATE INDEX ON "user" ("email");    
    
