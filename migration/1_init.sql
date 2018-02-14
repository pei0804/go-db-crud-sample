-- +migrate Up
CREATE TABLE [articles]
(
	[title] text NOT NULL,
	[body] text NOT NULL
);


CREATE TABLE [timelines]
(
	[title] text NOT NULL
);

-- +migrate Down
DROP TABLE [articles];
DROP TABLE [timelines];

