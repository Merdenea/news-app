CREATE TABLE IF NOT EXISTS "sources" (
    id UUID PRIMARY KEY,

    source_name TEXT NOT NULL,
    source_url TEXT NOT NULL,
    category TEXT NOT NULL
);

INSERT INTO "sources"("id", "source_name", "source_url", "category") VALUES ('b661d911-1c41-4d48-8ed2-c44dbab6a719', 'BBC NEWS', 'http://feeds.bbci.co.uk/news/uk/rss.xml', 'UK NEWS');