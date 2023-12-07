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

CREATE TABLE turns (
    id INT AUTO_INCREMENT PRIMARY KEY,
    week VARCHAR(255) NOT NULL,
);

CREATE TABLE characters (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE abilities (
    id INT AUTO_INCREMENT PRIMARY KEY,
    knowledge TINYINT UNSIGNED NOT NULL,
    strength TINYINT UNSIGNED NOT NULL,
    social TINYINT UNSIGNED NOT NULL,
    stress TINYINT UNSIGNED NOT NULL,
    decision TINYINT UNSIGNED NOT NULL,
    luck TINYINT UNSIGNED NOT NULL
);


CREATE TABLE character_abilities (
    character_id INT NOT NULL,
    ability_id INT NOT NULL,
    FOREIGN KEY (character_id) REFERENCES characters(id),
    FOREIGN KEY (ability_id) REFERENCES abilities(id),
    PRIMARY KEY (character_id, ability_id)
);

CREATE TABLE events (
    id INT AUTO_INCREMENT PRIMARY KEY,
    description VARCHAR(255) NOT NULL,
);

CREATE TABLE event_details (
    event_id INT NOT NULL,
    FOREIGN KEY (event_id) REFERENCES events(id),
    rate TINYINT UNSIGNED NOT NULL,
)

-- テストデータを挿入
INSERT INTO users (name, email, password, created_at, updated_at)
VALUES ('John Doe', 'johndoe@example.com', 'secretpassword', '2023-11-29 13:37:27.864', '2023-11-29 13:37:27.864');
