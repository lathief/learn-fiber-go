CREATE PROCEDURE ecommerce_product.COUNT_PRICE(
    IN ORDERID INT
)
BEGIN
    DECLARE totalPrice DECIMAL(10, 2) DEFAULT 0.0;

    UPDATE order_items oi
    INNER JOIN product p on oi.product_id = p.id
    SET oi.price = oi.quantity * p.price
    WHERE oi.order_id = ORDERID;

    SELECT SUM(price) INTO totalPrice FROM order_items;

    UPDATE `order` o
    INNER JOIN order_items oi on o.id = oi.order_id
    SET o.total_price = totalPrice
    WHERE o.id = ORDERID;
END;