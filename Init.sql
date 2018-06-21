drop schema if exists ExpertHubAuth;
create schema ExpertHubAuth;
use ExpertHubAuth;

create table users(userid bigint auto_increment key, nickname varchar(32) not null unique, hashkey text not null, credit int default 0);

create table experts(
 expertid bigint auto_increment key,
 userid bigint, 
 nickname varchar(32) not null,
 gender varchar(32),
 email varchar(32),
 tel varchar(20),
 subgroup text,
 category text,
 avatar text,
 captcha text,
 accountstatus int default -1);

insert into users(nickname, hashkey) values("Alice","123");
insert into users(nickname, hashkey) values("Bob","111");
insert into users(nickname, hashkey) values("YangTui","666");