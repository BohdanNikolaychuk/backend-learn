CREATE TABLE learn_backend.words(
    id                  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    deck_id             UUID NOT NULL REFERENCES learn_backend.decks(id),
    word                VARCHAR(255) NOT NULL,
    transcription       VARCHAR(255),
    part_of_speech      VARCHAR(50),
    translation         TEXT NOT NULL,
    definition          TEXT,
    example             TEXT,
    example_translation TEXT,
    notes               TEXT,
    level               VARCHAR(10),
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
