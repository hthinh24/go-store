-- Cart Service Database Schema

CREATE TABLE cart (
    id          BIGSERIAL    NOT NULL,
    user_id     int8         NOT NULL,
    status      varchar(255) NOT NULL DEFAULT 'ACTIVE', -- ACTIVE, ABANDONED, CONVERTED
    created_by  varchar(255) NOT NULL,
    updated_by  varchar(255) NOT NULL,
    created_at  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    version     int          NOT NULL DEFAULT 1,
    PRIMARY KEY (id)
);

CREATE TABLE cart_item (
    id             BIGSERIAL      NOT NULL,
    cart_id        int8           NOT NULL,
    product_id     int8           NOT NULL,  -- References product service's product
    product_sku_id int8           NOT NULL,  -- References product service's product_sku
    quantity       int4           NOT NULL DEFAULT 1,
    unit_price     DECIMAL(10, 2) NOT NULL,  -- Snapshot of price at time of adding
    total_price    DECIMAL(10, 2) NOT NULL GENERATED ALWAYS AS (quantity * unit_price) STORED,
    status         varchar(255)   NOT NULL DEFAULT 'ACTIVE', -- Follow product status
    created_by     varchar(255)   NOT NULL,
    updated_by     varchar(255)   NOT NULL,
    created_at     timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at     timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    version        int            NOT NULL DEFAULT 1,
    PRIMARY KEY (id),
    UNIQUE(cart_id, product_sku_id)  -- Prevent duplicate items
);

-- Add constraints
ALTER TABLE cart_item
    ADD CONSTRAINT CHKcart_item_quantity CHECK (quantity > 0),
    ADD CONSTRAINT CHKcart_item_unit_price CHECK (unit_price >= 0),
    ADD CONSTRAINT CHKcart_item_total_price CHECK (total_price >= 0);

-- Add foreign key constraints
ALTER TABLE cart_item
    ADD CONSTRAINT FKcart_item_cart FOREIGN KEY (cart_id) REFERENCES cart (id) ON DELETE CASCADE;

-- ====================================================================
-- INDEXES FOR PERFORMANCE OPTIMIZATION
-- ====================================================================

-- Cart table indexes
CREATE INDEX idx_cart_user_id ON cart (user_id);
CREATE INDEX idx_cart_status ON cart (status);

-- Cart newItem table indexes
CREATE INDEX idx_cart_item_cart_id ON cart_item (cart_id);
CREATE INDEX idx_cart_item_status ON cart_item (status);
CREATE INDEX idx_cart_item_created_at ON cart_item (created_at);
