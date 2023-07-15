CREATE TABLE IF NOT EXISTS user_role (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    role_id BIGINT NOT NULL
);

ALTER TABLE user_role ADD FOREIGN KEY (user_id) REFERENCES users (id);
ALTER TABLE user_role ADD FOREIGN KEY (role_id) REFERENCES roles (id);