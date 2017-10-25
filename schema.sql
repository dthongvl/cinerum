CREATE TABLE users (
  id           INTEGER PRIMARY KEY,
  username     VARCHAR(255),
  password     VARCHAR(255),
  live_at      VARCHAR(255),
  stream_title VARCHAR(255),
  stream_key   VARCHAR(255)
);

INSERT INTO users (username, password, live_at, stream_title, stream_key)
VALUES ('dthongvl', '123456', '', 'Untitled', 'live_key_zxcqwertyuiop');

INSERT INTO users (username, password, live_at, stream_title, stream_key)
VALUES ('dquang', '123456', '', 'Untitled', 'live_key_zxcqwevsvsvsfv');