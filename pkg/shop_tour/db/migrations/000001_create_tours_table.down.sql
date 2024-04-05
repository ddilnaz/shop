--C:\Users\Lenovo\Desktop\shop\pkg\abr-plus\db\migrations\000001_create_tours_table.down.sql-- Удаляем внешние ключи из таблицы order_product_item
ALTER TABLE order_product_item DROP CONSTRAINT IF EXISTS fk_order_id;
ALTER TABLE order_product_item DROP CONSTRAINT IF EXISTS fk_product_item_id;

DROP TABLE IF EXISTS order_product_item;

ALTER TABLE orders DROP CONSTRAINT IF EXISTS fk_user_id;

DROP TABLE IF EXISTS orders;

DROP TABLE IF EXISTS product_item;

DROP TABLE IF EXISTS users;
