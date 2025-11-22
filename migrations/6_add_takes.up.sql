CREATE TABLE takes (
  id SERIAL PRIMARY KEY,
  status TEXT DEFAULT 'PENDING',
  user_message_id BIGINT NOT NULL,
  admin_message_id BIGINT NOT NULL,
  user_channel_id INTEGER REFERENCES user_channels(id) NOT NULL,
  channel_id INTEGER REFERENCES channels(id) NOT NULL
);
