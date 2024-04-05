CREATE TABLE IF NOT EXISTS users (
    id          bigserial PRIMARY KEY,
    name        text NOT NULL,
    email       text NOT NULL,
    created_at  timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at  timestamp(0) with time zone NOT NULL DEFAULT NOW()
);
CREATE TABLE IF NOT EXISTS product_item (
    id          bigserial PRIMARY KEY,
    title       text NOT NULL,
    description text,
    price       int NOT NULL,
    created_at  timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at  timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS orders (
    order_id    bigserial PRIMARY KEY,
    item_id     bigserial REFERENCES product_item (id),
    user_id     bigserial REFERENCES users (id),
    created_at  timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at  timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    status      text NOT NULL DEFAULT 'Pending'
);