CREATE TABLE IF NOT EXISTS  users (
    id VARCHAR(255) NOT NULL,
    auth0_id VARCHAR(255),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS  creators (
    id VARCHAR(255) NOT NULL,
    user_id VARCHAR(255),
    type VARCHAR(255),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id)
);
