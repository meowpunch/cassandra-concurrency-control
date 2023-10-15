USE payment;

BEGIN
    BATCH
    INSERT INTO payment_id_by_conversation_id (conversation_id, payment_id)
    VALUES ('conv_id', 42895d36-6b6d-11ee-b962-0242ac120002)
        IF NOT EXISTS;

    INSERT INTO event_by_conversation_id (conversation_id, id, event_type, event_data_as_json)
    VALUES ('conv_id', 42895d36-6b6d-11ee-b962-0242ac120002, 'PAYMENT_INITIALIZED_EVENT',
            '{"data": "sample"}');
    APPLY BATCH;