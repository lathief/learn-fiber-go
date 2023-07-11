CREATE TABLE IF NOT EXISTS order_product (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    order_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL
);

ALTER TABLE order_product ADD FOREIGN KEY (order_id) REFERENCES `order` (id);
ALTER TABLE order_product ADD FOREIGN KEY (product_id) REFERENCES product (id);