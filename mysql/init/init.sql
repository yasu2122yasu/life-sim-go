-- DBの作成
DROP DATABASE IF EXISTS test_database;
CREATE DATABASE test_database;
-- DBを切り替え
USE test_database;
-- テーブルを作成
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);

INSERT INTO users (name, email, password, created_at, updated_at)
VALUES ('John Doe', 'johndoe@example.com', 'secretpassword', '2023-11-29 13:37:27.864', '2023-11-29 13:37:27.864');
