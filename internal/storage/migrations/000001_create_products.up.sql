CREATE TABLE Product
(
                         id BIGINT UNSIGNED PRIMARY KEY,
                         title VARCHAR(1024),
                         price BIGINT UNSIGNED,
                         description MEDIUMTEXT,
                         update_date TIMESTAMP,
                         images MEDIUMTEXT
);