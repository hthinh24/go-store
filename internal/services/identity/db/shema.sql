CREATE TABLE IF NOT EXISTS users
(
    id            BIGSERIAL    NOT NULL,
    email         varchar(255) NOT NULL UNIQUE,
    password      varchar(255),
    provider_id    varchar(255) NOT NULL DEFAULT 1,
    provider_name  varchar(255) NOT NULL DEFAULT 'APP',
    last_name     varchar(255) NOT NULL,
    first_name    varchar(255) NOT NULL,
    avatar        varchar(255),
    gender        varchar(255) NOT NULL,
    phone_number  varchar(20)  NOT NULL,
    date_of_birth timestamp    NOT NULL,
    status        varchar(255) NOT NULL,
    created_by    varchar(255) NOT NULL,
    updated_by    varchar(255) NOT NULL,
    created_at    timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at    timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    version       int          NOT NULL DEFAULT 1,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS roles
(
    id          BIGSERIAL    NOT NULL,
    name        varchar(255) NOT NULL,
    description varchar(255) NOT NULL,
    created_by  varchar(255) NOT NULL,
    updated_by  varchar(255) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS permissions
(
    id          BIGSERIAL    NOT NULL,
    name        varchar(255) NOT NULL,
    description varchar(255) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS user_has_role
(
    ID       BIGSERIAL NOT NULL,
    users_id int8      NOT NULL,
    roles_id int8      NOT NULL,
    PRIMARY KEY (ID)
);

CREATE TABLE IF NOT EXISTS role_has_permission
(
    ID             BIGSERIAL NOT NULL,
    permissions_id int8      NOT NULL,
    roles_id       int8      NOT NULL,
    PRIMARY KEY (ID)
);

ALTER TABLE user_has_role
    ADD CONSTRAINT FKuser_has_r352169 FOREIGN KEY (users_id) REFERENCES users (id);
ALTER TABLE user_has_role
    ADD CONSTRAINT FKuser_has_r355910 FOREIGN KEY (roles_id) REFERENCES roles (id);
ALTER TABLE role_has_permission
    ADD CONSTRAINT FKrole_has_p648170 FOREIGN KEY (permissions_id) REFERENCES permissions (id);
ALTER TABLE role_has_permission
    ADD CONSTRAINT FKrole_has_p131704 FOREIGN KEY (roles_id) REFERENCES roles (id);