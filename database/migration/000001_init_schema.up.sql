
CREATE TABLE IF NOT EXISTS "users" (
    "id" uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    "username" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    );

CREATE TABLE IF NOT EXISTS "match" (
    "user_id" UUID REFERENCES "users"(id) ON DELETE CASCADE,
    "match_id" UUID REFERENCES "users"(id) ON DELETE CASCADE,
);