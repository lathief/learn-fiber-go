CREATE TABLE IF NOT EXISTS roles (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    role_name VARCHAR(50) NOT NULL,
    description TEXT DEFAULT '',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

ALTER TABLE users ADD FOREIGN KEY (role_id) REFERENCES roles (id);

INSERT INTO roles (role_name, description) VALUES ('role_super_admin', 'Super Admin yang mengatur semua user baik pemilik toko maupun pembeli');
INSERT INTO roles (role_name, description) VALUES ('role_admin_toko', 'Admin Toko yang mengatur semua aktivitas di tokonya sendiri');
INSERT INTO roles (role_name, description) VALUES ('role_customer', 'Customer yang hanya melakukan aktivitas jual beli');

INSERT INTO users (username, first_name, last_name, email, password, phone_number, address, role_id)
VALUES  ('spadmin', 'super', 'admin', 'admin@gmail.com', '1234', '0895336958353', 'Jalan', 1)
