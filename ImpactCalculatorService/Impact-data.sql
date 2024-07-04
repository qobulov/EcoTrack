CREATE TYPE footprint_category AS ENUM ('transport', 'energy', 'food', 'consumption');

CREATE TABLE carbon_footprint_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id),
    category footprint_category,
    amount DECIMAL(10, 2),
    unit VARCHAR(20),
    logged_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE donations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id),
    cause VARCHAR(100),
    amount DECIMAL(10, 2),
    donated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE donation_causes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    description TEXT
);

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT UUID_RAND_GEN(),
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE groups (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL
);

CREATE TABLE group_members (
    user_id UUID REFERENCES users(id),
    group_id UUID REFERENCES groups(id),
    PRIMARY KEY (user_id, group_id)
);

INSERT INTO groups (name)
VALUES
    ('Group 1'),
    ('Group 2');

INSERT INTO group_members (user_id, group_id)
VALUES
    ((SELECT id FROM users WHERE username = 'user1'), (SELECT id FROM groups WHERE name = 'Group 1')),
    ((SELECT id FROM users WHERE username = 'user2'), (SELECT id FROM groups WHERE name = 'Group 1')),
    ((SELECT id FROM users WHERE username = 'user3'), (SELECT id FROM groups WHERE name = 'Group 2')),
    ((SELECT id FROM users WHERE username = 'user4'), (SELECT id FROM groups WHERE name = 'Group 2')),
    ((SELECT id FROM users WHERE username = 'user5'), (SELECT id FROM groups WHERE name = 'Group 2'));

INSERT INTO carbon_footprint_logs (user_id, category, amount, unit)
VALUES
    ((SELECT id FROM users WHERE username = 'user1'), 'transport', 20.50, 'kgCO2e'),
    ((SELECT id FROM users WHERE username = 'user2'), 'energy', 15.00, 'kgCO2e'),
    ((SELECT id FROM users WHERE username = 'user3'), 'food', 10.75, 'kgCO2e'),
    ((SELECT id FROM users WHERE username = 'user4'), 'consumption', 5.25, 'kgCO2e'),
    ((SELECT id FROM users WHERE username = 'user5'), 'transport', 8.90, 'kgCO2e');


INSERT INTO donations (user_id, cause, amount)
VALUES
    ((SELECT id FROM users WHERE username = 'user1'), 'Planting Trees', 50.00),
    ((SELECT id FROM users WHERE username = 'user2'), 'Clean Water Initiative', 75.00),
    ((SELECT id FROM users WHERE username = 'user3'), 'Renewable Energy Fund', 100.00),
    ((SELECT id FROM users WHERE username = 'user4'), 'Wildlife Protection', 30.00),
    ((SELECT id FROM users WHERE username = 'user5'), 'Ocean Cleanup', 45.00);


INSERT INTO donation_causes (name, description)
VALUES
    ('Planting Trees', 'Support tree planting initiatives to combat deforestation and climate change.'),
    ('Clean Water Initiative', 'Provide clean and safe drinking water to communities in need.'),
    ('Renewable Energy Fund', 'Promote the use of renewable energy sources to reduce carbon emissions.'),
    ('Wildlife Protection', 'Conserve wildlife habitats and protect endangered species.'),
    ('Ocean Cleanup', 'Support efforts to clean up plastic pollution in our oceans.');
