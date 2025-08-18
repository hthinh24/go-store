-- Inventory Service Database Schema

CREATE TABLE inventory (
    id                BIGSERIAL    NOT NULL,
    product_sku_id    int8         NOT NULL UNIQUE, -- References product service's product_sku
    available_stock   int4         NOT NULL DEFAULT 0,
    reserved_stock    int4         NOT NULL DEFAULT 0,
    total_stock       int4         NOT NULL GENERATED ALWAYS AS (available_stock + reserved_stock) STORED,
    reorder_level     int4         NOT NULL DEFAULT 10, -- Minimum stock level before reorder
    max_stock_level   int4,        -- Maximum stock capacity
    warehouse_location varchar(255),
    last_updated_by   varchar(255) NOT NULL,
    created_by        varchar(255) NOT NULL,
    updated_by        varchar(255) NOT NULL,
    created_at        timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at        timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    version           int          NOT NULL DEFAULT 1,
    PRIMARY KEY (id)
);

CREATE TABLE inventory_movements (
    id              BIGSERIAL    NOT NULL,
    inventory_id    int8         NOT NULL,
    movement_type   varchar(50)  NOT NULL, -- IN, OUT, RESERVED, RELEASED, ADJUSTMENT
    quantity        int4         NOT NULL,
    reference_type  varchar(50)  NOT NULL, -- ORDER, PURCHASE, RETURN, ADJUSTMENT, MANUAL
    reference_id    int8,
    reason          varchar(500),
    previous_stock  int4         NOT NULL,
    new_stock       int4         NOT NULL,
    created_by      varchar(255) NOT NULL,
    updated_by      varchar(255) NOT NULL,
    created_at      timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    version         int          NOT NULL DEFAULT 1,
    PRIMARY KEY (id)
);

CREATE TABLE stock_reservations (
    id               BIGSERIAL    NOT NULL,
    inventory_id     int8         NOT NULL,
    reference_type   varchar(50)  NOT NULL, -- CHECKOUT, ORDER
    reference_id     int8         NOT NULL, -- checkout_session_id or order_id
    reserved_quantity int4        NOT NULL,
    expiry_time      timestamp    NOT NULL,
    status           varchar(50)  NOT NULL DEFAULT 'ACTIVE', -- ACTIVE, EXPIRED, CONFIRMED
    created_by       varchar(255) NOT NULL,
    updated_by       varchar(255) NOT NULL,
    created_at       timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at       timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    version          int          NOT NULL DEFAULT 1,
    PRIMARY KEY (id)
);

-- Add constraints
ALTER TABLE inventory
    ADD CONSTRAINT CHK_inventory_available_stock CHECK (available_stock >= 0),
    ADD CONSTRAINT CHK_inventory_reserved_stock CHECK (reserved_stock >= 0),
    ADD CONSTRAINT CHK_inventory_reorder_level CHECK (reorder_level >= 0),
    ADD CONSTRAINT CHK_inventory_max_stock CHECK (max_stock_level IS NULL OR max_stock_level >= 0);

ALTER TABLE inventory_movements
    ADD CONSTRAINT CHK_movements_quantity CHECK (quantity != 0),
    ADD CONSTRAINT CHK_movements_previous_stock CHECK (previous_stock >= 0),
    ADD CONSTRAINT CHK_movements_new_stock CHECK (new_stock >= 0);

ALTER TABLE stock_reservations
    ADD CONSTRAINT CHK_reservations_quantity CHECK (reserved_quantity > 0),
    ADD CONSTRAINT CHK_reservations_expiry CHECK (expiry_time > created_at);

-- Add foreign key constraints
ALTER TABLE inventory_movements
    ADD CONSTRAINT FK_movements_inventory FOREIGN KEY (inventory_id) REFERENCES inventory (id) ON DELETE CASCADE;

ALTER TABLE stock_reservations
    ADD CONSTRAINT FK_reservations_inventory FOREIGN KEY (inventory_id) REFERENCES inventory (id) ON DELETE CASCADE;
