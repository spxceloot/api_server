CREATE TABLE IF NOT EXISTS Users  (
  uid varchar PRIMARY KEY,
  email varchar UNIQUE,
  password varchar,
  username varchar,
  created_a timestamp
);
