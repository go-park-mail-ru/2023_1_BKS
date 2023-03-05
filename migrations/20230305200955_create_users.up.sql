CREATE TABLE users (
    id bigserial not null primary key,
    email varchar not null unique,
    encrypted_password varchar not null,
    first_name varchar not null,
    second_name varchar not null,
    card_number varchar,
    card_cvv int
)