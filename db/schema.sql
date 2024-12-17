CREATE TABLE IF NOT EXISTS "schema_migrations" (version varchar(128) primary key);
CREATE TABLE PersonType(
    Id INTEGER PRIMARY KEY,
    TypeKey TEXT UNIQUE NOT NULL
);
CREATE TABLE Sponsorship(
    Id INTEGER PRIMARY KEY,
    PersonId INTEGER NOT NULL REFERENCES Person(Id),
    SponsorId INTEGER NOT NULL REFERENCES Company(Id)
);
CREATE TABLE Company(
    Id INTEGER PRIMARY KEY,
    CompanyName TEXT UNIQUE NOT NULL,
    NetWorth FLOAT NOT NULL DEFAULT 0.0
);
CREATE TABLE Person(
    Id INTEGER PRIMARY KEY,
    TypeId INTEGER REFERENCES PersonType(Id),

    Firstname TEXT NOT NULL,
    Surname TEXT NOT NULL,
    CityId INTEGER NOT NULL REFERENCES City(Id),
    BornDay INTEGER NOT NULL,
    -- Money earned from any sources: fights, sponsorships, side busineses.
    TotalMoneyEarned FLOAT
);
CREATE TABLE Timeline(
    Id INTEGER PRIMARY KEY,
    CurrentDay INTEGER
);
CREATE TABLE WorldEventType(
    Id INTEGER PRIMARY KEY,
    TypeKey TEXT UNIQUE NOT NULL
);
CREATE TABLE WorldEvent(
    -- Id serve as event order.
    Id INTEGER PRIMARY KEY,
    TypeId INTEGER NOT NULL REFERENCES WorldEventType(Id),
    Body JSONB,
    TimelineDay INTEGER
);
CREATE TABLE FightEndType(
    Id INTEGER PRIMARY KEY,
    TypeKey TEXT UNIQUE NOT NULL
);
CREATE TABLE Arena(
    Id INTEGER PRIMARY KEY,
    ArenaName TEXT UNIQUE NOT NULL,
    CityId INTEGER NOT NULL REFERENCES City(Id)
);
CREATE TABLE City(
    Id INTEGER PRIMARY KEY,
    CityName TEXT UNIQUE NOT NULL,
    CountryId INTEGER NOT NULL REFERENCES Country(Id)
);
CREATE TABLE Country(
    Id INTEGER PRIMARY KEY,
    CountryName TEXT UNIQUE NOT NULL
);
CREATE TABLE Fight(
    Id INTEGER PRIMARY KEY,

    Fighter0Id INTEGER NOT NULL REFERENCES Person(Id),
    Fighter1Id INTEGER NOT NULL REFERENCES Person(Id),
    RefereeId INTEGER NOT NULL REFERENCES Person(Id),
    ArenaId INTEGER NOT NULL REFERENCES Arena(Id),

    OfflineWatchers INTEGER NOT NULL,
    OnlineWatchers INTEGER NOT NULL,
    MaxRounds INTEGER NOT NULL,
    WinPrize FLOAT DEFAULT 0.0 NOT NULL,
    -- Prize issued in any case to every fighter.
    ParticipancePrize FLOAT,

    -- If null: still going.
    EndTypeId INTEGER REFERENCES FightEndType(Id),
    RoundsPassed INTEGER DEFAULT 0.0,

    CHECK (RoundsPassed <= MaxRounds)
);
-- Dbmate schema migrations
INSERT INTO "schema_migrations" (version) VALUES
  ('20241217214956');
