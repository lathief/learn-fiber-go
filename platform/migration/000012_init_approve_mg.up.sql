CREATE TABLE IF NOT EXISTS approval (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    name_merchant VARCHAR(150) NOT NULL,
    status VARCHAR(10) NOT NULL, ## approved, pending, rejected
    description_merchant TEXT NOT NULL,
    product_merchant TEXT NOT NULL,
    approval_comment TEXT,
    approved_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
ALTER TABLE approval ADD FOREIGN KEY (user_id) REFERENCES users (id);