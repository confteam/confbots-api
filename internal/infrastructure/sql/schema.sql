CREATE TABLE channels (
  id SERIAL PRIMARY KEY,
  code TEXT NOT NULL,
  channel_chat_id BIGINT,
  admin_chat_id BIGINT,
  discussions_chat_id BIGINT,
  decorations TEXT
);

CREATE TABLE bots (
  id SERIAL PRIMARY KEY,
  tgid BIGINT NOT NULL,
  type TEXT NOT NULL,
  channel_id INTEGER REFERENCES channels(id),

  CONSTRAINT bots_tgid_type_unique UNIQUE (tgid, type)
);
