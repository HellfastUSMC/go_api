CREATE TABLE users (
    id bigserial not null primary key,
    email varchar not null unique,
    enc_pass varchar not null
);