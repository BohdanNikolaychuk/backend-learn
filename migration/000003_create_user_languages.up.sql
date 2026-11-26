CREATE TABLE learn_backend.user_languages(
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id       UUID NOT NULL REFERENCES learn_backend.users(id),
    language_code VARCHAR(10) NOT NULL REFERENCES learn_backend.languages(code),
    UNIQUE(user_id, language_code)
);
