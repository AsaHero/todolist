CREATE TABLE Users 
(
    id int unsigned not null auto_increment,
    name varchar(255) not null,
    username varchar(255) not null,
    password_hash varchar(255) not null,
    primary key(id),
    unique(username)
);

CREATE TABLE ToDoList 
(
    id int unsigned not null auto_increment,
    title text not null,
    description text,
    primary key(id)
);

CREATE TABLE ToDoItem
(
    id int unsigned not null auto_increment,
    title text not null,
    description text,
    done boolean not null default false,
    primary key(id)
);

CREATE TABLE UsersList
(
    id int unsigned not null auto_increment,
    users_id int unsigned not null,
    list_id int unsigned not null,
    primary key(id),
    foreign key(users_id) references Users(id) on delete cascade,
    foreign key(list_id) references ToDoList(id) on delete cascade 
);

CREATE TABLE ItemList
(
    id int unsigned not null auto_increment,
    list_id int unsigned not null,
    item_id int unsigned not null,
    primary key(id),
    foreign key(list_id) references ToDoList(id) on delete cascade,
    foreign key(item_id) references ToDoItem(id) on delete cascade
);