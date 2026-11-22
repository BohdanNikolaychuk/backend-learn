CREATE TABLE learn_backend.decks(
	id BIGSERIAL PRIMARY KEY,
	user_id BIGINT NOT NULL references learn_backend.users(id),
	name VARCHAR(100) not null,
	created_at TIMESTAMPTZ not null default now()
);