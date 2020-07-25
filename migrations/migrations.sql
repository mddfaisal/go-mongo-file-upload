CREATE TABLE email (
    id serial PRIMARY KEY,
    from_email TEXT,
    to_email TEXT,
    cc_email TEXT,
    date_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    subject TEXT NOT NULL,
    muid TEXT NOT NULL UNIQUE,
    content_type TEXT,
    mongo_id TEXT
);

CREATE TABLE upload_path (
    id serial PRIMARY KEY,
    email_id INT,
    file_path TEXT,
    CONSTRAINT fk_email FOREIGN KEY (email_id) REFERENCES email(id)
    ON UPDATE NO ACTION ON DELETE NO ACTION
);