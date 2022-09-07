CREATE TABLE IF NOT EXISTS account (
  id BIGSERIAL PRIMARY KEY,
  name TEXT,
  username TEXT UNIQUE NOT NULL,
  email TEXT UNIQUE NOT NULL,
  photo_path TEXT,
  password_hash TEXT NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS artist (
  id BIGINT REFERENCES account (id) ON DELETE CASCADE UNIQUE NOT NULL,
  artist_name TEXT NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS beat (
  id BIGSERIAL PRIMARY KEY,
  artist_id BIGINT REFERENCES artist (id) ON DELETE CASCADE NOT NULL,
  name TEXT NOT NULL,
  bpm TEXT NOT NULL,
  key TEXT NOT NULL,
  photo_path TEXT NOT NULL,
  mp3_path TEXT NOT NULL,
  wav_path TEXT,
  likes BIGINT DEFAULT 0,
  genre TEXT DEFAULT 'All',
  mood TEXT DEFAULT 'All',
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS tag (
  id BIGSERIAL PRIMARY KEY,
  beat_id BIGINT REFERENCES beat (id) ON DELETE CASCADE NOT NULL,
  name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS price (
  id BIGSERIAL PRIMARY KEY,
  beat_id BIGINT REFERENCES beat (id) ON DELETE CASCADE UNIQUE NOT NULL,
  standart_price TEXT NOT NULL,
  premium_price TEXT,
  unlimited_price TEXT
);

CREATE TABLE IF NOT EXISTS account_beat (
  id BIGSERIAL PRIMARY KEY,
  account_id BIGINT REFERENCES account (id) ON DELETE CASCADE NOT NULL,
  beat_id BIGINT REFERENCES beat (id) ON DELETE CASCADE NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS playlist (
  id BIGSERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS playlist_beat (
  playlist_id BIGINT REFERENCES playlist (id) ON DELETE CASCADE NOT NULL,
  beat_id BIGINT REFERENCES beat (id) ON DELETE CASCADE NOT NULL,
  PRIMARY KEY (playlist_id, beat_id)
);

CREATE TABLE IF NOT EXISTS account_playlist (
  account_id BIGINT REFERENCES account (id) ON DELETE CASCADE NOT NULL,
  playlist_id BIGINT REFERENCES playlist (id) ON DELETE CASCADE NOT NULL,
  PRIMARY KEY (playlist_id, account_id)
);

-- QUERIES
-- Create Account
INSERT INTO
  account (
    name,
    username,
    email,
    photo_path,
    password_hash,
    created_at
  )
VALUES
  (
    'Alexander Turok',
    'alexander_turok',
    'alexander.turok2004@gmail.com',
    'avatar.jpeg',
    'some_password_hash',
    CURRENT_TIMESTAMP
  );

-- Select Account
SELECT
  *
FROM
  account;

-- Create Artist
INSERT INTO
  artist (id, artist_name, created_at)
VALUES
  (1, 'Alex Beats', CURRENT_TIMESTAMP);

-- Select Artist
SELECT
  *
FROM
  artist;

-- Select Account's Artist
SELECT
  *
FROM
  artist
WHERE
  id = 1;

--or
SELECT
  *
FROM
  account
  JOIN artist ON artist.id = account.id;

-- Create beat
INSERT INTO
  beat (
    artist_id,
    name,
    bpm,
    key,
    photo_path,
    mp3_path,
    wav_path,
    genre,
    mood,
    created_at
  )
VALUES
  (
    1,
    'Moonlight beat',
    '96',
    'Amaj5',
    'album.jpeg',
    'file.mp3',
    'file.wav',
    'electro',
    'energy',
    CURRENT_TIMESTAMP
  );

INSERT INTO
  tag (beat_id, name)
VALUES
  (1, 'moonlight');

INSERT INTO
  tag (beat_id, name)
VALUES
  (1, 'alex-beat');

INSERT INTO
  price (
    beat_id,
    standart_price,
    premium_price,
    unlimited_price
  )
VALUES
  (1, '9.99$', '19.99$', '49.99$');

-- Select Beat
SELECT
  *
FROM
  beat
  JOIN price ON price.beat_id = beat.id
  JOIN tag ON tag.beat_id = beat.id;

-- Select Artisit's beat
SELECT
  *
FROM
  beat
  JOIN price ON price.beat_id = beat.id
  JOIN tag ON tag.beat_id = beat.id
WHERE
  artist_id = 1;

-- Craete Account Beat (when account buys beat)
INSERT INTO
  account_beat (account_id, beat_id, created_at)
VALUES
  (1, 1, CURRENT_TIMESTAMP);

-- Select Accounts Bought Beat
SELECT
  *
FROM
  account_beat
  JOIN account ON account.id = account_beat.account_id
  JOIN beat ON beat.id = account_beat.beat_id
  JOIN price ON price.beat_id = beat.id
  JOIN tag ON tag.beat_id = beat.id;

-- Create Playlist
INSERT INTO
  playlist (name, created_at)
VALUES
  ('My Playlist', CURRENT_TIMESTAMP);

INSERT INTO
  account_playlist (account_id, playlist_id)
VALUES
  (1, 1);

-- Insert Beat Into Playlist
INSERT INTO
  playlist_beat (playlist_id, beat_id)
VALUES
  (1, 1);

-- Select Account's Playlist 
SELECT
  *
FROM
  account
  JOIN account_playlist ON account_playlist.account_id = account.id
  JOIN playlist ON playlist.id = account_playlist.playlist_id;

-- Select Playlist's Beat
SELECT
  *
FROM
  playlist_beat
  JOIN beat ON beat.id = playlist_beat.beat_id
WHERE
  playlist_beat.playlist_id = 1;