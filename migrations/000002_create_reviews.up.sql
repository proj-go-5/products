CREATE TABLE Review
(
    id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    product_id BIGINT UNSIGNED,
    user_id BIGINT UNSIGNED,
    score SMALLINT(1),
    text MEDIUMTEXT,
    pros MEDIUMTEXT,
    cons MEDIUMTEXT,
    FOREIGN KEY (product_id) REFERENCES Product(id)
);