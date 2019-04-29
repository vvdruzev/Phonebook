SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

DROP TABLE IF EXISTS `items`;
CREATE TABLE `PhoneBook` (
  `CountryCode` varchar(10) NOT NULL,
  `CountryName` varchar(45) DEFAULT NULL,
  `PhoneCode` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`CountryCode`),
  UNIQUE KEY `CountryCode_UNIQUE` (`CountryCode`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

