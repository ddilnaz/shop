-- db/migrations/{timestamp}_create_tables.up.sql
-- C:\Users\Lenovo\Desktop\shop\pkg\abr-plus\db\migrations\000001_create_tours_table.up.sql
-- Создание таблицы для продуктов, заказов и пользователей
CREATE TABLE IF NOT EXISTS product_item (
    id          bigserial PRIMARY KEY,
    title       text NOT NULL,
    description text,
    price       int NOT NULL,
    created_at  timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at  timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS "orderr" (
    id          bigserial PRIMARY KEY,
    user_id     bigserial REFERENCES "user" (id),
    created_at  timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at  timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    status      text NOT NULL,
    FOREIGN KEY (user_id) REFERENCES "user" (id)
);

CREATE TABLE IF NOT EXISTS "user" (
    id          bigserial PRIMARY KEY,
    name        text NOT NULL,
    email       text NOT NULL,
    created_at  timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at  timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS order_product_item (
    id             bigserial PRIMARY KEY,
    order_id       bigserial REFERENCES "order" (id),
    product_item_id bigserial REFERENCES product_item (id),
    quantity       int NOT NULL,
    total_price    int NOT NULL,
    created_at     timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at     timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    FOREIGN KEY (order_id) REFERENCES "order" (id),
    FOREIGN KEY (product_item_id) REFERENCES product_item (id)
);
