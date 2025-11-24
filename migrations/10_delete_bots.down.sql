CREATE TABLE bots (
  id SERIAL PRIMARY KEY,
  tgid INTEGER NOT NULL,
  type TEXT NOT NULL,
  channel_id INTEGER REFERENCES channels(id),
  CONSTRAINT bots_tgid_type_unique UNIQUE (tgid, type)
);
