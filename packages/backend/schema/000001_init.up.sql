create table users (
  id serial not null unique,
  email varchar(255) not null unique,
  password_hash varchar(255) not null
);
create table todo_lists (
  id serial not null unique,
  title varchar(255) not null,
  user_id int references users(id) on delete cascade not null
);
create table todo_items (
  id serial not null unique,
  title varchar(255) not null,
  done boolean not null default false,
  list_id int references todo_lists(id) on delete cascade not null
);