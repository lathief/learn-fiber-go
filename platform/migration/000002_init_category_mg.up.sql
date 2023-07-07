CREATE TABLE IF NOT EXISTS category (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    description TEXT DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

ALTER TABLE product ADD FOREIGN KEY (category_id) REFERENCES category (id);

INSERT INTO category (name, description) VALUES ('televisi',' salah satu media publik yang memiliki tiga fungsi sebagai alat komunikasi massa');
INSERT INTO category (name, description) VALUES ('smartphone','Smartphone merupakan gadget genggam elektronik yang mencakup fungsionalitas lanjutan selain melakukan panggilan telepon dan mengirim pesan teks.');
INSERT INTO category (name, description) VALUES ('laptop','komputer pribadi yang dapat dipindahkan dan dibawa dengan mudah sehingga dapat digunakan di banyak tempat.');

INSERT INTO product (name, price, description, category_id) VALUES ('Samsung A12',2590000,'Smartphone merk samsung',2);
INSERT INTO product (name, price, description, category_id) VALUES ('Samsung A13',4590000,'Smartphone merk samsung',2);
INSERT INTO product (name, price, description, category_id) VALUES ('TV TCL OLED',9590000,'SmartTV merk TCL',1);
INSERT INTO product (name, price, description, category_id) VALUES ('TV Samsung OLED',10590000,'SmartTV merk amsung',1);
INSERT INTO product (name, price, description, category_id) VALUES ('TV Xiomi OLED',5590000,'SmartTV merk xiomi',1);
INSERT INTO product (name, price, description, category_id) VALUES ('Laptop xiomi i5',6590000,'Laptop merk xiomi',3);
INSERT INTO product (name, price, description, category_id) VALUES ('Laptop Toshiba i3',3590000,'Laptop merk toshiba',3);
INSERT INTO product (name, price, description, category_id) VALUES ('Laptop Lenovo i7',8590000,'Laptop merk lenovo',3);