create database if not exists cmdb default charset utf8mb4;

drop table user;

create table if not exists user(
    id          bigint primary key auto_increment,
    staff_id    varchar (32) not null default '',
    name        varchar(64) not null default '',
    nickname    varchar (64) not null default '',
    password    varchar(1024) not null default '',
    gender      int not null default 0 comment '0: 女, 1: 男',
    tel         varchar (32) not null default '',
    email       varchar (64) not null default '',
    addr        varchar (128) not null default '',
    department  varchar (128) not null default '',
    status      int not null default 0 comment '0: 正常, 1: 锁定, 2: 离职',
    created_at  datetime not null,
    updated_at  datetime not null,
    deleted_at  datetime
)engine = innodb default charset utf8mb4;

select id, staff_id, name, nickname, password, gender, tel, email, addr, department, status, created_at, updated_at, deleted_at from user;

insert into user(staff_id, name, nickname, password, gender, tel, email, addr, department, status, created_at, updated_at)
values ('L0001', 'lhq', 'LHQ', '$2a$10$TKP0wGzN41tmCyyf0Dr98Op6CCOFNw/3KpAqcW45zGopZvRmbxA3C', '1', '15209880622', '1029806879@qq.com', '之江家园二区', 'sb', '0', '2021/08/23', '2021/08/23');

insert into user(staff_id, name, nickname, password, gender, tel, email, addr, department, status, created_at, updated_at)
values ('L0002', 'hab', 'HAB', '$2a$10$TKP0wGzN41tmCyyf0Dr98Op6CCOFNw/3KpAqcW45zGopZvRmbxA3C', '2', '15868783887', '1029806879@qq.com', '之江家园二区', 'sb', '0', '2021/08/23', '2021/08/23');
