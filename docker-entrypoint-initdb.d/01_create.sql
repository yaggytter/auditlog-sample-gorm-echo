DROP TABLE IF EXISTS users;

CREATE TABLE users (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255),
  `email` VARCHAR(255),
  `secret` VARCHAR(255),
  `personalinfo` VARCHAR(255),
  PRIMARY KEY (`id`)
);

INSERT INTO users (name, email, secret, personalinfo) VALUES ("Yamada", "yamada@momiage.com", "himitsu", "83 years old");
INSERT INTO users (name, email, secret, personalinfo) VALUES ("Yamada2", "yamada2@momiage.com", "my number is 123", "82 years old");


