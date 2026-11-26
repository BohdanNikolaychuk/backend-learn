
CREATE TABLE learn_backend.refresh_tokens (
	id         	UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	user_id		UUID NOT NULL REFERENCES learn_backend.users(id),
	token 		TEXT NOT NULL UNIQUE,
	expires_at 	TIMESTAMPTZ NOT NULL,
	created_at 	TIMESTAMPTZ NOT NULL DEFAULT NOW()
);