-- Shipping Service Database Schema

CREATE TABLE shipping_providers
(
    id           BIGSERIAL    NOT NULL,
    name         varchar(255) NOT NULL UNIQUE,
    code         varchar(50)  NOT NULL UNIQUE,
    api_endpoint varchar(500),
    is_active    boolean      NOT NULL DEFAULT true,
    created_by   varchar(255) NOT NULL,
    updated_by   varchar(255) NOT NULL,
    created_at   timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    version      int          NOT NULL DEFAULT 1,
    PRIMARY KEY (id)
);

CREATE TABLE shipping_methods
(
    id             BIGSERIAL      NOT NULL,
    provider_id    int8           NOT NULL,
    service_type   varchar(100)   NOT NULL,              -- STANDARD, EXPRESS, SAME_DAY
    from_city      varchar(100)   NOT NULL,
    to_city        varchar(100)   NOT NULL,
    min_weight     DECIMAL(8, 2)  NOT NULL DEFAULT 0.00, -- in kg
    max_weight     DECIMAL(8, 2)  NOT NULL,
    base_rate      DECIMAL(10, 2) NOT NULL,
    rate_per_kg    DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    estimated_days int4           NOT NULL DEFAULT 1,
    is_active      boolean        NOT NULL DEFAULT true,
    created_by     varchar(255)   NOT NULL,
    updated_by     varchar(255)   NOT NULL,
    created_at     timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at     timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    version        int            NOT NULL DEFAULT 1,
    PRIMARY KEY (id)
);

CREATE TABLE shipments
(
    id                 BIGSERIAL      NOT NULL,
    order_id           int8           NOT NULL,
    provider_id        int8           NOT NULL,
    tracking_number    varchar(255) UNIQUE,
    service_type       varchar(100)   NOT NULL,
    status             varchar(50)    NOT NULL DEFAULT 'PENDING', -- PENDING, PICKED_UP, IN_TRANSIT, OUT_FOR_DELIVERY, DELIVERED, RETURNED, CANCELLED
    weight             DECIMAL(8, 2)  NOT NULL,
    dimensions         varchar(100),                              -- LxWxH in cm
    shipping_cost      DECIMAL(10, 2) NOT NULL,
    estimated_delivery timestamp,
    actual_delivery    timestamp,
    pickup_address     text           NOT NULL,
    delivery_address   text           NOT NULL,
    recipient_name     varchar(255)   NOT NULL,
    recipient_phone    varchar(20)    NOT NULL,
    notes              text,
    created_by         varchar(255)   NOT NULL,
    updated_by         varchar(255)   NOT NULL,
    created_at         timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at         timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    version            int            NOT NULL DEFAULT 1,
    PRIMARY KEY (id)
);

CREATE TABLE shipment_tracking
(
    id          BIGSERIAL    NOT NULL,
    shipment_id int8         NOT NULL,
    status      varchar(50)  NOT NULL,
    location    varchar(255),
    description varchar(500) NOT NULL,
    timestamp   timestamp    NOT NULL,
    created_by  varchar(255) NOT NULL,
    updated_by  varchar(255) NOT NULL,
    created_at  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    version     int          NOT NULL DEFAULT 1,
    PRIMARY KEY (id)
);

-- Add constraints
ALTER TABLE shipping_methods
    ADD CONSTRAINT CHK_shipping_rates_weight CHECK (min_weight >= 0 AND max_weight > min_weight),
    ADD CONSTRAINT CHK_shipping_rates_base_rate CHECK (base_rate >= 0),
    ADD CONSTRAINT CHK_shipping_rates_rate_per_kg CHECK (rate_per_kg >= 0),
    ADD CONSTRAINT CHK_shipping_rates_estimated_days CHECK (estimated_days > 0);

ALTER TABLE shipments
    ADD CONSTRAINT CHK_shipments_weight CHECK (weight > 0),
    ADD CONSTRAINT CHK_shipments_shipping_cost CHECK (shipping_cost >= 0);

-- Add foreign key constraints
ALTER TABLE shipping_methods
    ADD CONSTRAINT FK_shipping_rates_provider FOREIGN KEY (provider_id) REFERENCES shipping_providers (id) ON DELETE CASCADE;

ALTER TABLE shipments
    ADD CONSTRAINT FK_shipments_provider FOREIGN KEY (provider_id) REFERENCES shipping_providers (id) ON DELETE RESTRICT;

ALTER TABLE shipment_tracking
    ADD CONSTRAINT FK_shipment_tracking_shipment FOREIGN KEY (shipment_id) REFERENCES shipments (id) ON DELETE CASCADE;
