# myLibrary
Web app with mysql db and go backend to capture and organize my books based on [this example](http://www.golangprograms.com/advance-programs/example-of-golang-crud-using-mysql-from-scratch.html).

### Database
- Version: MySQL 8.0
- DB Name: golibrary
- Connection info
- Table
  DROP TABLE IF EXISTS `book`;
  CREATE TABLE `book` (
  `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(256) NOT NULL,
  `author` varchar(256) NOT NULL,
  `formattype` varchar(30) NOT NULL,
  `location` varchar(256) NOT NULL,
  PRIMARY KEY (`id`)
) AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

### Usage
- localhost:8090
- `/`: main screen
- `/add`: add a book
- `/show`: see full list
- `/print`: print full list FUTURE
