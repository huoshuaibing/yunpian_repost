SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

CREATE DATABASE IF NOT EXISTS `yunpian_db` DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci;
USE `yunpian_db`;

DROP TABLE IF EXISTS `send_info`;
CREATE TABLE `send_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `errordetail` varchar(60) NOT NULL ,
  `sid` bigint(50) NOT NULL,
  `uid` varchar(40) NOT NULL,
  `userreceivetime` varchar(40) NOT NULL,
  `errormsg` varchar(300),
  `mobile` varchar(20) NOT NULL,
  `reportstatus` varchar(40) NOT NULL,
  `alarmstatus` int NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `sid` (`sid`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;