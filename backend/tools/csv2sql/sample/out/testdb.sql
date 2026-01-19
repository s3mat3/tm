--------------------
-- Tables for testdb
--------------------

--------------------
-- Table: table_info
-- Description: Table info
--------------------
DROP TABLE IF EXISTS table_info CASCADE;
CREATE TABLE IF NOT EXISTS table_info(
	id serial CONSTRAINT not null, -- table ID
	name varchar(50) CONSTRAINT not null unique, -- table name
	comment varchar(100), -- Description of this table
	PRIMARY KEY (id)
);
COMMENT ON TABLE table_info IS 'Table info';
COMMENT ON COLUMN table_info.id IS 'table ID';
COMMENT ON COLUMN table_info.name IS 'table name';
COMMENT ON COLUMN table_info.comment IS 'Description of this table';
--------------------
-- Table: column_info
-- Description: Column info
--------------------
DROP TABLE IF EXISTS column_info CASCADE;
CREATE TABLE IF NOT EXISTS column_info(
	id int CONSTRAINT not null, -- Column ID
	table_id int CONSTRAINT not null, -- Table info ID
	name varchar(50) CONSTRAINT not null, -- Column name
	display varchar(50), -- Display name for this column
	pk char, -- Is this column a primary key? T | t | “”
	fk varchar(50), -- table.column
	type varchar(50) CONSTRAINT not null, -- Data type of this column
	constraint varchar(100), -- This column constraint
	comment varchar(100), -- Description for this column
	PRIMARY KEY (id,table_id),
	FOREIGN KEY (table_id) REFERENCES table_info (id) ON UPDATE CASCADE ON DELETE CASCADE
);
COMMENT ON TABLE column_info IS 'Column info';
COMMENT ON COLUMN column_info.table_id IS 'Table info ID';
COMMENT ON COLUMN column_info.pk IS 'Is this column a primary key? T | t | “”';
COMMENT ON COLUMN column_info.fk IS 'table.column';
COMMENT ON COLUMN column_info.constraint IS 'This column constraint';
COMMENT ON COLUMN column_info.id IS 'Column ID';
COMMENT ON COLUMN column_info.name IS 'Column name';
COMMENT ON COLUMN column_info.display IS 'Display name for this column';
COMMENT ON COLUMN column_info.type IS 'Data type of this column';
COMMENT ON COLUMN column_info.comment IS 'Description for this column';
--------------------
-- Table: ui_field_type
-- Description: UI field type list
--------------------
DROP TABLE IF EXISTS ui_field_type CASCADE;
CREATE TABLE IF NOT EXISTS ui_field_type(
	id serial CONSTRAINT not null, -- ID
	name Varchar(20) CONSTRAINT not null unique, -- Type name in UI (dropdown
	PRIMARY KEY (id)
);
COMMENT ON TABLE ui_field_type IS 'UI field type list';
COMMENT ON COLUMN ui_field_type.id IS 'ID';
COMMENT ON COLUMN ui_field_type.name IS 'Type name in UI (dropdown';
--------------------
-- Table: ui_info
-- Description: Column info for UI extentions
--------------------
DROP TABLE IF EXISTS ui_info CASCADE;
CREATE TABLE IF NOT EXISTS ui_info(
	column_id int CONSTRAINT not null, -- Column info table ID
	table_id int CONSTRAINT not null, -- Table info tble ID
	type int CONSTRAINT not null, -- Column type in UI
	default varchar(100), -- The default value of the string representation
	dropdown varchar(100), -- Drop-down list source table name
	requier bool, -- Is requier this column
	readonly bool, -- Is read only this column
	hidden bool, -- Is hidden this column
	PRIMARY KEY (column_id,table_id),
	FOREIGN KEY (column_id) REFERENCES column_info (id) ON UPDATE CASCADE ON DELETE CASCADE,
	FOREIGN KEY (table_id) REFERENCES table_info (id) ON UPDATE CASCADE ON DELETE CASCADE,
	FOREIGN KEY (type) REFERENCES ui_field_type (id) ON UPDATE CASCADE ON DELETE CASCADE
);
COMMENT ON TABLE ui_info IS 'Column info for UI extentions';
COMMENT ON COLUMN ui_info.default IS 'The default value of the string representation';
COMMENT ON COLUMN ui_info.dropdown IS 'Drop-down list source table name';
COMMENT ON COLUMN ui_info.requier IS 'Is requier this column';
COMMENT ON COLUMN ui_info.readonly IS 'Is read only this column';
COMMENT ON COLUMN ui_info.hidden IS 'Is hidden this column';
COMMENT ON COLUMN ui_info.column_id IS 'Column info table ID';
COMMENT ON COLUMN ui_info.table_id IS 'Table info tble ID';
COMMENT ON COLUMN ui_info.type IS 'Column type in UI';
