CREATE TABLE store (
    sku VARCHAR(32),
    name TEXT,
    price REAL
);

CREATE INDEX store_sku ON store(sku);
