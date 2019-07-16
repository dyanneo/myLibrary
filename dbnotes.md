DB Notes

  DROP TABLE IF EXISTS `book`;
  CREATE TABLE `book` (
  `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(256) NOT NULL,
  `author` varchar(256) NOT NULL,
  `formattype` varchar(30) NOT NULL,
  `location` varchar(256) NOT NULL,
  PRIMARY KEY (`id`)
) AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

ALTER TABLE `golibrary`.`book` 
ADD COLUMN `isbn` VARCHAR(45) NULL AFTER `location`;