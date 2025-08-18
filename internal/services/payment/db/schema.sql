-- Payment Service Database Schema

CREATE TABLE payments
(
    id               BIGSERIAL      NOT NULL,
    idempotence_key  varchar(255)   NOT NULL UNIQUE,            -- Idempotency key from gateway
    order_id         int8           NOT NULL,
    user_id          int8           NOT NULL,
    payment_method   varchar(50)    NOT NULL,                   -- CREDIT_CARD, DEBIT_CARD, PAYPAL, BANK_TRANSFER, COD, WALLET
    amount           DECIMAL(12, 2) NOT NULL,
    currency         varchar(3)     NOT NULL DEFAULT 'VND',
    status           varchar(50)    NOT NULL DEFAULT 'PENDING', -- PENDING, PROCESSING, SUCCESS, FAILED, CANCELLED, REFUNDED
    transaction_id   varchar(255),                              -- External payment gateway transaction ID
    gateway_provider varchar(50),                               -- STRIPE, PAYPAL, VNPAY, MOMO, etc.
    gateway_response jsonb,                                     -- Store raw gateway response
    failure_reason   varchar(500),
    processed_at     timestamp,
    created_by       varchar(255)   NOT NULL,
    updated_by       varchar(255)   NOT NULL,
    created_at       timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at       timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    version          int            NOT NULL DEFAULT 1,
    PRIMARY KEY (id)
);

CREATE TABLE payment_methods
(
    id              BIGSERIAL    NOT NULL,
    user_id         int8         NOT NULL,
    method_type     varchar(50)  NOT NULL, -- CREDIT_CARD, DEBIT_CARD, PAYPAL, BANK_ACCOUNT, COD, WALLET
    provider        varchar(50),           -- VISA, MASTERCARD, PAYPAL, etc. (NULL for COD)
    masked_number   varchar(20),           -- Last 4 digits for cards (NULL for COD)
    cardholder_name varchar(255),          -- NULL for COD
    expiry_month    int2,                  -- NULL for COD
    expiry_year     int4,                  -- NULL for COD
    is_default      boolean      NOT NULL DEFAULT false,
    is_active       boolean      NOT NULL DEFAULT true,
    external_id     varchar(255),          -- ID from payment gateway (NULL for COD)
    created_by      varchar(255) NOT NULL,
    updated_by      varchar(255) NOT NULL,
    created_at      timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    version         int          NOT NULL DEFAULT 1,
    PRIMARY KEY (id)
);

CREATE TABLE refunds
(
    id               BIGSERIAL      NOT NULL,
    payment_id       int8           NOT NULL,
    order_id         int8           NOT NULL,
    refund_amount    DECIMAL(12, 2) NOT NULL,
    reason           varchar(500)   NOT NULL,
    status           varchar(50)    NOT NULL DEFAULT 'PENDING', -- PENDING, PROCESSING, SUCCESS, FAILED
    transaction_id   varchar(255),                              -- External refund transaction ID
    gateway_response jsonb,
    processed_at     timestamp,
    created_by       varchar(255)   NOT NULL,
    updated_by       varchar(255)   NOT NULL,
    created_at       timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at       timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    version          int            NOT NULL DEFAULT 1,
    PRIMARY KEY (id)
);

-- Add constraints
ALTER TABLE payments
    ADD CONSTRAINT CHK_payments_amount CHECK (amount > 0),
    ADD CONSTRAINT CHK_payments_currency CHECK (currency IN ('VND', 'USD', 'EUR'));

ALTER TABLE payment_methods
    ADD CONSTRAINT CHK_payment_methods_expiry_month CHECK (expiry_month IS NULL OR (expiry_month >= 1 AND expiry_month <= 12)),
    ADD CONSTRAINT CHK_payment_methods_expiry_year CHECK (expiry_year IS NULL OR expiry_year >= EXTRACT(YEAR FROM CURRENT_DATE));

ALTER TABLE refunds
    ADD CONSTRAINT CHK_refunds_amount CHECK (refund_amount > 0);

-- Add foreign key constraints
ALTER TABLE refunds
    ADD CONSTRAINT FK_refunds_payment FOREIGN KEY (payment_id) REFERENCES payments (id) ON DELETE CASCADE;
