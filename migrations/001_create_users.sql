CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name varchar(100) NOT NULL,
    email varchar(150) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    role VARCHAR(20) DEFAULT 'USER',
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);