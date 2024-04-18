CREATE TABLE pois (
    id SERIAL,
    name VARCHAR(80) NOT NULL,
    x_coord INTEGER NOT NULL CHECK (x_coord >= 0),
    y_coord INTEGER NOT NULL CHECK (y_coord >= 0)
);
