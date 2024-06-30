CREATE TYPE footprint_category AS ENUM ('transport', 'energy', 'food', 'consumption');

CREATE TABLE carbon_footprint_logs (
    id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    category footprint_category,
    amount DECIMAL(10, 2),
    unit VARCHAR(20),
    logged_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE donations (
    id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    cause VARCHAR(100),
    amount DECIMAL(10, 2),
    donated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE donation_causes (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT
);
