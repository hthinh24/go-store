-- Order Service Database Schema

CREATE TABLE orders (
    id              BIGSERIAL      NOT NULL,
    idempotence_key      varchar(255)   NOT NULL UNIQUE, -- Idempotency key from gateway
    user_id         int8           NOT NULL,
    order_number    varchar(50)    NOT NULL UNIQUE,
    status          varchar(50)    NOT NULL DEFAULT 'PENDING', -- PENDING, PROCESSING, SHIPPING, SHIPPED, DELIVERED, CANCELLED, REFUNDED
    total_amount    DECIMAL(12, 2) NOT NULL DEFAULT 0.00,
    shipping_fee    DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    tax_amount      DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    discount_amount DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    final_amount    DECIMAL(12, 2) NOT NULL GENERATED ALWAYS AS (total_amount + shipping_fee + tax_amount - discount_amount) STORED,
    payment_status  varchar(50)    NOT NULL DEFAULT 'PENDING', -- PENDING, PAID, FAILED, REFUNDED
    shipping_status varchar(50)    NOT NULL DEFAULT 'PENDING', -- PENDING, PROCESSING, SHIPPED, DELIVERED
    notes           text,
    created_by      varchar(255)   NOT NULL,
    updated_by      varchar(255)   NOT NULL,
    created_at      timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    version         int            NOT NULL DEFAULT 1,
    PRIMARY KEY (id)
);

CREATE TABLE order_items (
    id             BIGSERIAL      NOT NULL,
    order_id       int8           NOT NULL,
    product_id     int8           NOT NULL,
    product_sku_id int8           NOT NULL,
    product_name   varchar(255)   NOT NULL, -- Snapshot at order time
    sku            varchar(255)   NOT NULL, -- Snapshot at order time
    quantity       int4           NOT NULL,
    unit_price     DECIMAL(10, 2) NOT NULL, -- Price at order time
    total_price    DECIMAL(10, 2) NOT NULL GENERATED ALWAYS AS (quantity * unit_price) STORED,
    created_by     varchar(255)   NOT NULL,
    updated_by     varchar(255)   NOT NULL,
    created_at     timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at     timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    version        int            NOT NULL DEFAULT 1,
    PRIMARY KEY (id)
);

CREATE TABLE order_addresses (
    id           BIGSERIAL    NOT NULL,
    order_id     int8         NOT NULL,
    full_name    varchar(255) NOT NULL,
    phone        varchar(20)  NOT NULL,
    email        varchar(255),
    address_line varchar(500) NOT NULL,
    city         varchar(100) NOT NULL,
    state        varchar(100) NOT NULL,
    postal_code  varchar(20)  NOT NULL,
    country      varchar(100) NOT NULL DEFAULT 'VietNam',
    created_by   varchar(255) NOT NULL,
    updated_by   varchar(255) NOT NULL,
    created_at   timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    version      int          NOT NULL DEFAULT 1,
    PRIMARY KEY (id)
);

-- Add constraints
ALTER TABLE orders
    ADD CONSTRAINT CHK_orders_total_amount CHECK (total_amount >= 0),
    ADD CONSTRAINT CHK_orders_shipping_fee CHECK (shipping_fee >= 0),
    ADD CONSTRAINT CHK_orders_tax_amount CHECK (tax_amount >= 0),
    ADD CONSTRAINT CHK_orders_discount_amount CHECK (discount_amount >= 0),
    ADD CONSTRAINT CHK_orders_final_amount CHECK (final_amount >= 0);

ALTER TABLE order_items
    ADD CONSTRAINT CHK_order_items_quantity CHECK (quantity > 0),
    ADD CONSTRAINT CHK_order_items_unit_price CHECK (unit_price >= 0),
    ADD CONSTRAINT CHK_order_items_total_price CHECK (total_price >= 0);

-- Add foreign key constraints
ALTER TABLE order_items
    ADD CONSTRAINT FK_order_items_order FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE;

ALTER TABLE order_addresses
    ADD CONSTRAINT FK_order_addresses_order FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE;
