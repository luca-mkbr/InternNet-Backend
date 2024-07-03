CREATE DATABASE internnetDatabase

CREATE TABLE events (
    id SERIAL,
    "type" TEXT,
    title TEXT,
    "location" TEXT,
    time TIMESTAMP WITH TIME ZONE,
    userids TEXT[],
	latitude float,
	longitude float
)

CREATE TABLE users (
    id SERIAL,
    first_name TEXT,
    last_name TEXT,
    pass_word TEXT,
    email TEXT,
    company TEXT,
    friends INTEGER[],
    event_ids TEXT[],
    PRIMARY KEY (id)
);