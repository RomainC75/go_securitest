CREATE TABLE scans (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGSERIAL NOT NULL,
    scenario INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE port_ranges (
    id BIGSERIAL PRIMARY KEY,
    scan_id BIGSERIAL NOT NULL,
    range_min INT NOT NULL,
    range_max INT,
    is_unique BOOLEAN,
    FOREIGN KEY (scan_id) REFERENCES scans(id)
);

CREATE TABLE ip_ranges (
    id BIGSERIAL PRIMARY KEY,
    scan_id BIGSERIAL NOT NULL,
    ip_min TEXT NOT NULL,
    ip_max TEXT,
    is_unique BOOLEAN,
    FOREIGN KEY (scan_id) REFERENCES scans(id)
);



