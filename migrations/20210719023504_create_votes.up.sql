CREATE TABLE IF NOT EXISTS work_votes (
                                       id bigserial not null primary key,
                                       work_id bigint not null,
                                       user_email varchar(50) not null,
                                       contest varchar(50) not null,
                                       UNIQUE (work_id, user_email)
);