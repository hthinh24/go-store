CREATE TABLE category
(
    id                 BIGSERIAL    NOT NULL,
    name               varchar(255) NOT NULL UNIQUE,
    description        varchar(255) NOT NULL,
    parent_category_id int8,
    PRIMARY KEY (id)
);

-- Create the brand table
CREATE TABLE brand
(
    id          BIGSERIAL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    created_by  varchar(255) NOT NULL,
    updated_by  varchar(255) NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE product
(
    id                BIGSERIAL      NOT NULL,
    name              varchar(255)   NOT NULL,
    description       TEXT           NOT NULL,
    short_description varchar(255)   NOT NULL,
    image_url         varchar(500)   NOT NULL,
    slug              varchar(255)   NOT NULL UNIQUE,
    base_price        DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    sale_price        DECIMAL(10, 2)          DEFAULT NULL,
    is_featured       BOOLEAN                 DEFAULT false,
    sale_start_date   TIMESTAMP,
    sale_end_date     TIMESTAMP,
    status            varchar(255)   NOT NULL DEFAULT 'ACTIVE',
    brand_id          int8           NOT NULL,
    category_id       int8           NOT NULL,
    user_id           int8           NOT NULL,
    created_by        varchar(255)   NOT NULL,
    updated_by        varchar(255)   NOT NULL,
    created_at        timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at        timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    version           int4           NOT NULL DEFAULT 1,
    PRIMARY KEY (id)
);

CREATE TABLE product_attribute_info
(
    id              BIGSERIAL PRIMARY KEY,
    attribute_name  VARCHAR(255) NOT NULL,
    attribute_value VARCHAR(255) NOT NULL,
    product_id      int8         NOT NULL
);

CREATE TABLE product_option_info
(
    id           BIGSERIAL PRIMARY KEY,
    option_name  VARCHAR(255) NOT NULL,
    option_value VARCHAR(255) NOT NULL,
    product_id   int8         NOT NULL
);



CREATE TABLE product_attribute
(
    id         BIGSERIAL    NOT NULL,
    name       varchar(255) NOT NULL,
    created_by varchar(255) NOT NULL,
    updated_by varchar(255) NOT NULL,
    created_at timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE product_attribute_category
(
    id                   BIGSERIAL NOT NULL,
    product_attribute_id int8      NOT NULL,
    category_id          int8      NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE product_attribute_value
(
    id                   BIGSERIAL    NOT NULL,
    value                varchar(255) NOT NULL,
    product_attribute_id int8         NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE product_option
(
    id         BIGSERIAL    NOT NULL,
    name       varchar(255) NOT NULL,
    created_by varchar(255) NOT NULL,
    updated_by varchar(255) NOT NULL,
    created_at timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE product_option_combination
(
    ID                BIGSERIAL NOT NULL,
    product_id        int8      NOT NULL,
    product_option_id int8      NOT NULL,
    display_order     int4      NOT NULL,
    PRIMARY KEY (ID)
);

CREATE TABLE product_option_value
(
    id                BIGSERIAL    NOT NULL,
    value             varchar(255) NOT NULL,
    product_option_id int8         NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE product_product_attribute_value
(
    ID                         BIGSERIAL NOT NULL,
    product_id                 int8      NOT NULL,
    product_attribute_value_id int8      NOT NULL,
    display_order              int8      NOT NULL,
    PRIMARY KEY (ID)
);

CREATE TABLE product_sku
(
    id              BIGSERIAL      NOT NULL,
    sku             varchar(255)   NOT NULL,
    sku_signature   varchar(255)   NOT NULL UNIQUE,
    extra_price     DECIMAL(10, 2) NOT NULL DEFAULT 0, -- Additional price for the SKU
    sale_type       VARCHAR(20)             DEFAULT NULL, -- "Percentage" or "Fixed"
    sale_value      DECIMAL(10, 2)          DEFAULT NULL,
    sale_start_date TIMESTAMP               DEFAULT NULL,
    sale_end_date   TIMESTAMP               DEFAULT NULL,
    status          varchar(255)   NOT NULL DEFAULT 'ACTIVE',
    product_id      int8           NOT NULL,
    created_by      varchar(255)   NOT NULL,
    updated_by      varchar(255)   NOT NULL,
    created_at      timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    version         int4           NOT NULL DEFAULT 1,
    PRIMARY KEY (id),

    CONSTRAINT CHKproduct_sku_price CHECK (extra_price >= 0)
);

CREATE TABLE product_sku_value
(
    id                      BIGSERIAL NOT NULL,
    productSKU_id           int8      NOT NULL,
    product_option_value_id int8      NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE product_inventory
(
    id              BIGSERIAL PRIMARY KEY,
    product_id      BIGINT       NOT NULL,
    product_sku_id  BIGINT       NOT NULL,
    available_stock INTEGER      NOT NULL DEFAULT 0,
    reserved_stock  INTEGER      NOT NULL DEFAULT 0, -- For pending orders
    damaged_stock   INTEGER      NOT NULL DEFAULT 0,
    total_stock     INTEGER      NOT NULL DEFAULT 0,
    created_by      VARCHAR(255) NOT NULL,
    updated_by      VARCHAR(255) NOT NULL,
    created_at      TIMESTAMP             DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP             DEFAULT CURRENT_TIMESTAMP,
    version         int4         NOT NULL DEFAULT 1,

    CONSTRAINT CHK_stock_values CHECK (available_stock >= 0 AND reserved_stock >= 0 AND damaged_stock >= 0),
    CONSTRAINT CHK_inventory_reference CHECK (product_id IS NOT NULL OR product_sku_id IS NOT NULL)
);

CREATE TABLE product_review
(
    id                   BIGSERIAL PRIMARY KEY,
    product_id           BIGINT       NOT NULL,
    user_id              BIGINT       NOT NULL,
    rating               INTEGER      NOT NULL,
    title                VARCHAR(255),
    review_text          TEXT,
    is_verified_purchase BOOLEAN   DEFAULT false,
    reviewer_name        VARCHAR(255),
    reviewer_email       VARCHAR(255),
    created_by           VARCHAR(255) NOT NULL,
    updated_by           VARCHAR(255) NOT NULL,
    created_at           TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at           TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT CHK_rating_range CHECK (rating >= 1 AND rating <= 5)
);

ALTER TABLE product
    ADD CONSTRAINT FKproduct822402 FOREIGN KEY (category_id) REFERENCES category (id) ON DELETE SET NULL;
ALTER TABLE product
    ADD CONSTRAINT FKproduct_123456 FOREIGN KEY (brand_id) REFERENCES brand (id) ON DELETE SET NULL;
ALTER TABLE product_attribute_info
    ADD CONSTRAINT FK_product_attribute_info_product FOREIGN KEY (product_id) REFERENCES product (id) ON DELETE CASCADE;
ALTER TABLE product_option_info
    ADD CONSTRAINT FK_product_option_info_product FOREIGN KEY (product_id) REFERENCES product (id) ON DELETE CASCADE;
ALTER TABLE category
    ADD CONSTRAINT FKcategory465764 FOREIGN KEY (parent_category_id) REFERENCES category (id) ON DELETE SET NULL;
ALTER TABLE product_attribute_value
    ADD CONSTRAINT FKproduct_at463715 FOREIGN KEY (product_attribute_id) REFERENCES product_attribute (id) ON DELETE CASCADE;
ALTER TABLE product_option_combination
    ADD CONSTRAINT FKproduct_op57306 FOREIGN KEY (product_option_id) REFERENCES product_option (id) ON DELETE CASCADE;
ALTER TABLE product_sku
    ADD CONSTRAINT FKproduct_sk755757 FOREIGN KEY (product_id) REFERENCES product (id) ON DELETE CASCADE;
ALTER TABLE product_inventory
    ADD CONSTRAINT FKproduct_invent123456 FOREIGN KEY (product_id) REFERENCES product (id) ON DELETE CASCADE;
ALTER TABLE product_inventory
    ADD CONSTRAINT FKproduct_invent789012 FOREIGN KEY (product_sku_id) REFERENCES product_sku (id) ON DELETE CASCADE;
ALTER TABLE product_sku_value
    ADD CONSTRAINT FKproduct_sk119171 FOREIGN KEY (productSKU_id) REFERENCES product_sku (id) ON DELETE SET NULL;
ALTER TABLE product_sku_value
    ADD CONSTRAINT FKproduct_sk81262 FOREIGN KEY (product_option_value_id) REFERENCES product_option_value (id) ON DELETE SET NULL;
ALTER TABLE product_attribute_category
    ADD CONSTRAINT FKproduct_at802748 FOREIGN KEY (product_attribute_id) REFERENCES product_attribute (id);
ALTER TABLE product_attribute_category
    ADD CONSTRAINT FKproduct_at169936 FOREIGN KEY (category_id) REFERENCES category (id) ON DELETE SET NULL;
ALTER TABLE product_option_value
    ADD CONSTRAINT FKproduct_op735050 FOREIGN KEY (product_option_id) REFERENCES product_option (id) ON DELETE SET NULL;
ALTER TABLE product_option_combination
    ADD CONSTRAINT FKproduct_op606934 FOREIGN KEY (product_id) REFERENCES product (id) ON DELETE CASCADE;
ALTER TABLE product_product_attribute_value
    ADD CONSTRAINT FKproduct_pr376641 FOREIGN KEY (product_id) REFERENCES product (id) ON DELETE CASCADE;
ALTER TABLE product_product_attribute_value
    ADD CONSTRAINT FKproduct_pr98112 FOREIGN KEY (product_attribute_value_id) REFERENCES product_attribute_value (id) ON DELETE CASCADE;
AlTER TABLE product_review
    ADD CONSTRAINT FKproduct_re123456 FOREIGN KEY (product_id) REFERENCES product (id) ON DELETE CASCADE;