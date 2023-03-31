CREATE DATABASE `platzidemo` 
use  `platzidemo` 
CREATE TABLE platzidemo.notas (
	id BIGINT auto_increment NOT NULL,
	title varchar(100) NULL,
	description varchar(100) NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
	CONSTRAINT notas_PK PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8
COLLATE=utf8_general_ci;


