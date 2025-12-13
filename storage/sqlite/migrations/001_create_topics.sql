CREATE TABLE topics (
    id TEXT PRIMARY KEY,
    tag TEXT NOT NULL,
    skipped INTEGER NOT NULL DEFAULT 0,
    completed INTEGER NOT NULL DEFAULT 0,
    skippable INTEGER NOT NULL DEFAULT 1,
    last_recall INTEGER
);