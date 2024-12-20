-- migrate:up
CREATE TABLE Sponsorship(
    Id INTEGER PRIMARY KEY,
    PersonId INTEGER NOT NULL REFERENCES Person(Id),
    SponsorId INTEGER NOT NULL REFERENCES Company(Id)
);

CREATE TABLE Company(
    Id INTEGER PRIMARY KEY,
    Name TEXT UNIQUE NOT NULL,
    NetWorth FLOAT NOT NULL DEFAULT 0.0
);

CREATE TABLE FighterSkill(
    Id INTEGER PRIMARY KEY,
    PersonId INTEGER REFERENCES Person(Id),

    Strength INTEGER NOT NULL,
    Agility INTEGER NOT NULL,
    Endurance INTEGER NOT NULL,
    Speed INTEGER NOT NULL,
    Intelligence INTEGER NOT NULL,

    -- Variants:
    -- * In-Fighter
    -- * Out-Boxer
    -- * Slugger
    -- * Boxer-Puncher
    -- https://en.wikipedia.org/wiki/Boxing_styles_and_technique
    Style TEXT UNIQUE NOT NULL
);

CREATE TABLE Person(
    Id INTEGER PRIMARY KEY,
    Type TEXT NOT NULL,

    Rating INTEGER DEFAULT 0,

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

CREATE TABLE World(
    Id INTEGER PRIMARY KEY,
    -- Viable to keep track on current level of synchronization between events
    -- and database state.
    LastProcessedStateEventId INTEGER REFERENCES StateEvent(Id)
    SimulationStarted INTEGER NOT NULL DEFAULT 0 CHECK (Simulated IN (0, 1))
);

CREATE TABLE StateEvent(
    -- Id serves as StateEvent order.
    Id INTEGER PRIMARY KEY,
    Type TEXT NOT NULL,
    Body JSONB,
    Time INTEGER NOT NULL,
    TimelineDay INTEGER NOT NULL
);

CREATE TABLE City(
    Id INTEGER PRIMARY KEY,
    Name TEXT UNIQUE NOT NULL,
    CountryId INTEGER NOT NULL REFERENCES Country(Id),
    Population INTEGER
    Capital INTEGER CHECK (Capital IN (0, 1))
    Archetype TEXT
    Influence FLOAT CHECK (Influence BETWEEN 0.0 AND 1.0)
);

CREATE TABLE Country(
    Id INTEGER PRIMARY KEY,
    Name TEXT UNIQUE NOT NULL
);

CREATE TABLE Fight(
    Id INTEGER PRIMARY KEY,
    TimelineDay INTEGER,

    Fighter0Id INTEGER NOT NULL REFERENCES Person(Id),
    Fighter1Id INTEGER NOT NULL REFERENCES Person(Id),
    RefereeId INTEGER NOT NULL REFERENCES Person(Id),
    CityId INTEGER NOT NULL REFERENCES City(Id),

    OfflineWatchers INTEGER NOT NULL,
    OnlineWatchers INTEGER NOT NULL,
    MaxRounds INTEGER NOT NULL,
    WinPrize FLOAT DEFAULT 0.0 NOT NULL,
    LosePrize FLOAT DEFAULT 0.0 NOT NULL
);

CREATE TABLE FightRound(
    Id INTEGER PRIMARY KEY,
    FightId INTEGER NOT NULL REFERENCES Fight(Id),

    KnockdownsBy0 INTEGER,
    KnockdownsBy1 INTEGER,

    -- 0: Win by 0, 1: Win by 1, 2: Draw.
    Evaluation INTEGER CHECK (Evaluation > 0 AND Evaluation < 3),

    -- If null: still going.
    EndType TEXT
);

-- migrate:down

