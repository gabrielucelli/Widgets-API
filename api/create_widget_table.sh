#!/bin/sh

DB_NAME=spa.sqlite
DB_TABLE=widgets

sqlite3 $DB_NAME << EOF


CREATE TABLE IF NOT EXISTS $DB_TABLE(
					id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
					name VARCHAR(100) NOT NULL,
					color VARCHAR(20) NOT NULL,
					price FLOAT NOT NULL,
					inventory INTEGER NOT NULL,
					melts BOOLEAN NOT NULL
				);


INSERT INTO $DB_TABLE(id, name, color, price, inventory, melts)
VALUES
(
1, 'Losenoid', 'blue', 9.99, 51, 'true'
);


INSERT INTO $DB_TABLE(id, name, color, price, inventory, melts)
VALUES
(
2, 'Rowlow', 'red', 4.00, 7, 'true'
);


INSERT INTO $DB_TABLE(id, name, color, price, inventory, melts)
VALUES
(
3, 'Printure', 'green', 5.55, 18, 'false'
);


INSERT INTO $DB_TABLE(id, name, color, price, inventory, melts)
VALUES
(
4, 'Claster', 'off-white', 12.56, 9, 'true'
);


INSERT INTO $DB_TABLE(id, name, color, price, inventory, melts)
VALUES
(
5, 'Pepelexa', 'purple', 0.99, 0, 'false'
);


INSERT INTO $DB_TABLE(id, name, color, price, inventory, melts)
VALUES
(
6, 'Dropellet', 'speckled', 16.00, 99, 'true'
);


INSERT INTO $DB_TABLE(id, name, color, price, inventory, melts)
VALUES
(
7, 'Jeebus', 'depends on the viewing angle', 18.11, 36, 'true'
);


INSERT INTO $DB_TABLE(id, name, color, price, inventory, melts)
VALUES
(
8, 'Nodile', 'red', 50.00, 20, 'true'
);


INSERT INTO $DB_TABLE(id, name, color, price, inventory, melts)
VALUES
(
9, 'Kaloobon', 'white', 8.00, 40, 'true'
);


INSERT INTO $DB_TABLE(id, name, color, price, inventory, melts)
VALUES
(
10, 'Bioyino', 'turtle-shell', 90.12, 3, 'true'
);


INSERT INTO $DB_TABLE(id, name, color, price, inventory, melts)
VALUES
(
11, 'Dizoolexa', 'magenta', 67.23, 976, 'false'
);


INSERT INTO $DB_TABLE(id, name, color, price, inventory, melts)
VALUES
(
12, 'Skippy', 'green', 0.01, 404, 'true'
);


INSERT INTO $DB_TABLE(id, name, color, price, inventory, melts)
VALUES
(
13, 'what', 'black', 50.00, 4, 'true'
);


INSERT INTO $DB_TABLE(id, name, color, price, inventory, melts)
VALUES
(
14, 'My id is actually 333', 'purple', 600.00, 900, 'true'
);


INSERT INTO $DB_TABLE(id, name, color, price, inventory, melts)
VALUES
(
15, 'Test', 'blue', 7.77, 13, 'true'
);


INSERT INTO $DB_TABLE(id, name, color, price, inventory, melts)
VALUES
(
16, 'Testing', 'red', 4.00, 7, 'true'
);


INSERT INTO $DB_TABLE(id, name, color, price, inventory, melts)
VALUES
(
17, 'Testing', 'test', 4.00, 7, 'true'
);


INSERT INTO $DB_TABLE(id, name, color, price, inventory, melts)
VALUES
(
18, 'Testing', 'test', 4.00, 7, 'true'
);


INSERT INTO $DB_TABLE(id, name, color, price, inventory, melts)
VALUES
(
19, 'Testing', 'test', 0, 7, 'true'
);


INSERT INTO $DB_TABLE(id, name, color, price, inventory, melts)
VALUES
(
20, 'Testing', 'test', 0, 7, 'true'
);


INSERT INTO $DB_TABLE(id, name, color, price, inventory, melts)
VALUES
(
21, 'Testing', 'test', 0, 7, 'true'
);


INSERT INTO $DB_TABLE(id, name, color, price, inventory, melts)
VALUES
(
22, 'Testing', 'test', 2.88, 7, 'true'
);


INSERT INTO $DB_TABLE(id, name, color, price, inventory, melts)
VALUES
(
23, 'Testing', 'test', 1.00, 7, 'true'
);


INSERT INTO $DB_TABLE(id, name, color, price, inventory, melts)
VALUES
(
24, 'Testing', 'test', 1.00, 7, 'true'
);


INSERT INTO $DB_TABLE(id, name, color, price, inventory, melts)
VALUES
(
25, 'Testing', 'test', 1.00, 0, 'true'
);


INSERT INTO $DB_TABLE(id, name, color, price, inventory, melts)
VALUES
(
26, 'Testing', 'test', 1.00, 0, 'true'
);


INSERT INTO $DB_TABLE(id, name, color, price, inventory, melts)
VALUES
(
27, 'Testing', 'test', 1.00, 52, 'true'
);


INSERT INTO $DB_TABLE(id, name, color, price, inventory, melts)
VALUES
(
28, 'Testing', 'test', 1.00, 52, 'true'
);


EOF