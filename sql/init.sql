
create database pbl_app1;
use pbl_app1;

create table if not exists pbl_app1.user (
    id         int          not null auto_increment primary key,
    name       varchar(50)  not null,
    password   varchar(100) not null,
    gender     varchar(10)  not null,
    birthday   date         not null,
    state      boolean      not null default false,
    point      int          not null default 0,
    created_at datetime     not null,
    updated_at datetime,
    deleted_at datetime
)ENGINE=INNODB DEFAULT CHARSET=utf8;

create table if not exists pbl_app1.restaurant (
    id             int          not null auto_increment primary key,
    name           varchar(50)  not null,
    business_hours varchar(50)  not null,
    image          varchar(100) not null,
    created_at     datetime     not null,
    updated_at     datetime,
    deleted_at     datetime
)ENGINE=INNODB DEFAULT CHARSET=utf8;

create table if not exists pbl_app1.recognize (
    id            int       not null auto_increment primary key,
    restaurant_id int       not null,
    user_id       int       not null,
    created_at    datetime  not null,
    deleted_at    datetime,
    constraint fk_recognize_restaurant_id
        foreign key (restaurant_id)
            references pbl_app1.restaurant (id)
            on delete restrict on update restrict,
    constraint fk_recognize_user_id
        foreign key (user_id)
            references pbl_app1.user (id)
            on delete restrict on update restrict
)ENGINE=INNODB DEFAULT CHARSET=utf8;

create table if not exists pbl_app1.point (
    id            int         not null auto_increment primary key,
    restaurant_id int         not null,
    user_id       int         not null,
    transaction   varchar(10) not null,
    created_at    datetime    not null,
    constraint fk_point_restaurant_id
        foreign key (restaurant_id)
            references pbl_app1.restaurant (id)
            on delete restrict on update restrict,
    constraint fk_point_user_id
        foreign key (user_id)
            references pbl_app1.user (id)
            on delete restrict on update restrict
)ENGINE=INNODB DEFAULT CHARSET=utf8;

create table if not exists pbl_app1.post (
    id            int           not null auto_increment primary key,
    user_id       int           not null,
    restaurant_id int           not null,
    content       varchar(100)  not null,
    good          int           not null default 0,
    genre         varchar(50)   not null,
    comment       varchar(100),
    created_at    datetime      not null,
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
    state   boolean     not null default true,
    constraint fk_like_post_id
        foreign key (post_id)
            references pbl_app1.post (id)
            on delete restrict on update restrict,
    constraint fk_like_user_id
        foreign key (user_id)
            references pbl_app1.user (id)
            on delete restrict on update restrict
)ENGINE=INNODB DEFAULT CHARSET=utf8;

set time_zone = '+09:00';

# user table
insert into pbl_app1.user (name, password, created_at) values ('sato', 'pass', now());
insert into pbl_app1.user (name, password, created_at) values ('suzuki', 'su0902', now());
insert into pbl_app1.user (name, password, created_at) values ('watanabe', 'nabenabe', now());
insert into pbl_app1.user (name, password, created_at) values ('yamada', 'yamayamada', now());
insert into pbl_app1.user (name, password, created_at) values ('wakazono', 'yuta1126', now());

# restaurant table
insert into pbl_app1.restaurant (name, business_hours, image, created_at) values ('乃が美', '11:00-24:00', 'https://storage.googleapis.com/pbl-lookin-storage/restaurant1.jpeg', now());
insert into pbl_app1.restaurant (name, business_hours, image, created_at) values ('Pommeno-ki', '10:00-22:00', 'https://storage.googleapis.com/pbl-lookin-storage/restaurant2.jpeg', now());
insert into pbl_app1.restaurant (name, business_hours, image, created_at) values ('VELDE', '19:00-02:00', 'https://storage.googleapis.com/pbl-lookin-storage/restaurant3.jpeg', now());
insert into pbl_app1.restaurant (name, business_hours, image, created_at) values ('備後屋', '19:00-01:00', 'https://storage.googleapis.com/pbl-lookin-storage/restaurant4.jpeg', now());

