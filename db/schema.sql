CREATE TABLE IF NOT EXISTS "schema_migrations" (version varchar(128) primary key);
CREATE TABLE PersonType (
    Id INTEGER PRIMARY KEY,
    TypeKey TEXT
);
-- Dbmate schema migrations
INSERT INTO "schema_migrations" (version) VALUES
  ('20241217214956');
