CREATE TABLE users (
    Id uuid PRIMARY KEY,
    Email text NOT NULL UNIQUE,
    PhoneNumber text NOT NULL UNIQUE,
    Login text NOT NULL UNIQUE,
    Password text NOT NULL,
    SecondName text NOT NULL,
    FirstName text NOT NULL,
    Patronimic text,
    PathToAvatar text
);