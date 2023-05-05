
ALTER TABLE orders 
    ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE order_items 
    ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;


ALTER TABLE orders
ADD promo_code_id varchar(4) DEFAULT NULL;


ALTER TABLE orders ADD FOREIGN KEY (promo_code_id) REFERENCESpromo_codes(promo_code_id); 