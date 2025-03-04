-- ALTER SYSTEM SET wal_level = logical;
-- CREATE PUBLICATION debezium_pub FOR ALL TABLES;
-- -- CREATE PUBLICATION debezium_pub FOR TABLE students, student_notifications;
-- SELECT * FROM pg_create_logical_replication_slot('debezium_slot', 'pgoutput');

-- 🔥 Bật chế độ WAL Level Logical để hỗ trợ Debezium
ALTER SYSTEM SET wal_level = logical;

-- 🚀 Tạo publication cho toàn bộ bảng (hoặc chỉ các bảng cụ thể)
DROP PUBLICATION IF EXISTS debezium_pub;
CREATE PUBLICATION debezium_pub FOR ALL TABLES;

-- 🛠️ Tạo replication slot (nếu đã có thì xóa trước)
SELECT pg_drop_replication_slot('debezium_slot') WHERE EXISTS 
    (SELECT 1 FROM pg_replication_slots WHERE slot_name = 'debezium_slot');

SELECT * FROM pg_create_logical_replication_slot('debezium_slot', 'pgoutput');

ALTER TABLE students REPLICA IDENTITY FULL;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS CITEXT;




-- DO $$ 
-- BEGIN 
--    IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'student_management') THEN
--       CREATE DATABASE student_management;
--    END IF;
-- END $$;

-- \connect student_management;

-- CREATE TABLE IF NOT EXISTS genders (
--     id SERIAL PRIMARY KEY,
--     name VARCHAR(10) NOT NULL
-- );

-- CREATE TABLE IF NOT EXISTS faculties (
--     id SERIAL PRIMARY KEY,
--     name VARCHAR(255) NOT NULL
-- );

-- CREATE TABLE IF NOT EXISTS statuses (
--     id SERIAL PRIMARY KEY,
--     name VARCHAR(50) NOT NULL
-- );

-- CREATE TABLE IF NOT EXISTS students (
--     id INT PRIMARY KEY,
--     full_name VARCHAR(255) NOT NULL,
--     birth_date DATE NOT NULL,
--     gender_id INT NOT NULL,
--     faculty_id INT NOT NULL,
--     course VARCHAR(100),
--     program VARCHAR(255),
--     address TEXT,
--     email VARCHAR(255) UNIQUE NOT NULL,
--     phone VARCHAR(20) UNIQUE NOT NULL,
--     status_id INT NOT NULL,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
-- );

-- ALTER TABLE students
-- ADD CONSTRAINT fk_gender FOREIGN KEY (gender_id) REFERENCES genders(id),
-- ADD CONSTRAINT fk_faculty FOREIGN KEY (faculty_id) REFERENCES faculties(id),
-- ADD CONSTRAINT fk_status FOREIGN KEY (status_id) REFERENCES statuses(id);

-- -- Insert data
-- INSERT INTO genders (name) VALUES ('Male'), ('Female'), ('Other') ON CONFLICT DO NOTHING;

-- INSERT INTO faculties (name) VALUES ('Law'), ('Business English'), ('Japanese'), ('French') ON CONFLICT DO NOTHING;

-- INSERT INTO statuses (name) VALUES ('Studying'), ('Graduated'), ('Dropped Out'), ('Paused') ON CONFLICT DO NOTHING;

-- INSERT INTO students (id, full_name, birth_date,  gender_id, faculty_id, course, program, address, email, phone, status_id) 
-- VALUES 
-- (22127180, 'John Doe', '1995-05-15', 1, 1, 'Law 101', 'Law', '123 Main St', 'john.doe@example.com', '123-456-7890', 1) 
-- ON CONFLICT DO NOTHING;
