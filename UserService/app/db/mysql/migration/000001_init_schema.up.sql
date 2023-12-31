CREATE TABLE IF NOT EXISTS`User` (
  `id` INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `fullName` VARCHAR(255) NOT NULL,
  `gender` VARCHAR(255) NOT NULL,
  `email` VARCHAR(255) NOT NULL,
  `datecreated` DATETIME NOT NULL,
  `dateupdated` DATETIME
);

CREATE TABLE IF NOT EXISTS `User_Password` (
  `id` INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user_id` INT UNIQUE NOT NULL,
  `password` VARCHAR(255) NOT NULL
);

CREATE UNIQUE INDEX `User_index_email` ON `User` (`email`);

