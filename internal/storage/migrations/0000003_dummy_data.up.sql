INSERT INTO Product (id, title, price, description, update_date, images) 
VALUES
(1, 'Trousers', 10000, 'The best trousers on the planet', CURRENT_TIMESTAMP(), 'static/trouesrs.png'),
(2, 'T-Shirt', 5000, 'Cool t-shirt you are likely to buy', CURRENT_TIMESTAMP(), 'static/tshirt.png')

INSERT INTO Review (id, product_id, user_id, score, text, pros, cons)
VALUES
(1, 1, 1, 5, 'This is literally the best trousers on the planet Earth!', 'So much stretch', 'None!')
