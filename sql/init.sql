
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
    user_id       int       not null,
    created_at    datetime,
    updated_at    datetime,
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

set time_zone = '+09:00';

# user table
insert into pbl_app1.user (name, password, created_at) values ('sato', 'pass', now());
insert into pbl_app1.user (name, password, created_at) values ('suzuki', 'su0902', now());
insert into pbl_app1.user (name, password, created_at) values ('watanabe', 'nabenabe', now());
insert into pbl_app1.user (name, password, created_at) values ('yamada', 'yamayamada', now());
insert into pbl_app1.user (name, password, created_at) values ('wakazono', 'yuta1126', now());

# restaurant table
insert into pbl_app1.restaurant (name, business_hours, image, created_at) values ('乃が美', '11:00-24:00', 'https://storage.cloud.google.com/pbl-lookin/image/jpeg/restaurant/restaurant1.jpeg?hl=ja', now());
insert into pbl_app1.restaurant (name, business_hours, image, created_at) values ('Pommeno-ki', '10:00-22:00', 'https://storage.cloud.google.com/pbl-lookin/image/jpeg/restaurant/restaurant2.jpg?hl=ja', now());
insert into pbl_app1.restaurant (name, business_hours, image, created_at) values ('VELDE', '19:00-02:00', 'https://storage.cloud.google.com/pbl-lookin/image/jpeg/restaurant/restaurant3.jpeg?hl=ja', now());

# post table
insert into pbl_app1.post (user_id, restaurant_id, image, genre, comment, created_at) values (1, 1, 'https://storage.cloud.google.com/pbl-lookin/image/jpeg/post/mood/mood1.jpg?hl=ja', 'mood', '居心地がいい', now());
insert into pbl_app1.post (user_id, restaurant_id, image, genre, comment, created_at) values (2, 1, 'https://storage.cloud.google.com/pbl-lookin/image/jpeg/post/food/food1.jpeg?hl=ja', 'food', '焼き鳥うまい', now());
insert into pbl_app1.post (user_id, restaurant_id, image, genre, comment, created_at) values (4, 1, 'https://storage.cloud.google.com/pbl-lookin/image/jpeg/post/drink/drink2.jpeg?hl=ja', 'drink', 'ビール最高！！', now());
insert into pbl_app1.post (user_id, restaurant_id, image, genre, comment, created_at) values (2, 1, 'https://storage.cloud.google.com/pbl-lookin/image/jpeg/post/dessert/dessert1.jpeg?hl=ja', 'dessert', 'あまい---', now());

insert into pbl_app1.post (user_id, restaurant_id, image, genre, comment, created_at) values (5, 2, 'https://storage.cloud.google.com/pbl-lookin/image/jpeg/post/food/food2.jpg?hl=ja', 'food', 'チーズ好きにはたまらない！', now());
insert into pbl_app1.post (user_id, restaurant_id, image, genre, comment, created_at) values (3, 2, 'https://storage.cloud.google.com/pbl-lookin/image/jpeg/post/dessert/dessert3.jpeg?hl=ja', 'dessert', '抹茶がやばい！', now());
insert into pbl_app1.post (user_id, restaurant_id, image, genre, comment, created_at) values (4, 2, 'https://storage.cloud.google.com/pbl-lookin/image/jpeg/post/mood/mood2.jpeg?hl=ja', 'mood', 'お洒落', now());
insert into pbl_app1.post (user_id, restaurant_id, image, genre, comment, created_at) values (4, 2, 'https://storage.cloud.google.com/pbl-lookin/image/jpeg/post/drink/drink1.jpg?hl=ja', 'drink', 'お酒LOVE', now());

insert into pbl_app1.post (user_id, restaurant_id, image, genre, comment, created_at) values (1, 3, 'https://storage.cloud.google.com/pbl-lookin/image/jpeg/post/food/food3.jpeg?hl=ja', 'food', '淡々タンタン、ナポリタン！！', now());
insert into pbl_app1.post (user_id, restaurant_id, image, genre, comment, created_at) values (1, 3, 'https://storage.cloud.google.com/pbl-lookin/image/jpeg/post/food/food3.jpeg?hl=ja', 'dessert', 'パフェの断面がめっちゃ綺麗', now());
insert into pbl_app1.post (user_id, restaurant_id, image, genre, comment, created_at) values (5, 3, 'https://storage.cloud.google.com/pbl-lookin/image/jpeg/post/mood/mood3.jpeg?hl=ja', 'mood', 'ずっとここにいたい', now());
insert into pbl_app1.post (user_id, restaurant_id, image, genre, comment, created_at) values (2, 3, 'https://storage.cloud.google.com/pbl-lookin/image/jpeg/post/drink/drink3.jpg?hl=ja', 'drink', 'ゆっくりと、、、ワイン', now());

# good table
insert into pbl_app1.good (post_id, user_id) values (1, 1);
insert into pbl_app1.good (post_id, user_id) values (1, 2);
insert into pbl_app1.good (post_id, user_id) values (1, 4);
insert into pbl_app1.good (post_id, user_id) values (1, 2);
insert into pbl_app1.good (post_id, user_id) values (2, 5);
insert into pbl_app1.good (post_id, user_id) values (2, 3);
insert into pbl_app1.good (post_id, user_id) values (2, 4);
insert into pbl_app1.good (post_id, user_id) values (2, 4);
insert into pbl_app1.good (post_id, user_id) values (3, 1);
insert into pbl_app1.good (post_id, user_id) values (3, 1);
insert into pbl_app1.good (post_id, user_id) values (3, 5);
insert into pbl_app1.good (post_id, user_id) values (3, 2);

#recognize table
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
