CREATE TABLE hardware (
	id integer PRIMARY KEY AUTOINCREMENT,
	name string,
	serial_number string,
	manufacturer integer,
	model integer,
	type integer,
	production_date date,
	purchase_date date,
	barcode string,
	cpu_name string,
	cpu_max_clock float,
	cpu_cores integer,
	ram integer,
	disk integer,
	user_name string
);

CREATE TABLE manufacturer (
	id integer PRIMARY KEY AUTOINCREMENT,
	name string,
	logo string
);

CREATE TABLE model (
	id integer PRIMARY KEY AUTOINCREMENT,
	manufacturer integer,
	name string,
	type integer
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

CREATE TABLE user (
	id integer PRIMARY KEY AUTOINCREMENT,
	login string,
	password string,
	first_name string,
	last_name string,
	group_id integer
);

CREATE TABLE group (
	id integer PRIMARY KEY AUTOINCREMENT,
	name string
);


CREATE VIEW hardware_full
AS
SELECT
	hardware.id as id,
	hardware.name as name,
	hardware.user_name as user_name,
	manufacturer.name as manufacturer,
	model.name as model,
	hardware_type.name as type,
	hardware.production_date as production_date,
	hardware.purchase_date as purchase_date,
	hardware.barcode as barcode,
	hardware.cpu_name as cpu_name,
	hardware.cpu_max_clock as cpu_max_clock,
	hardware.cpu_cores as cpu_cores,
	hardware.ram as ram,
	hardware.disk as disk
FROM
	hardware
INNER JOIN manufacturer ON manufacturer.id = hardware.manufacturer
INNER JOIN model ON model.id = hardware.model
INNER JOIN hardware_type ON hardware_type.id = hardware.type;

CREATE VIEW hardware_short
AS
SELECT
	hardware.id as id,
	hardware.barcode as barcode,
	manufacturer.name as manufacturer,
	model.name as model,
	hardware.user_name as user_name
FROM
	hardware
INNER JOIN manufacturer ON manufacturer.id = hardware.manufacturer
INNER JOIN model ON model.id = hardware.model;

CREATE UNIQUE INDEX "hardwaretypename" ON "hardware_type" (
	"name"	ASC
);

CREATE UNIQUE INDEX "manufacturername" ON "manufacturer" (
	"name"	ASC
);

CREATE UNIQUE INDEX "modelname" ON "model" (
	"name"	ASC
);

CREATE UNIQUE INDEX "username" ON "user" (
	"login"	ASC
);