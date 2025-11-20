CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  tgid BIGINT NOT NULL
);

CREATE TABLE user_channels (
  id SERIAL PRIMARY KEY,
  user_id INTEGER REFERENCES users(id) NOT NULL,
  channel_id INTEGER REFERENCES channels(id) NOT NULL,
  role TEXT DEFAULT 'MEMBER',
  anonimity BOOLEAN DEFAULT TRUE,
  CONSTRAINT user_channels_user_id_channel_id_unique UNIQUE (user_id, channel_id)
);
