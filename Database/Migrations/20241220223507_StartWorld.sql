-- migrate:up
INSERT INTO World (LastProcessedStateEventId) VALUES (-1);
INSERT INTO Timeline (CurrentDay) VALUES (0);

-- migrate:down

