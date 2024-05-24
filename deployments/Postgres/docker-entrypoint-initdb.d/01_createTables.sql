CREATE TABLE IF NOT EXISTS orders (
    order_uid VARCHAR(255) PRIMARY KEY,
    track_number VARCHAR(255),
    entry VARCHAR(255),
    locale VARCHAR(50),
    internal_signature VARCHAR(255),
    customer_id VARCHAR(255),
    delivery_service VARCHAR(255),
    shardkey VARCHAR(50),
    sm_id INT,
    date_created TIMESTAMP,
    oof_shard VARCHAR(50)
);

CREATE TABLE IF NOT EXISTS delivery (
    order_uid VARCHAR(255) REFERENCES orders(order_uid),
    name VARCHAR(255),
    phone VARCHAR(11),
    zip VARCHAR(50),
    city VARCHAR(255),
    address VARCHAR(255),
    region VARCHAR(255),
    email VARCHAR(255),
    PRIMARY KEY (order_uid)
);

CREATE TABLE IF NOT EXISTS  payment (
    transaction VARCHAR(255) PRIMARY KEY,
    order_uid VARCHAR(255) REFERENCES orders(order_uid),
    request_id VARCHAR(255),
    currency VARCHAR(50),
    provider VARCHAR(255),
    amount INT,
    payment_dt BIGINT,
    bank VARCHAR(255),
    delivery_cost INT,
    goods_total INT,
    custom_fee INT
);

CREATE TABLE IF NOT EXISTS  items (
    chrt_id BIGINT PRIMARY KEY,
    order_uid VARCHAR(255) REFERENCES orders(order_uid),
    track_number VARCHAR(255),
    price INT,
    rid VARCHAR(255),
    name VARCHAR(255),
    sale INT,
    size VARCHAR(50),
    total_price INT,
    nm_id BIGINT,
    brand VARCHAR(255),
    status INT
);