DROP TABLE IF EXISTS Users;
CREATE TABLE Users (
  Id INT  PRIMARY KEY auto_increment,
  Age INT,
  Name VARCHAR (255),
  Description VARCHAR(255),
  SiblingInfo json
);

INSERT Users (Age, Name, Description, SiblingInfo) values (22, "suzuki", "ダミーデータ1", '[{"sibling":"brother", "name":"akito"},{"sibling":"sister", "name":"ayaka"}, {"sibling":"brother", "name":"taro"}]');
INSERT Users (Age, Name, Description, SiblingInfo) values (34, "sato", "ダミーデータ2", '[{"sibling":"sister", "name":"kana"},{"sibling":"brother", "name":"ken"}]');
INSERT Users (Age, Name, Description, SiblingInfo) values (45, "kobayashi", "ダミーデータ3", '[{"sibling":"brother", "name":"ryo"}]');
