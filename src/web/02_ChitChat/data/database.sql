drop table if exists posts;
drop table if exists threads;
drop table if exists sessions;
drop table if exists users;

--- 用户信息表
create table users (
    id serial primary key,
    uuid varchar(64) not null unique,
    name varchar(255),
    email varchar(255) not null unique,
    password varchar(255) not null,
    created_at timestamp not null
);

--- 用户登陆会话
create table sessions (
    id serial primary key,
    uuid varchar(64) not null unique,
    email varchar(255),
    user_id integer references users(id),
    password varchar(255) not null,
    created_at timestamp not null
);

--- 帖子
create table threads (
    id serial primary key,
    uuid varchar(64) not null unique,
    topic text,
    user_id integer references users(id),
    created_at timestamp not null
);

--- 帖子的回复
create table posts (
    id serial primary key,
    uuid varchar(64) not null unique,
    body text,
    user_id integer references users(id),
    thread_id integer references threads(id),
    created_at timestamp not null
);