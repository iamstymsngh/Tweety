-- SCHEMA FOR TWEETY
DROP DATABASE IF EXISTS tweety CASCADE;
CREATE DATABASE IF NOT EXISTS tweety;
SET DATABASE = tweety;

-- Create USERS table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL NOT NULL PRIMARY KEY,
    email varchar NOT NULL UNIQUE,
    username varchar NOT NULL UNIQUE
);