-- create schema version table
create table schema_version (
    version text not null
);

create table keys (
    id integer primary key autoincrement,
    key text unique check (key <> '' and length(key) <= 20),
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    user_id integer unique,
    foreign key (user_id) references users(id)
);

create table users (
    id integer primary key autoincrement,
    name text unique CHECK (name <> '' and length(name) <= 15),
    hash text not null CHECK (hash <> ''),
    about TEXT not null DEFAULT '',
    picture text not null default '',
    is_admin boolean default false,
    key_id integer not null unique,
    foreign key (key_id) references keys(id)
);

create table settings (
    name text not null,
    css text not null default ''
);

create table boards (
    id integer primary key autoincrement,
    name text not null check ( name <> '' and length(name) < 120 ),
    description text
);

create table posts (
    id integer primary key autoincrement,
    subject text not null check ( length(subject) <= 120 ),
    content text not null check ( length(content) <= 50000 ),
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    topic_id integer references posts(id),
    board_id integer references boards(id),
    user_id integer references users(id)
);

create view topics as
    select
        t.id,
        t.user_id,
        u.name,
        t.board_id,
        t.subject,
        count(r.topic_id) replies,
        t.created_at,
        ifnull(r.updated_at, t.updated_at) updated_at
    from posts t
             left join posts r on r.topic_id = t.id
             left join users u on t.user_id = u.id
    where t.topic_id is null
    group by t.id order by r.updated_at desc;

create view boardStats as
    select
        b.id,
        b.name,
        b.description,
        count(p.id) as posts,
        sum(case when (p.topic_id is null and p.id is not null) then 1 else 0 end) topics
    from boards b left join posts p on b.id = p.board_id
    group by p.board_id;

create view postUsers as
    select
        p.id as post_id,
        p.subject,
        p.content,
        p.created_at,
        p.topic_id,
        u.id as user_id,
        u.name,
        u.picture
    from posts p
         left join users u on p.user_id = u.id;

insert into settings (name) values ('vpub');
insert into keys (key) values ('admin');