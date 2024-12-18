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

CREATE TABLE FighterStyle(
    Id INTEGER PRIMARY KEY,
    -- In-Fighter
    -- Out-Boxer
    -- Slugger
    -- Boxer-Puncher
    -- https://en.wikipedia.org/wiki/Boxing_styles_and_technique
    StyleKey TEXT UNIQUE NOT NULL,
)

CREATE TABLE FighterSkills(
    Id INTEGER PRIMARY KEY,
    PersonId INTEGER REFERENCES Person(Id),

    Strength INTEGER NOT NULL,
    Agility INTEGER NOT NULL,
    Endurance INTEGER NOT NULL,
    Speed INTEGER NOT NULL,
    Intelligence INTEGER NOT NULL,

    StyleId INTEGER REFERENCES FighterStyle(Id)
);

CREATE TABLE Person(
    Id INTEGER PRIMARY KEY,
    TypeId INTEGER NOT NULL REFERENCES PersonType(Id),

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

CREATE TABLE RoundEndType(
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
    TimelineDay INTEGER,

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
);

CREATE TABLE FightRound(
    Id INTEGER PRIMARY KEY,
    FightId INTEGER NOT NULL REFERENCES Fight(Id),

    KnockdownsBy0 INTEGER,
    KnockdownsBy1 INTEGER,

    -- 0: Win by 0, 1: Win by 1, 2: Draw.
    Evaluation INTEGER,
    CHECK (Evaluation > 0 AND Evaluation < 3),

    -- If null: still going.
    EndTypeId INTEGER REFERENCES RoundEndType(Id),
);

-- migrate:down

