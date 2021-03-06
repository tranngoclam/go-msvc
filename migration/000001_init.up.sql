CREATE SCHEMA IF NOT EXISTS todo;

USE todo;

CREATE TABLE IF NOT EXISTS tasks
(
    id          BIGINT(10) AUTO_INCREMENT,
    description VARCHAR(128) NOT NULL,
    priority    VARCHAR(9)   NOT NULL DEFAULT '' NOT NULL,
    status      VARCHAR(15)  NOT NULL DEFAULT '' NOT NULL,
    created_at  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);
