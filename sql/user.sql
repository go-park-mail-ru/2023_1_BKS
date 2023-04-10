CREATE TABLE users (
    Id uuid PRIMARY KEY,
    Email varchar(20) NOT NULL UNIQUE,
    PhoneNumber varchar(10) NOT NULL UNIQUE,
    Login varchar(10) NOT NULL UNIQUE,
    Password varchar(20) NOT NULL,
    FirstName varchar(30) NOT NULL,
    SecondName varchar(30) NOT NULL,
    Patronimic varchar(30),
    Avatar text
);