CREATE TABLE replies (
  id SERIAL PRIMARY KEY,
  user_message_id BIGINT NOT NULL,
  admin_message_id BIGINT NOT NULL,
  take_id INTEGER REFERENCES takes(id) NOT NULL
);
