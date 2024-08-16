-- Cria o banco de dados
CREATE DATABASE IF NOT EXISTS testdb;

-- Usa o banco de dados criado
USE testdb;

-- Cria a tabela de usuários
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE
);

-- Insere alguns usuários de exemplo
INSERT INTO users (name, email) VALUES ('Alice Johnson', 'alice@example.com');
INSERT INTO users (name, email) VALUES ('Bob Smith', 'bob@example.com');
INSERT INTO users (name, email) VALUES ('Charlie Brown', 'charlie@example.com');
INSERT INTO users (name, email) VALUES ('Diana Prince', 'diana@example.com');
INSERT INTO users (name, email) VALUES ('Eve Adams', 'eve@example.com');


CREATE TABLE IF NOT EXISTS challenge_day (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

INSERT INTO challenge_day (name) VALUES ('Higor Diego', 'higor_challenge@gmail.com');