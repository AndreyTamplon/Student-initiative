CREATE TABLE users (
   id bigserial not null primary key,
   name varchar(255) not null,
   email varchar not null UNIQUE,
   confirmed boolean not null default false,
   encrypted_confirmation_code varchar not null,
   encrypted_password varchar not null
);
