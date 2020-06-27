#!/bin/sh
set -e

psql -e -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-END
	\timing
	CREATE USER gmail_dba WITH PASSWORD 'pass';
	CREATE DATABASE gmail_db OWNER gmail_dba;
	GRANT ALL PRIVILEGES ON DATABASE gmail_db TO gmail_dba;
	CREATE EXTENSION tablefunc;
END

export PGUSER=gmail_dba
export PGPASSWORD=pass
psql -d gmail_db < /src/repo/migrations/migrations.sql