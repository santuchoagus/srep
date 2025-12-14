DROP TABLE IF EXISTS global_state;
CREATE TABLE global_state (
    key TEXT PRIMARY KEY,
    value TEXT
);

INSERT INTO global_state
(key, value)
VALUES
('CURRENT_TOPIC', 'foo');