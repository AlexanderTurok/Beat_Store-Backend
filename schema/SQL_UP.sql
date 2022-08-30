CREATE TABLE users (
  id SERIAL NOT NULL UNIQUE,
  name VARCHAR(255) NOT NULL,
  username VARCHAR(255) UNIQUE NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE beat (
  id SERIAL NOT NULL UNIQUE,
  bpm INT NOT NULL,
  key VARCHAR(16) NOT NULL,
  path VARCHAR(255) NOT NULL,
  tag VARCHAR(255),
  price FLOAT NOT NULL
);

CREATE TABLE users_beat (
  id SERIAL NOT NULL UNIQUE,
  user_id INT REFERENCES users (id) ON DELETE CASCADE NOT NULL,
  beat_id INT REFERENCES beat (id) ON DELETE CASCADE NOT NULL,
  bought BOOLEAN DEFAULT (FALSE) -- if bought=false -> it is beat which user sells
  -- if bought=true -> it is beat which user bought 
);