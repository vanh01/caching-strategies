CREATE TABLE IF NOT EXISTS "user" (
  "id"   uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "name" text NOT NULL
)
