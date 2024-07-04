-- Active: 1718919020656@@127.0.0.1@5432@postgres


CREATE TABLE habits (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    frequency frequency,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Habit logs table
CREATE TABLE habit_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    habit_id UUID REFERENCES habits(id),
    logged_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    notes TEXT
);


-- Users table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- User profiles table
CREATE TABLE user_profiles (
    user_id UUID PRIMARY KEY REFERENCES users(id),
    full_name VARCHAR(100),
    bio TEXT,
    location VARCHAR(100),
    avatar_url VARCHAR(255)
);


-- Insert sample data into users table
INSERT INTO users (username, email, password_hash)
VALUES
('user1', 'user1@example.com', 'hashedpassword1'),
('user2', 'user2@example.com', 'hashedpassword2'),
('user3', 'user3@example.com', 'hashedpassword3'),
('user4', 'user4@example.com', 'hashedpassword4'),
('user5', 'user5@example.com', 'hashedpassword5');

-- Insert sample data into user_profiles table
INSERT INTO user_profiles (user_id, full_name, bio, location, avatar_url)
SELECT id, 'User One', 'Bio for User One', 'Location One', 'http://example.com/avatar1.png' FROM users WHERE username = 'user1'
UNION ALL
SELECT id, 'User Two', 'Bio for User Two', 'Location Two', 'http://example.com/avatar2.png' FROM users WHERE username = 'user2'
UNION ALL
SELECT id, 'User Three', 'Bio for User Three', 'Location Three', 'http://example.com/avatar3.png' FROM users WHERE username = 'user3'
UNION ALL
SELECT id, 'User Four', 'Bio for User Four', 'Location Four', 'http://example.com/avatar4.png' FROM users WHERE username = 'user4'
UNION ALL
SELECT id, 'User Five', 'Bio for User Five', 'Location Five', 'http://example.com/avatar5.png' FROM users WHERE username = 'user5';

-- Insert sample data into habits table
INSERT INTO habits (user_id, name, description, frequency)
VALUES 
((SELECT id FROM users WHERE username = 'user1'), 'Exercise', 'Daily morning workout', 'daily'),
((SELECT id FROM users WHERE username = 'user2'), 'Read Book', 'Read at least one chapter', 'weekly'),
((SELECT id FROM users WHERE username = 'user3'), 'Meditate', 'Meditation session', 'monthly'),
((SELECT id FROM users WHERE username = 'user4'), 'Clean House', 'Weekly cleaning routine', 'daily'),
((SELECT id FROM users WHERE username = 'user5'), 'Write Journal', 'Daily journaling', 'weekly');

-- Insert sample data into habit_logs table
INSERT INTO habit_logs (habit_id, notes)
SELECT id, 'Completed 30 minutes of exercise' FROM habits WHERE name = 'Exercise'
UNION ALL
SELECT id, 'Read two chapters of a novel' FROM habits WHERE name = 'Read Book'
UNION ALL
SELECT id, 'Meditated for 20 minutes' FROM habits WHERE name = 'Meditate'
UNION ALL
SELECT id, 'Cleaned the living room and kitchen' FROM habits WHERE name = 'Clean House'
UNION ALL
SELECT id, 'Wrote about my day and thoughts' FROM habits WHERE name = 'Write Journal';
