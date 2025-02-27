CREATE TABLE if NOT EXISTS users(
    id SERIAL,
    name TEXT PRIMARY KEY,
    email TEXT,
    password TEXT,
    salt TEXT,
    token TEXT DEFAULT 1,
    role TEXT
);
