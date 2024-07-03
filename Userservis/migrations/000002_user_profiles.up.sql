CREATE TABLE IF NOT EXISTS user_profiles (
    user_id UUID PRIMARY KEY REFERENCES users(id),
    full_name VARCHAR(100),
    bio TEXT,
    location VARCHAR(100),
    avatar_url VARCHAR(255)
);