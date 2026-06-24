CREATE TABLE books (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    author TEXT NOT NULL,
    published_year INTEGER,
    status TEXT CHECK(status IN (
        'SUGGESTED',
        'READING',
        'DISCARDED',
        'COMPLETED',
        'SAVED'
    )) NOT NULL
    -- TODO: add genre column later
);