CREATE DATABASE IF NOT EXISTS belajarmysql;

USE belajarmysql;

CREATE TABLE IF NOT EXISTS person(
    id INT PRIMARY KEY not null auto_increment,
    name VARCHAR(200) not null
);