# post table
insert into pbl_app1.post (user_id, restaurant_id, content, genre, comment, created_at) values (1, 1, 'https://storage.googleapis.com/pbl-lookin-storage/post1_mood.jpeg', 'mood', '居心地がいい', now());
insert into pbl_app1.post (user_id, restaurant_id, content, genre, comment, created_at) values (2, 1, 'https://storage.googleapis.com/pbl-lookin-storage/post7_food.jpeg', 'food', '焼き鳥うまい', now());
insert into pbl_app1.post (user_id, restaurant_id, content, genre, comment, created_at) values (4, 1, 'https://storage.googleapis.com/pbl-lookin-storage/post2_drink.jpeg', 'drink', 'ビール最高！！', now());
insert into pbl_app1.post (user_id, restaurant_id, content, genre, comment, created_at) values (2, 1, 'https://storage.googleapis.com/pbl-lookin-storage/post4_dessert.jpeg', 'dessert', 'あまい---', now());

insert into pbl_app1.post (user_id, restaurant_id, content, genre, comment, created_at) values (5, 2, 'https://storage.googleapis.com/pbl-lookin-storage/post8_food.jpeg', 'food', 'チーズ好きにはたまらない！', now());
insert into pbl_app1.post (user_id, restaurant_id, content, genre, comment, created_at) values (3, 2, 'https://storage.googleapis.com/pbl-lookin-storage/post11_dessert.jpeg', 'dessert', '抹茶がやばい！', now());
insert into pbl_app1.post (user_id, restaurant_id, content, genre, comment, created_at) values (4, 2, 'https://storage.googleapis.com/pbl-lookin-storage/post5_mood.jpeg', 'mood', 'お洒落', now());
insert into pbl_app1.post (user_id, restaurant_id, content, genre, comment, created_at) values (4, 2, 'https://storage.googleapis.com/pbl-lookin-storage/post3_drink.jpeg', 'drink', 'お酒LOVE', now());

insert into pbl_app1.post (user_id, restaurant_id, content, genre, comment, created_at) values (1, 3, 'https://storage.googleapis.com/pbl-lookin-storage/post9_food.jpeg', 'food', '淡々タンタン、ナポリタン！！', now());
insert into pbl_app1.post (user_id, restaurant_id, content, genre, comment, created_at) values (1, 3, 'https://storage.googleapis.com/pbl-lookin-storage/post14_dessert.jpeg', 'dessert', 'パフェの断面がめっちゃ綺麗', now());
insert into pbl_app1.post (user_id, restaurant_id, content, genre, comment, created_at) values (5, 3, 'https://storage.googleapis.com/pbl-lookin-storage/post6_mood.jpeg', 'mood', 'ずっとここにいたい', now());
insert into pbl_app1.post (user_id, restaurant_id, content, genre, comment, created_at) values (2, 3, 'https://storage.googleapis.com/pbl-lookin-storage/post10_drink.jpeg', 'drink', 'ゆっくりと、、、ワイン', now());

insert into pbl_app1.post (user_id, restaurant_id, content, genre, comment, created_at) values (2, 4, 'https://storage.googleapis.com/pbl-lookin-storage/post12_mood.jpeg', 'mood', 'オシャレ', now());
insert into pbl_app1.post (user_id, restaurant_id, content, genre, comment, created_at) values (3, 4, 'https://storage.googleapis.com/pbl-lookin-storage/post13_drink.jpeg', 'drink', '今日は、贅沢にシャンパン', now());

update pbl_app1.post set content = 'https://storage.googleapis.com/pbl-lookin-storage/post10_drink.jpeg' where id = 10;
update pbl_app1.post set content = 'https://storage.googleapis.com/pbl-lookin-storage/post13_drink.jpeg' where id = 13;

update pbl_app1.user set name='sasa', password='update_test', gender='woman', birthday=cast('19981006' as date), updated_at=now() where id=1;

