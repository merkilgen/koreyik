CREATE TABLE IF NOT EXISTS animes (
    id INTEGER PRIMARY KEY,
    thumbnail_url TEXT DEFAULT '',

    title_kk TEXT DEFAULT 'Атсыз',
    title_en TEXT DEFAULT 'Untitled',
    title_jp TEXT DEFAULT 'Untitled'
);