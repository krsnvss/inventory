CREATE TABLE hardware (
	id integer PRIMARY KEY AUTOINCREMENT,
	name string,
	serial_number string,
	manufacturer integer,
	model integer,
	type integer,
	production_date date,
	purchase_date date,
	barcode string
);

CREATE TABLE manufacturer (
	id integer PRIMARY KEY AUTOINCREMENT,
	name string
);

CREATE TABLE model (
	id integer PRIMARY KEY AUTOINCREMENT,
	manufacturer integer,
	name string
);

CREATE TABLE hardware_type (
	id integer PRIMARY KEY AUTOINCREMENT,
	name string
);

CREATE TABLE repair (
	id integer PRIMARY KEY AUTOINCREMENT,
	hardware integer,
	repair_date date,
	comment text
);

CREATE TABLE decomission (
	id integer PRIMARY KEY AUTOINCREMENT,
	hardware integer,
	decomission_date date,
	reason string
);

