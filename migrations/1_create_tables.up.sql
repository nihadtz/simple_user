CREATE TABLE Users (
    id int(11) NOT NULL AUTO_INCREMENT,
    name varchar(128) NOT NULL DEFAULT '',
    updated int(11) NOT NULL DEFAULT 0,
    year_of_birth int(11) NOT NULL DEFAULT 0,
    PRIMARY KEY(id)
    ) ENGINE=INNODB;