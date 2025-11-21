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

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  tgid BIGINT NOT NULL UNIQUE
);

CREATE TABLE user_channels (
  id SERIAL PRIMARY KEY,
  user_id INTEGER REFERENCES users(id) NOT NULL,
  channel_id INTEGER REFERENCES channels(id) NOT NULL,
  role TEXT NOT NULL,
  anonimity BOOLEAN DEFAULT TRUE,

  CONSTRAINT user_channels_user_id_channel_id_unique UNIQUE (user_id, channel_id)
);
