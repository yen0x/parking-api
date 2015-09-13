-- Database init
CREATE USER parking WITH UNENCRYPTED PASSWORD 'parking';
CREATE DATABASE "parking";
GRANT ALL ON DATABASE "parking" TO "parking";

-- Switch to the parking db as the parking user.
\connect "parking";
set role "parking";

-- User

CREATE TABLE "user" (
    "uid" text NOT NULL,
    "email" text NOT NULL,
    "firstname" text NOT NULL,
    "lastname" text DEFAULT '',
    "password" text NOT NULL,
    "gender" text DEFAULT 'UNKNOWN',
    "phone" text DEFAULT '',
    "address" text DEFAULT '',
    "creation_time" timestamp with time zone DEFAULT now(),
    "last_update" timestamp with time zone
);

CREATE UNIQUE INDEX ON "user" ("uid");
CREATE INDEX ON "user" ("email");    
    
-- Parking

CREATE TABLE "parking" (
    "uid" text NOT NULL,
    "user_id" TEXT NOT NULL,
    "description" TEXT DEFAULT '',
    "address" TEXT DEFAULT '',
    "zip" TEXT DEFAULT '',
    "city" TEXT DEFAULT '',
    "latitude" DOUBLE PRECISION DEFAULT 0.0,
    "longitude" DOUBLE PRECISION DEFAULT 0.0,
    "daily_price" integer DEFAULT 0,
    "currency" text default 'EUR',
    "creation_time" timestamp with time zone DEFAULT now(),
    "last_update" timestamp with time zone
);

CREATE UNIQUE INDEX ON "parking" ("uid");
CREATE INDEX ON "parking" ("user_id");
