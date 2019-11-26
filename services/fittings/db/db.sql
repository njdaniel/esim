-- CREATE DATABASE fittings;
-- CREATE USER esim WITH PASSWORD 'password';
-- ALTER USER esim WITH LOGIN;
-- GRANT ALL PRIVILEGES ON DATABASE postgres TO esim;

CREATE TABLE fittings (
    name VARCHAR(255) PRIMARY KEY UNIQUE NOT NULL ,
    ship_type_id INT NOT NULL ,
    description VARCHAR(2000) ,
);

CREATE TABLE items (
    flag INT ,
    quantity INT ,
    type_id INT ,
    fittings FOREIGN KEY REFERENCES fittings(name)
);