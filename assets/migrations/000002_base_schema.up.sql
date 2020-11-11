-- CREATE TYPE keyvalue AS
-- (
--     key          text,
--     value_type   text,
--     value_string text,
--     value_bool   boolean,
--     value_long   bigint,
--     value_double decimal,
--     value_binary bytea
-- );
--
-- CREATE TYPE log AS
-- (
--     ts     bigint,
--     fields keyvalue[]
-- );
--
-- CREATE TYPE span_ref AS
-- (
--     ref_type text,
--     trace_id bytea,
--     span_id  bigint
-- );
--
-- CREATE TYPE process AS
-- (
--     service_name text,
--     tags         keyvalue[]
--
-- );

CREATE TABLE traces
(
    trace_id       bytea,
    span_id        bigint,
    span_hash      bigint,
    parent_id      bigint,
    operation_name text,
    flags          int,
    start_time     bigint,
    duration       bigint,
    tags           jsonb,
    logs           jsonb,
    refs           jsonb,
    process        jsonb,
    PRIMARY KEY (trace_id, span_id, span_hash)
);

CREATE TABLE service_names
(
    service_name text,
    PRIMARY KEY (service_name)
);

CREATE TABLE operation_names_v2
(
    service_name   text,
    span_kind      text,
    operation_name text,
    PRIMARY KEY (service_name, span_kind, operation_name)
);

CREATE TABLE service_operation_index
(
    service_name   text,
    operation_name text,
    start_time     bigint,
    trace_id       bytea,
    PRIMARY KEY (service_name, operation_name, start_time)
);

CREATE TABLE service_name_index
(
    service_name text,
    bucket       int,
    start_time   bigint,
    trace_id     bytea,
    PRIMARY KEY (service_name, bucket, start_time)
);

CREATE TABLE duration_index
(
    service_name   text,
    operation_name text,
    bucket         timestamp,
    duration       bigint,
    start_time     bigint,
    trace_id       bytea,
    PRIMARY KEY (service_name, operation_name, bucket, duration, start_time, trace_id)
);

CREATE TABLE tag_index
(
    service_name text,
    tag_key      text,
    tag_value    text,
    start_time   bigint,
    trace_id     bytea,
    span_id      bigint,
    PRIMARY KEY (service_name, tag_key, tag_value, start_time, trace_id, span_id)
);

-- CREATE TYPE dependency AS
-- (
--     parent     text,
--     child      text,
--     call_count bigint,
--     source     text
-- );

CREATE TABLE dependencies_v2
(
    ts_bucket    timestamp,
    ts           timestamp,
    dependencies jsonb,
    PRIMARY KEY (ts_bucket, ts)
);
