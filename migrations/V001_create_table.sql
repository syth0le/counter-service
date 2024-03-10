CREATE TABLE IF NOT EXISTS counter_table
(
    id         TEXT                     NOT NULL,
    reader_id  TEXT                     NOT NULL,
    sender_id  TEXT                     NOT NULL,
    is_read    bool                     NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE,

    CONSTRAINT pk_counter_table PRIMARY KEY (id)
);