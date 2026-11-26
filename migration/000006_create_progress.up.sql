CREATE TABLE learn_backend.progress(
    id               UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id          UUID NOT NULL REFERENCES learn_backend.users(id),
    word_id          UUID NOT NULL REFERENCES learn_backend.words(id),
    repeat_count     INT NOT NULL DEFAULT 0,
    last_reviewed_at TIMESTAMPTZ,
    next_review_at   TIMESTAMPTZ,
    is_known         BOOLEAN NOT NULL DEFAULT FALSE
);
