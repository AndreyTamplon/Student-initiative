CREATE TABLE petitions (
  id bigserial not null PRIMARY KEY,
  title varchar(255) not null UNIQUE,
  author_name varchar(50) not null,
  author_email varchar(50) not null,
  date_of_creation varchar(30) not null,
  date_of_expiration varchar(30) not null,
  tags varchar(100)[],
  petition_content TEXT not null,
  number_of_signatures INTEGER not null,
  signatures_target INTEGER not null,
  signatories INTEGER[]
);