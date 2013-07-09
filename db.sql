create extension postgis;

create table groups (
	id serial primary key,
	name text
);

create table users (
	id serial primary key,
	login text,
	hash text,
	gid int references groups (id),
	name text
);

create table hold_types (
	id serial primary key,
	name text
);
create index hold_type_idx on hold_types (id);

create table woody_types ( 
	id serial primary key,
	name text
);
create index woody_type_idx on woody_types (id);

create table woodies (
	id serial primary key,
	uid int references users (id), 
	gid int references groups (id),
	name text,
	geom geometry
);
create index woodies_type_idx on woodies (id);
create index woodies_geom_idx on woodies using gist(geom);

create table holds ( 
	id serial primary key,
	type int references hold_types (id),
	woody int references woodies (id),
	name text,
	geom geometry
);
create index holds_idx on holds (id);
create index holds_geom_idx on holds using gist(geom);
