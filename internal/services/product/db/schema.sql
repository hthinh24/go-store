CREATE TABLE category
(
    id                 SERIAL       NOT NULL,
    name               varchar(255) NOT NULL UNIQUE,
    description        varchar(255) NOT NULL,
    parent_category_id int4,
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
    id          BIGSERIAL    NOT NULL,
    name        varchar(255) NOT NULL,
    description varchar(255) NOT NULL,
    slug        varchar(255) NOT NULL UNIQUE,
    status      varchar(255) NOT NULL,
    category_id int4         NOT NULL,
    user_id     int8         NOT NULL,
    created_by  varchar(255) NOT NULL,
    updated_by  varchar(255) NOT NULL,
    created_at  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    version     int          NOT NULL DEFAULT 1,
    PRIMARY KEY (id)
);

CREATE TABLE product_attribute
(
    id         SERIAL       NOT NULL,
    name       varchar(255) NOT NULL,
    created_by varchar(255) NOT NULL,
    updated_by varchar(255) NOT NULL,
    created_at timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE product_attribute_category
(
    id                   SERIAL NOT NULL,
    product_attribute_id int4   NOT NULL,
    category_id          int4   NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE product_attribute_value
(
    id                   SERIAL       NOT NULL,
    value                varchar(255) NOT NULL,
    product_attribute_id int4         NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE product_option
(
    id         SERIAL       NOT NULL,
    name       varchar(255) NOT NULL,
    created_by varchar(255) NOT NULL,
    updated_by varchar(255) NOT NULL,
    created_at timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE product_option_combination
(
    ID                SERIAL NOT NULL,
    product_id        int8   NOT NULL,
    product_option_id int4   NOT NULL,
    display_order     int4   NOT NULL,
    PRIMARY KEY (ID)
);

CREATE TABLE product_option_value
(
    id                SERIAL       NOT NULL,
    value             varchar(255) NOT NULL,
    product_option_id int4         NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE product_product_attribute_value
(
    ID                         BIGSERIAL NOT NULL,
    product_id                 int8      NOT NULL,
    product_attribute_value_id int4      NOT NULL,
    display_order              int4      NOT NULL,
    PRIMARY KEY (ID)
);

CREATE TABLE product_sku
(
    id            BIGSERIAL      NOT NULL,
    sku           varchar(255)   NOT NULL,
    sku_signature varchar(255)   NOT NULL UNIQUE,
    stock         int4           NOT NULL DEFAULT 0,
    price         DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    product_id    int8           NOT NULL,
    created_by    varchar(255)   NOT NULL,
    updated_by    varchar(255)   NOT NULL,
    created_at    timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at    timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    version       int            NOT NULL DEFAULT 1,
    PRIMARY KEY (id)
);

CREATE TABLE product_sku_value
(
    id                      SERIAL NOT NULL,
    productSKU_id           int8   NOT NULL,
    product_option_value_id int4   NOT NULL,
    PRIMARY KEY (id)
);

ALTER TABLE product_sku
    ADD CONSTRAINT CHKproduct_sku_stock CHECK (stock >= 0),
    ADD CONSTRAINT CHKproduct_sku_price CHECK (price >= 0);

ALTER TABLE product
    ADD CONSTRAINT FKproduct822402 FOREIGN KEY (category_id) REFERENCES category (id) ON DELETE SET NULL;
ALTER TABLE category
    ADD CONSTRAINT FKcategory465764 FOREIGN KEY (parent_category_id) REFERENCES category (id) ON DELETE SET NULL;
ALTER TABLE product_attribute_value
    ADD CONSTRAINT FKproduct_at463715 FOREIGN KEY (product_attribute_id) REFERENCES product_attribute (id) ON DELETE CASCADE;
ALTER TABLE product_option_combination
    ADD CONSTRAINT FKproduct_op57306 FOREIGN KEY (product_option_id) REFERENCES product_option (id) ON DELETE CASCADE;
ALTER TABLE product_sku
    ADD CONSTRAINT FKproduct_sk755757 FOREIGN KEY (product_id) REFERENCES product (id) ON DELETE CASCADE;
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