DROP TABLE IF EXISTS students;

CREATE TABLE students (
  id VARCHAR(255) PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  age INTEGER NOT NULL
);

DROP TABLE IF EXISTS tests;

CREATE TABLE tests (
  id VARCHAR(255) PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);
