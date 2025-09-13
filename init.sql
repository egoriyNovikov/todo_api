-- Создание базы данных и пользователя (если не существуют)
-- CREATE DATABASE IF NOT EXISTS todo_db;
-- CREATE USER IF NOT EXISTS 'todo_user'@'%' IDENTIFIED BY 'todo_password';
-- GRANT ALL PRIVILEGES ON todo_db.* TO 'todo_user'@'%';

-- Создание таблицы для задач
CREATE TABLE IF NOT EXISTS todos (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Создание индексов для оптимизации
CREATE INDEX IF NOT EXISTS idx_todos_completed ON todos(completed);
CREATE INDEX IF NOT EXISTS idx_todos_created_at ON todos(created_at);

-- Вставка тестовых данных
INSERT INTO todos (title, description, completed) VALUES
('Изучить Docker', 'Изучить основы Docker и Docker Compose', false),
('Создать API', 'Создать REST API для управления задачами', false),
('Настроить PostgreSQL', 'Настроить подключение к базе данных', true)
ON CONFLICT DO NOTHING;
