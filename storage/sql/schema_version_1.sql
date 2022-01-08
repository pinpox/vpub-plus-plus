-- create schema version table
create table schema_version (
    version text not null
);

-- create users table
create table users
(
    name text primary key CHECK (name <> '' and length(name) <= 15),
    hash text not null CHECK (hash <> ''),
    about TEXT not null DEFAULT ''
);

-- create posts table
create table posts
(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    author TEXT references users(name) NOT NULL,
    title TEXT NOT NULL CHECK (title <> '' and length(title) <= 120),
    content TEXT NOT NULL CHECK (content <> '' and length(title) <= 50000),
    topic TEXT,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);

-- create replies table
create table replies
(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    author TEXT references users(name) NOT NULL,
    content TEXT NOT NULL CHECK (content <> ''),
    post_id int references posts(id) NOT NULL,
    parent_id int references replies(id),
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    foreign key (post_id) references posts(id),
    foreign key (parent_id) references replies(id)
);

-- create notification table
create table notifications
(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    author text references users(name),
    reply_id int references replies(id),
    foreign key (author) references users(name) on delete cascade,
    foreign key (reply_id) references replies(id) on delete cascade
);

-- indices
CREATE INDEX idx_replies_parent_id ON replies(parent_id);
CREATE INDEX idx_replies_post_id ON replies(post_id);

-- New data model

-- A forum has topics
create table topics (
    id integer primary key autoincrement,
    name text,
    description text
);

-- Could we not have threads? Instead have posts that don't have a thread
create table tposts (
    id integer primary key autoincrement,
    author text,
    subject text not null check ( length(subject) < 120 ),
    content text not null check ( length(content) < 50000 ),
    thread integer,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    topic integer,
    foreign key (topic) references topics(id),
    foreign key (thread) references tposts(id),
    foreign key (author) references users(name)
)