# good table
insert into pbl_app1.good (post_id, user_id) values (1, 1);
insert into pbl_app1.good (post_id, user_id) values (1, 2);
insert into pbl_app1.good (post_id, user_id) values (1, 4);
insert into pbl_app1.good (post_id, user_id) values (1, 5);
insert into pbl_app1.good (post_id, user_id) values (2, 5);
insert into pbl_app1.good (post_id, user_id) values (2, 3);
insert into pbl_app1.good (post_id, user_id) values (2, 2);
insert into pbl_app1.good (post_id, user_id) values (2, 4);
insert into pbl_app1.good (post_id, user_id) values (3, 4);
insert into pbl_app1.good (post_id, user_id) values (3, 1);
insert into pbl_app1.good (post_id, user_id) values (3, 5);
insert into pbl_app1.good (post_id, user_id) values (3, 2);

# tableに入っているデータを全て消す(auto incrementの場合は1からデータが入る)
truncate table good;

# recognize table
insert into pbl_app1.recognize (restaurant_id, user_id, created_at) values (1, 1, now());
insert into pbl_app1.recognize (restaurant_id, user_id, created_at) values (1, 2, now());
insert into pbl_app1.recognize (restaurant_id, user_id, created_at) values (1, 4, now());
insert into pbl_app1.recognize (restaurant_id, user_id, created_at) values (2, 5, now());
insert into pbl_app1.recognize (restaurant_id, user_id, created_at) values (2, 3, now());
insert into pbl_app1.recognize (restaurant_id, user_id, created_at) values (2, 4, now());
insert into pbl_app1.recognize (restaurant_id, user_id, created_at) values (3, 1, now());
insert into pbl_app1.recognize (restaurant_id, user_id, created_at) values (3, 5, now());
insert into pbl_app1.recognize (restaurant_id, user_id, created_at) values (3, 2, now());


# insert into pbl_app1.user (name, password, created_at) values ('aaa', 'test', now());
# insert into pbl_app1.restaurant (name, business_hours, image, created_at) values ('rappi', '10:00-24:00', 'r.png', now());
# insert into pbl_app1.post (user_id, restaurant_id, image, genre, comment, created_at) values (1, 1, 'a.jpg', 'atmosphere', 'very delicious!', now());
#
# insert into pbl_app1.user (name, password, created_at) values ('bbb', 'pass', now());
# insert into pbl_app1.restaurant (name, business_hours, image, created_at) values ('hako-ya', '11:00-23:00', 'h.png', now());
# insert into pbl_app1.post (user_id, restaurant_id, image, genre, comment, created_at) values (2, 2, 'kaisen.jpg', 'food', 'very very delicious!', now());
#
# insert into pbl_app1.user (name, password, created_at) values ('ccc', 'ppap', now());
# insert into pbl_app1.restaurant (name, business_hours, image, created_at) values ('aburi', '11:00-21:30', 'ramen.png', now());
# insert into pbl_app1.post (user_id, restaurant_id, image, genre, comment, created_at, updated_at, deleted_at) values (3, 3, 'c.jpg', 'sweet', 'delicious!', now(), now(), now());

# insert into pbl_app1.user values (1, 'aaa', 'test', now(), now(), now());
# insert into pbl_app1.restaurant values (1, 'rappi', '10:00-24:00', 'r.png', now(), now(), now());
# insert into pbl_app1.post values (1, 1, 1, 'a.jpg', 11, 'atmosphere', 'very delicious!', now(), now(), now());
#
# insert into pbl_app1.user values (2, 'bbb', 'pass', now(), now(), now());
# insert into pbl_app1.restaurant values (2, 'hako-ya', '11:00-23:00', 'h.png', now(), now(), now());
# insert into pbl_app1.post values (2, 2, 2, 'kaisen.jpg', 22, 'food', 'very very delicious!', now(), now(), now());
#
# insert into pbl_app1.user values (3, 'ccc', 'ppap', now(), now(), null);
# insert into pbl_app1.restaurant values (3, 'rappi', '10:00-24:00', 'r.png', now(), now(), now());
# insert into pbl_app1.post values (3, 3, 3, 'a.jpg', 0, 'atmosphere', 'very delicious!', now(), now(), now());
