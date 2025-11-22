ALTER TABLE replies
ADD COLUMN channel_id INTEGER REFERENCES channels(id) NOT NULL,
ADD CONSTRAINT replies_take_id_channel_id_unique UNIQUE (take_id, channel_id);
