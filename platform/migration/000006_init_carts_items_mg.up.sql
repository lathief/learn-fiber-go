CREATE TABLE IF NOT EXISTS cart_items (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    cart_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
    quantity INT NOT NULL
);

ALTER TABLE cart_items ADD FOREIGN KEY (cart_id) REFERENCES cart (id);
ALTER TABLE cart_items ADD FOREIGN KEY (product_id) REFERENCES product (id);