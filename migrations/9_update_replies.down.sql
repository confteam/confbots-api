ALTER TABLE replies
DROP COLUMN channel_id,
DROP CONSTRAINT replies_take_id_channel_id_unique;
