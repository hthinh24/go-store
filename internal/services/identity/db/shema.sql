CREATE TABLE IF NOT EXISTS users
(
    id            BIGSERIAL    NOT NULL,
    email         varchar(255) NOT NULL UNIQUE,
    password      varchar(255),
    provider_id   varchar(255) NOT NULL DEFAULT 1,
    provider_name varchar(255) NOT NULL DEFAULT 'APP',
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
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS permissions
(
    id          BIGSERIAL    NOT NULL,
    name        varchar(255) NOT NULL,
    description varchar(255) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS user_roles
(
    user_id int8 NOT NULL,
    role_id int8 NOT NULL,
    PRIMARY KEY (user_id, role_id)
);

CREATE TABLE IF NOT EXISTS role_permissions
(
    role_id       int8 NOT NULL,
    permission_id int8 NOT NULL,
    PRIMARY KEY (role_id, permission_id)
);

ALTER TABLE user_roles
    ADD CONSTRAINT FKuser_has_r352169 FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE;
ALTER TABLE user_roles
    ADD CONSTRAINT FKuser_has_r355910 FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE CASCADE;
ALTER TABLE role_permissions
    ADD CONSTRAINT FKrole_has_p648170 FOREIGN KEY (permission_id) REFERENCES permissions (id) ON DELETE SET NULL;
ALTER TABLE role_permissions
    ADD CONSTRAINT FKrole_has_p131704 FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE SET NULL;