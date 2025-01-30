CREATE TABLE orders
(
    id            SERIAL PRIMARY KEY,
    user_id       INT            NOT NULL,
    total_amount  DECIMAL(10, 2) NOT NULL CHECK (total_amount >= 0),
    status        VARCHAR(20)    NOT NULL CHECK (status IN ('pending', 'paid', 'shipped', 'delivered', 'cancelled')),
    shipping_addr TEXT           NOT NULL,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE order_items
(
    id          SERIAL PRIMARY KEY,
    order_id    INT            NOT NULL,
    product_id  INT            NOT NULL,
    quantity    INT            NOT NULL CHECK (quantity > 0),
    unit_price  DECIMAL(10, 2) NOT NULL CHECK (unit_price >= 0),
    total_price DECIMAL(10, 2) NOT NULL CHECK (total_price >= 0),
    FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE RESTRICT
);

CREATE TABLE products
(
    id             SERIAL PRIMARY KEY,
    name           VARCHAR(255)   NOT NULL,
    description    TEXT,
    price          DECIMAL(10, 2) NOT NULL CHECK (price >= 0),
    stock_quantity INT            NOT NULL CHECK (stock_quantity >= 0),
    category       VARCHAR(100),
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for faster queries
CREATE INDEX idx_product_name ON products (name);
CREATE INDEX idx_product_category ON products (category);

-- Indexes for faster queries
CREATE INDEX idx_order_status ON orders (status);
CREATE INDEX idx_order_user ON orders (user_id);
CREATE INDEX idx_order_items_order ON order_items (order_id);
CREATE INDEX idx_order_items_product ON order_items (product_id);