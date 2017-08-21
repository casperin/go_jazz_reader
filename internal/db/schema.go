package db

const (
	dropAppTable = `drop table app;`
	appTable     = `create table app (
		id bigserial primary key,
		last_items_update timestamp NOT NULL,
		last_email_update timestamp NOT NULL
	);`

	dropFeedTable = `drop table feeds;`
	feedTable     = `create table feeds (
		id bigserial primary key,
		url varchar(200) NOT NULL UNIQUE,
		link varchar(200) NOT NULL,
		title varchar(200) NOT NULL,
		description text NULL,
		error varchar(200) NULL
	);`

	dropItemTable = `drop table posts;`
	itemTable     = `create table posts (
		id bigserial primary key,
		status varchar(20) NOT NULL default 'new',
		saved boolean NOT NULL default 'false',
		feed_id bigserial NOT NULL,
		guid varchar(200) NOT NULL UNIQUE,
		author varchar(200) NOT NULL,
		link varchar(200) NOT NULL,
		title varchar(200) NOT NULL,
		feed_title varchar(200) NOT NULL,
		description text NOT NULL,
		content text NOT NULL,
		published timestamp NOT NULL,
		created_at timestamp NOT NULL,
		FOREIGN KEY (feed_id) REFERENCES feeds(id)
	);`

	dropUrlTable = `drop table urls;`
	urlTable     = `create table urls (
		id bigserial primary key,
		url varchar(200) NOT NULL,
		title varchar(200) NOT NULL default '',
		created_at timestamp NOT NULL
	);`
)
