-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS employee (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    pass_hash VARCHAR(255) NOT NULL,
    coins INTEGER DEFAULT 1000
);

CREATE INDEX IF NOT EXISTS idx_employee_id ON employee(id);

CREATE TABLE IF NOT EXISTS merch (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    price INT NOT NULL CHECK (price > 0)
);

CREATE INDEX IF NOT EXISTS idx_merch_id ON merch(id);

CREATE TABLE IF NOT EXISTS purchase (
    id BIGSERIAL PRIMARY KEY,
    employee_id BIGINT REFERENCES employee(id) ON DELETE CASCADE,
    merch_id INT REFERENCES merch(id) ON DELETE CASCADE,
    quantity INT NOT NULL CHECK (quantity > 0),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_purchases_employee_id ON purchase(employee_id);

CREATE TABLE IF NOT EXISTS transactions (
    id BIGSERIAL PRIMARY KEY,
    sender_id BIGINT REFERENCES employee(id) ON DELETE SET NULL,
    receiver_id BIGINT REFERENCES employee(id) ON DELETE SET NULL,
    quantity INT NOT NULL CHECK (quantity > 0),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_transactions_sender_id ON transactions(sender_id);
CREATE INDEX IF NOT EXISTS idx_transactions_receiver_id ON transactions(receiver_id);

INSERT INTO merch (name, price) VALUES
    ('t-shirt', 80),
    ('cup', 20),
    ('book', 50),
    ('pen', 10),
    ('powerbank', 200),
    ('hoody', 300),
    ('umbrella', 200),
    ('socks', 10),
    ('wallet', 50),
    ('pink-hoody', 500);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS purchase;
DROP TABLE IF EXISTS merch;
DROP TABLE IF EXISTS employee;
-- +goose StatementEnd
