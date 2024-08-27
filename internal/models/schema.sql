-- DROP TABLE animes;

CREATE TABLE IF NOT EXISTS animes (
    id INTEGER PRIMARY KEY,
    thumbnail_url TEXT,
    description TEXT,
    rating TEXT,

    title_kk TEXT,
    title_en TEXT,
    title_jp TEXT,

    status TEXT,
    started_airing TIMESTAMP,
    finished_airing TIMESTAMP,

    genres TEXT ARRAY,
    themes TEXT ARRAY,

    seasons INTEGER,
    episodes INTEGER,
    duration INTEGER,

    studios TEXT ARRAY,
    producers TEXT ARRAY
);