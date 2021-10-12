psql << EOF
    CREATE USER "auth" WITH PASSWORD 'P@ssword';
    ALTER ROLE "auth" SUPERUSER;
    CREATE DATABASE "authdb";
    GRANT ALL PRIVILEGES ON DATABASE "authdb" TO "auth";

    \c authdb;

    CREATE TABLE users (
      id serial PRIMARY KEY,
      username VARCHAR(100) UNIQUE NOT NULL,
      password_hash CHAR(64) NOT NULL
    );
EOF
