-- migrate:up
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

CREATE TABLE Fight(
    Id INTEGER PRIMARY KEY,

    Fighter0Id INTEGER NOT NULL REFERENCES Person(Id),
    Fighter1Id INTEGER NOT NULL REFERENCES Person(Id),
    RefereeId INTEGER NOT NULL REFERENCES Person(Id),
    ArenaId INTEGER NOT NULL,
    FOREIGN KEY (ArenaId) REFERENCES Person(Id),

    OfflineWatchers INTEGER NOT NULL,
    OnlineWatchers INTEGER NOT NULL,
    MaxRounds INTEGER NOT NULL,
    WinPrize FLOAT DEFAULT 0.0 NOT NULL,
    -- Prize issued in any case to every fighter.
    ParticipancePrize FLOAT,

    -- If null: still going.
    EndTypeId INTEGER REFERENCES FightEndType(Id)
);

-- migrate:down

