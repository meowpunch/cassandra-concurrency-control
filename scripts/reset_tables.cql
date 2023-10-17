USE payment;

DROP TABLE IF EXISTS payment_id_by_conversation_id;
DROP TABLE IF EXISTS event_by_conversation_id;
DROP TABLE IF EXISTS event_by_payment_id;

CREATE TABLE IF NOT EXISTS payment_id_by_conversation_id
(
    conversation_id text,
    payment_id      timeuuid,
    PRIMARY KEY (conversation_id)
);

CREATE TABLE IF NOT EXISTS event_by_conversation_id
(
    conversation_id    text,
    id                 timeuuid,
    event_type         text,
    event_data_as_json text,
    PRIMARY KEY ((conversation_id), id)
);

CREATE TABLE IF NOT EXISTS event_by_payment_id
(
    payment_id         timeuuid,
    conversation_id    text,
    id                 timeuuid,
    event_type         text,
    event_data_as_json text,
    PRIMARY KEY ((payment_id), id)
);