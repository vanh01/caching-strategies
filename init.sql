CREATE TABLE IF NOT EXISTS "user" (
  "id"   uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "name" text NOT NULL
)


insert into "user" ("id", "name")
values
('7aface47-7ce7-4b6a-ba16-392d19aa2785', 'test1'),
('f51ee719-58fc-4f5f-9b32-3ebaa7648e52', 'test2');