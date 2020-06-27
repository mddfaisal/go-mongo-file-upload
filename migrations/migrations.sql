CREATE TABLE email (
    id serial PRIMARY KEY,
    from_email TEXT,
    to_email TEXT,
    cc_email TEXT,
    date_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    subject TEXT NOT NULL,
    muid TEXT NOT NULL UNIQUE,
    mongo_id TEXT
);