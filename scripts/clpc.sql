CREATE TABLE issue(
    id BIGSERIAL PRIMARY KEY,
    ticket VARCHAR NOT NULL,
    result VARCHAR NOT NULL,
    schedule bigint NOT NULL,
    block_time bigint NOT NULL,
    block_number bigint NOT NULL,
    period bigint NOT NULL, 
    next_period bigint NOT NULL,
    next_period_schedule bigint NOT NULL,
    block_hash VARCHAR NOT NULL,
    dateline bigint not NULL,
    date_str VARCHAR not null
);
