CREATE TABLE users (
    id   SERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password  TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- CREATE TABLE lists (
--     id   SERIAL PRIMARY KEY,
--     name TEXT NOT NULL,
--     created_at TIMESTAMP NOT NULL,
--     updated_at TIMESTAMP NOT NULL,
--     user_id INT NOT NULL,
--     FOREIGN KEY (user_id) REFERENCES users(id),
--     UNIQUE (name, user_id)
-- );

-- CREATE TABLE items (
--     id   SERIAL PRIMARY KEY,
--     name TEXT NOT NULL,
--     description TEXT,
--     date TIMESTAMP,
--     created_at TIMESTAMP NOT NULL,
--     updated_at TIMESTAMP NOT NULL,
--     user_creator_id INT NOT NULL,
--     FOREIGN KEY (user_creator_id) REFERENCES users(id)
-- );

-- CREATE TABLE list_item (
--     id SERIAL PRIMARY KEY,
--     list_id INT NOT NULL,
--     item_id INT NOT NULL,
--     FOREIGN KEY (list_id) REFERENCES lists(id),
--     FOREIGN KEY (item_id) REFERENCES items(id)
-- );
