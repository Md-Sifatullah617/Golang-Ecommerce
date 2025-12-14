INSERT INTO products (title, description, price, image_url) VALUES
('Wireless Headphones', 'High-quality noise-cancelling headphones.', 129.99, 'https://www.lovefoodhatewaste.com/sites/default/files/styles/twitter_card_image/public/2022-07/Citrus%20fruits.jpg.webp?itok=H1j9CCCS'),
('Smart Watch', 'Stylish smart watch with health tracking.', 199.99, 'https://i0.wp.com/post.healthline.com/wp-content/uploads/2021/05/apples-1296x728-header.jpg?w=1155&h=1528'),
('Running Shoes', 'Lightweight shoes for everyday running.', 89.50, 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRZbMOVB8a8wRQ6e-UKZggiu7-edRAN1GolPQ&s'),
('Laptop Stand', 'Ergonomic adjustable laptop stand.', 49.99, 'https://example.com/laptop-stand.jpg'),
('Bluetooth Speaker', 'Portable speaker with 12-hour battery life.', 79.99, 'https://example.com/speaker.jpg'),
('USB-C Cable', 'Fast charging 6ft USB-C cable.', 15.99, 'https://example.com/usb-cable.jpg')
ON CONFLICT DO NOTHING;