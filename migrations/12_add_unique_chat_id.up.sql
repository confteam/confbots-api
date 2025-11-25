ALTER TABLE channels
ADD CONSTRAINT channels_code_unique UNIQUE (code),
ADD CONSTRAINT channels_admin_chat_id_unique UNIQUE (admin_chat_id),
ADD CONSTRAINT channels_channel_chat_id_unique UNIQUE (channel_chat_id),
ADD CONSTRAINT channels_discussions_chat_id_unique UNIQUE (discussions_chat_id);
