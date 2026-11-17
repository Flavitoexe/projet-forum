CREATE DATABASE IF NOT EXISTS bluedit;
USE bluedit;

CREATE TABLE users (
    id_user INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(64) NOT NULL UNIQUE,
    email VARCHAR(64) NOT NULL UNIQUE,
    password VARCHAR(256) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    role VARCHAR(8) NOT NULL DEFAULT "user",
    is_banned BOOLEAN DEFAULT false
);

CREATE TABLE categories (
    id_category INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(64) NOT NULL UNIQUE,
    description TEXT
);

CREATE TABLE threads (
    id_thread INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(256) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status INT NOT NULL DEFAULT 1,
    fk_id_category INT NOT NULL,
    fk_id_user INT DEFAULT NULL,

    FOREIGN KEY (fk_id_category) REFERENCES categories(id_category) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (fk_id_user) REFERENCES users(id_user) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE TABLE messages (
    id_message INT AUTO_INCREMENT PRIMARY KEY,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    fk_id_user INT DEFAULT NULL,
    fk_id_thread INT NOT NULL,

    FOREIGN KEY (fk_id_user) REFERENCES users(id_user) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (fk_id_thread) REFERENCES threads(id_thread) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE thread_likes (
    id_thread_like INT AUTO_INCREMENT PRIMARY KEY,
    fk_id_thread INT NOT NULL,
    fk_id_user INT NOT NULL,

    FOREIGN KEY (fk_id_user) REFERENCES users(id_user) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (fk_id_thread) REFERENCES threads(id_thread) ON UPDATE CASCADE ON DELETE CASCADE,

    CONSTRAINT unique_user_thread_like UNIQUE (fk_id_user, fk_id_thread)
);

CREATE TABLE message_likes (
    id_message_like INT AUTO_INCREMENT PRIMARY KEY,
    fk_id_message INT NOT NULL,
    fk_id_user INT NOT NULL,

    FOREIGN KEY (fk_id_user) REFERENCES users(id_user) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (fk_id_message) REFERENCES messages(id_message) ON UPDATE CASCADE ON DELETE CASCADE,

    CONSTRAINT unique_user_message_like UNIQUE (fk_id_user, fk_id_message)
);