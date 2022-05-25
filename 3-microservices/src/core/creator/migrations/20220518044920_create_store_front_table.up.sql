CREATE TABLE IF NOT EXISTS store_fronts (
    id BIGSERIAL PRIMARY KEY,
    store_name TEXT NOT NULL,
    store_url TEXT NOT NULL,
    store_logo_filename TEXT NOT NULL,
    storefront_banner_video_filename TEXT,
    storefront_title_header TEXT,
    storefront_sub_header TEXT,
    promotional_message TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);