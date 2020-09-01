-- +goose Up
CREATE TABLE IF NOT EXISTS users {
    id SERIAL PRIMARY KEY
    username VARCHAR(100),
    password VARCHAR(255),
    email VARCHAR(255),
    avatar VARCHAR(255),
    signature VARCHAR(255),
    email_confirmed_at TIMESTAMP,
    created_on TIMESTAMP NOT NULL,
    last_seen TIMESTAMP
};

CREATE TABLE IF NOT EXISTS groups {
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    description VARCHAR(255)
};

CREATE TABLE IF NOT EXISTS user_groups {
    group_id INTEGER REFERENCES groups(id) NOT NULL,
    user_id INTEGER REFERENCES users(id) NOT NULL,
    is_display TINYINT DEFAULT 0
};

CREATE TABLE IF NOT EXISTS permissions {
    id SERIAL PRIMARY KEY,
    permission_name VARCHAR(100),
    content_name VARCHAR(100),
    default_value TINYINT DEFAULT 0
};

CREATE TABLE IF NOT EXISTS permission_role {
    permission_id INTEGER REFERENCES permissions(id) NOT NULL,
    group_id INTEGER REFERENCES groups(id) NOT NULL,
    value TINYINT DEFAULT 0
};

CREATE TABLE IF NOT EXISTS forums {
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    description VARCHAR(255),
    is_link TINYINT DEFAULT 0,
    link VARCHAR(255),
    parent_id INTEGER REFERENCES forums(id) DEFAULT NULL,
    order INTEGER DEFAULT 0
};

CREATE TABLE IF NOT EXISTS topics {
    id SERIAL PRIMARY KEY,
    forum_id INTEGER REFERENCES forums(id) NOT NULL,
    user_id INTEGER REFERENCES users(id) NOT NULL,
    title VARCHAR(255),
    body TEXT,
    locked TINYINT DEFAULT 0,
    stiky TINYINT DEFAULT 0
};

CREATE TABLE IF NOT EXISTS posts {
    id SERIAL PRIMARY KEY,
    thread_id INTEGER REFERENCES threads(id) NOT NULL,
    user_id INTEGER REFERENCES users(id) NOT NULL,
    body TEXT,
    created_on TIMESTAMP NOT NULL,
    updated_on TIMESTAMP
};

-- +goose Down
DROP TABLE posts;
DROP TABLE topics;
DROP TABLE forums;
DROP TABLE permission_role;
DROP TABLE permissions;
DROP TABLE user_groups;
DROP TABLE groups;
DROP TABLE users;
