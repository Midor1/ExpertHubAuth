drop schema if exists ExpertHubAuth;
create schema ExpertHubAuth;
use ExpertHubAuth;

create table users(userid bigint auto_increment key, nickname varchar(32) not null, hashkey text not null);

insert into users(nickname, hashkey) values("Alice","123");
insert into users(nickname, hashkey) values("Bob","111");
insert into users(nickname, hashkey) values("YangTui","666");