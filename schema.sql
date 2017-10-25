CREATE TABLE users (
  id           INTEGER PRIMARY KEY AUTOINCREMENT,
  username     VARCHAR(255),
  password     VARCHAR(255),
  is_display   INTEGER,
  is_private   INTEGER,
  live_at      VARCHAR(255),
  stream_title VARCHAR(255),
  stream_key   VARCHAR(255)
);

INSERT INTO users (username, password, is_display, is_private, live_at, stream_title, stream_key)
VALUES ('dthongvl', '123456', 1, 0, '', 'Untitled', 'live_key_zxcqwertyuiop');

INSERT INTO users (username, password, is_display, is_private, live_at, stream_title, stream_key)
VALUES ('dquang', '123456', 0, 1, '', 'Untitled', 'live_key_zxcqwevsvsvsfv');