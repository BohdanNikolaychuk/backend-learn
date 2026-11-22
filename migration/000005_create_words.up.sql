CREATE TABLE learn_backend.words(
	id BIGSERIAL PRIMARY KEY,
	deck_id bigint not null REFERENCES learn_backend.decks(id),
	word varchar(255) not null,
	transcription varchar(255),
	part_of_speech varchar(50),
	translation TEXT not null,
	definition TEXT,
	example TEXT,
	example_translation TEXT,
	notes TEXT,
	level varchar(10),
	created_at TIMESTAMPTZ not null default now()
);