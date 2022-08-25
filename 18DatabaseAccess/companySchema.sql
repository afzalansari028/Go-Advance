CREATE DATABASE 'company'

USE 'company'

CREATE TABLE IF NOT EXISTS 'employees'(
    'id' int unsined NOT NULL AUTO_INCREMENT,
    'name' VARCHAR(20)
    'city' VARCHAR(20) DEFAULT "fredericton",
    'dept' VARCHAR(20) DEFAULT "onboarding",
    PRIMARY KEY ('id')
)ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb4;