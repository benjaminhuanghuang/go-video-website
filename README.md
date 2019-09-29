# Video website Go + MySql
RESTful API + Front end


## Key point
- native http lib
- native tempalte lib

## Workflow
 handler -> validation (request, user) -> business logic -> reponse

## API Design
- Create user  /user  POST
- Login        /user/:username POST
- Get user info  /user/:username GET
- Delete user    /user/:username DELETE


- List all videos:   /user/:username/videos GET
- Get one videos:   /user/:username/videos/:id GET
- Delete one videos:   /user/:username/videos/:id Delete


- Show video comments:   /videos/:id/comments   GET
- Post video comments:   /videos/:id/comments   POST
- Delete video comments:   /videos/:id/comments   DELETE

## Database Design
- users
```
CREATE TABLE users (
  id  INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  login_name VARCHAR(64)  UNIQUE KEY,
  pwd TEXT
)
```
- video_info
```
CREATE TABLE video_info(
  id VARCHAR(64) PRIMARY KEY NOT NULL,
  author_id  INT UNSIGNED,
  name TEXT,
  display_ctime TEXT,
  creat_time DATETIME
)
```
- comments
```
CREATE TABLE comments(
  id VARCHAR(64) PRIMARY KEY NOT NULL,
  video_id VARCHAR(64),
  author_id UNSIGNED INT,
  content TEXT,
  time DATETIME
)
```
- sessions
```
CREATE TABLE sessions(
  session_id TINYTEXT PRIMARY KEY NOT NULL,
  login_nname VARCHAR(64)
)
```

## MySql
$ docker pull mysql
$ docker images
$ docker run --name ben-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=my-pass -d mysql:latest
$ docker ps

./mysql -u root -p my-pass
use video_server;
show tables;
describe users;
 
CREATE DATABASE video_server;
USE video_server;

CREATE TABLE users (
  id  INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  login_name VARCHAR(64)  UNIQUE KEY,
  pwd TEXT
)

