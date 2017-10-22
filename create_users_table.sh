#!/bin/sh

DB_NAME=spa.sqlite
DB_TABLE=users

sqlite3 $DB_NAME << EOF


CREATE TABLE IF NOT EXISTS $DB_TABLE (
					id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
					name VARCHAR(100) NOT NULL,
					gravatar VARCHAR(200)
				);


INSERT INTO $DB_TABLE(name, id, gravatar)
VALUES
(
'Colin', 1, 'http://www.gravatar.com/avatar/a51972ea936bc3b841350caef34ea47e?s=64&d=monsterid'
);

INSERT INTO $DB_TABLE(name, id, gravatar)
VALUES
(
'Kyle', 2, 'http://www.gravatar.com/avatar/432f3e353c689fc37af86ae861d934f9?s=64&d=monsterid'
);

INSERT INTO $DB_TABLE(name, id, gravatar)
VALUES
(
'Thomas', 3, 'http://www.gravatar.com/avatar/48009c6a27d25ef0ea03f985d1f186b0?s=64&d=monsterid'
);

INSERT INTO $DB_TABLE(name, id, gravatar)
VALUES
(
'James', 4, 'http://www.gravatar.com/avatar/9372f138140c8578c82bbc77b2eca602?s=64&d=monsterid'
);

EOF