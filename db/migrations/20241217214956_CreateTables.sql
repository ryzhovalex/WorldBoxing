-- migrate:up
CREATE TABLE IF NOT EXISTS PersonType(
    Id INTEGER PRIMARY KEY,
    TypeKey VARCHAR(255),
);

CREATE TABLE IF NOT EXISTS Person(
    Id INTEGER PRIMARY KEY,
    PersonType INTEGER,
    FOREIGN KEY (PersonType) REFERENCES PersonType(id)

    Firstname VARCHAR(255),
    Surname VARCHAR(255),
    BornDay INTEGER,
);

CREATE TABLE IF NOT EXISTS Timeline(
    Id INTEGER PRIMARY KEY,
    CurrentDay INTEGER,
);

CREATE TABLE IF NOT EXISTS WorldEvent(
    -- Id serve as event order.
    Id INTEGER PRIMARY KEY,
    Body JSONB,
    TimelineDay INTEGER,
);

-- migrate:down

