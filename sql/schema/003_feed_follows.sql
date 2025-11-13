-- +goose Up
CREATE TABLE feed_follows (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    user_id uuid NOT NULL,
    CONSTRAINT fk_user_id
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    feed_id uuid NOT NULL,
    CONSTRAINT fk_feed_id
        FOREIGN KEY (feed_id) REFERENCES feeds(id) ON DELETE CASCADE,
    CONSTRAINT unique_user_feed 
        UNIQUE (user_id, feed_id)
);


-- +goose Down
DROP TABLE feed_follows;