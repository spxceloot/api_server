CREATE TABLE Users IF NOT EXISTS (
  uid varchar PRIMARY KEY,
  email varchar UNIQUE,
  password varchar,
  username varchar,
  created_a timestamp
);
