create table if not exists books
(
    id     INTEGER
        primary key autoincrement,
    title  TEXT
        unique,
    author TEXT
);
