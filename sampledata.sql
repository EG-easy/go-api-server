DROP TABLE IF EXISTS ItemInfo;
CREATE TABLE Users (
  Id INT  PRIMARY KEY auto_increment,
  Age INT,
  Name VARCHAR (255),
  Description VARCHAR(255)
);

INSERT Users (Age, Name, Description) values (22, "suzuki", "https://www.cryptokitties.co/img/1");
INSERT Users (Age, Name, Description) values (34, "sato", "https://www.mycryptoheroes.net/img/1");
INSERT Users (Age, Name, Description) values (45, "kobayashi", "https://www.etheremon.com/img/1");
