CREATE table learn_backend.progress(
	id bigserial primary key,
	user_id bigint not null references learn_backend.users(id),
	word_id bigint not null references learn_backend.words(id),
	repeat_count int not null DEFAULT 0,
	last_reviewed_at TIMESTAMPTZ,
	next_review_at TIMESTAMPTZ,
	is_known boolean not null DEFAULT FALSE
);
