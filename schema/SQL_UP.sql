CREATE TABLE user (
  id PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  username VARCHAR(255) UNIQUE NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  password_hash VARCHAR(255) NOT NULL,
);

CREATE TABLE tag (
  id PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
);

CREATE TABLE beat (
  id PRIMARY KEY,
  bpm INT NOT NULL,
  key VARCHAR(16) NOT NULL,
  path VARCHAR(255) NOT NULL,
  tags [] INT REFERENCES tag(id) ON DELETE CASCADE NOT NULL,
  price FLOAT NOT NULL,
);

CREATE TABLE users_beat (
  id PRIMARY KEY,
  user_id INT REFERENCES user (id) ON DELETE CASCADE NOT NULL,
  beat_id INT REFERENCES beat (id) ON DELETE CASCADE NOT NULL,
  bought BOOLEAN DEFAULT (FALSE),
  -- if bought=false -> it is beat which user sells
  -- if bought=true -> it is beat which user bought 
);