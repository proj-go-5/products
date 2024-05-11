BEGIN;
INSERT INTO Product (id, title, price, description, update_date, images) 
VALUES
(1, 'Trousers', 10000, 'The best trousers on the planet', CURRENT_TIMESTAMP(), 'static/trouesrs.png'),
(2, 'T-Shirt', 5000, 'Cool t-shirt you are likely to buy', CURRENT_TIMESTAMP(), 'static/tshirt.png');
