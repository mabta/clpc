CREATE TABLE issue(
    id BIGSERIAL PRIMARY KEY,
    ticket VARCHAR NOT NULL,
    result VARCHAR NOT NULL,
    schedule VARCHAR NOT NULL,
    block_time bigint NOT NULL,
    block_number bigint NOT NULL,
    dateline bigint not NULL,
    date_str VARCHAR not null
);
