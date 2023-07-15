CREATE TABLE IF NOT EXISTS order_items (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    order_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
    quantity INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL
);

ALTER TABLE order_items ADD FOREIGN KEY (order_id) REFERENCES `order` (id);
ALTER TABLE order_items ADD FOREIGN KEY (product_id) REFERENCES product (id);