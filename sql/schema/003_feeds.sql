-- +goose Up
CREATE TABLE feeds {
    id UUID PRIMARY_KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL,
    user_id UUID REFRENCES users(id) ON DELETE CASCADE
};

-- +goose Down

DROP TABLES feeds;