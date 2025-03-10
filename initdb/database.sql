-- Enable WAL Level Logical to support Debezium
ALTER SYSTEM SET wal_level = logical;

-- Create publication for all tables (or specific tables)
DROP PUBLICATION IF EXISTS debezium_pub;
CREATE PUBLICATION debezium_pub FOR ALL TABLES;

-- Create replication slot (drop if it already exists)
SELECT pg_drop_replication_slot('debezium_slot') WHERE EXISTS 
    (SELECT 1 FROM pg_replication_slots WHERE slot_name = 'debezium_slot');

SELECT * FROM pg_create_logical_replication_slot('debezium_slot', 'pgoutput');

ALTER TABLE students REPLICA IDENTITY FULL;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS CITEXT;
