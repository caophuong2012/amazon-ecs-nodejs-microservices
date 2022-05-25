CREATE TABLE IF NOT EXISTS collections
(
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    image_url TEXT,
    network TEXT NOT NULL,
    type TEXT,
    display_order INT,
    slug TEXT NOT NULL,
    status TEXT NOT NULL,
    creator_id UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    delete_at TIMESTAMP WITH TIME ZONE NULL
);
