CREATE TABLE IF NOT EXISTS reactionTypes (
  id CHAR(27) PRIMARY KEY,
  reacts VARCHAR(24) UNIQUE NOT NULL
);

CREATE TABLE reactions(
    id VARCHAR PRIMARY KEY ,
    reactionType VARCHAR NOT NULL ,
    postId VARCHAR Not NULL ,
    userId VARCHAR Not NULL

);
