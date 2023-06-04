create table users 
(
    id serial not null unique,
    first_name varchar(255) not null,
    last_name varchar(255) not null,
    email varchar(255) not null,
    password varchar(255) not null
);

create table todo_lists
(
    id serial not null unique,
    title varchar(255) not null,
    description varchar(255) not null
);

create table users_lists
(
    id serial not null unique,
    user_id int references users (id) on delete cascade not null,
    list_id int references todo_lists (id) on delete cascade not null    
);

create table todo_items
(
    id serial not null unique,
    title varchar(255) not null,
    description varchar(255),
    done boolean not null default false
);

create table lists_items
(
    id serial not null unique,
    item_id int references todo_items (id) on delete cascade not null,
    list_id int references todo_lists (id) on delete cascade not null    
);