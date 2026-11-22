CREATE SCHEMA learn_backend;

CREATE TABLE learn_backend.users (
    id        BIGSERIAL PRIMARY KEY,
	version	  BIGINT NOT NULL DEFAULT 1,
    email VARCHAR(255) NOT NULL UNIQUE CHECK(email LIKE '%@%'),
    password  VARCHAR(255) NOT NULL,
    full_name VARCHAR(100) NOT NULL CHECK(char_length(full_name) BETWEEN 3 AND 100),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

