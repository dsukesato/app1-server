
create database pbl_app1;
use pbl_app1;

create table if not exists pbl_app1.user (
    id         int          not null auto_increment primary key,
    name       varchar(50)  not null,
    password   varchar(100) not null,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime
)ENGINE=INNODB DEFAULT CHARSET=utf8;

create table if not exists pbl_app1.restaurant (
    id             int not null auto_increment primary key ,
    name           varchar(50)  not null,
    business_hours varchar(50)  not null,
    image          varchar(100) not null,
    created_at     datetime,
    updated_at     datetime,
    deleted_at     datetime
)ENGINE=INNODB DEFAULT CHARSET=utf8;

create table if not exists pbl_app1.recognize (
    id            int       not null auto_increment primary key,
    restaurant_id int       not null,
    created_at    datetime,
    updated_at    datetime,
    deleted_at    datetime,
    constraint fk_recognize_restaurant_id
        foreign key (restaurant_id)
            references pbl_app1.restaurant (id)
            on delete restrict on update restrict
)ENGINE=INNODB DEFAULT CHARSET=utf8;

create table if not exists pbl_app1.post (
    id            int not null auto_increment primary key,
    user_id       int not null,
    restaurant_id int not null,
    image         varchar(100),
    good          int default 0 not null,
    genre         varchar(50) not null,
    comment       varchar(100),
    created_at    datetime,
    updated_at    datetime,
    deleted_at    datetime,
    constraint fk_post_user_id
        foreign key (user_id)
            references pbl_app1.user (id)
            on delete restrict on update restrict,
    constraint fk_post_restaurant_id
        foreign key (restaurant_id)
            references pbl_app1.restaurant (id)
            on delete restrict on update restrict
)ENGINE=INNODB DEFAULT CHARSET=utf8;

create table if not exists pbl_app1.good (
    id      int         not null auto_increment primary key,
    post_id int         not null,
    user_id int         not null,
    constraint fk_like_post_id
        foreign key (post_id)
            references pbl_app1.post (id)
            on delete restrict on update restrict,
    constraint fk_like_user_id
        foreign key (user_id)
            references pbl_app1.user (id)
            on delete restrict on update restrict
)ENGINE=INNODB DEFAULT CHARSET=utf8;


insert into pbl_app1.user (name, password, created_at) values ('aaa', 'test', now());
insert into pbl_app1.restaurant (name, business_hours, image, created_at) values ('rappi', '10:00-24:00', 'r.png', now());
insert into pbl_app1.post (user_id, restaurant_id, image, genre, comment, created_at) values (1, 1, 'a.jpg', 'atmosphere', 'very delicious!', now());

insert into pbl_app1.user (name, password, created_at) values ('bbb', 'pass', now());
insert into pbl_app1.restaurant (name, business_hours, image, created_at) values ('hako-ya', '11:00-23:00', 'h.png', now());
insert into pbl_app1.post (user_id, restaurant_id, image, genre, comment, created_at) values (2, 2, 'kaisen.jpg', 'food', 'very very delicious!', now());

insert into pbl_app1.user (name, password, created_at) values ('ccc', 'ppap', now());
insert into pbl_app1.restaurant (name, business_hours, image, created_at) values ('aburi', '11:00-21:30', 'ramen.png', now());
insert into pbl_app1.post (user_id, restaurant_id, image, genre, comment, created_at, updated_at, deleted_at) values (3, 3, 'c.jpg', 'sweet', 'delicious!', now(), now(), now());


insert into pbl_app1.user values (1, 'aaa', 'test', now(), now(), now());
insert into pbl_app1.restaurant values (1, 'rappi', '10:00-24:00', 'r.png', now(), now(), now());
insert into pbl_app1.post values (1, 1, 1, 'a.jpg', 11, 'atmosphere', 'very delicious!', now(), now(), now());

insert into pbl_app1.user values (2, 'bbb', 'pass', now(), now(), now());
insert into pbl_app1.restaurant values (2, 'hako-ya', '11:00-23:00', 'h.png', now(), now(), now());
insert into pbl_app1.post values (2, 2, 2, 'kaisen.jpg', 22, 'food', 'very very delicious!', now(), now(), now());

insert into pbl_app1.user values (3, 'ccc', 'ppap', now(), now(), null);
insert into pbl_app1.restaurant values (3, 'rappi', '10:00-24:00', 'r.png', now(), now(), now());
insert into pbl_app1.post values (3, 3, 3, 'a.jpg', 0, 'atmosphere', 'very delicious!', now(), now(), now());
