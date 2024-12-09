ALTER Table history ADD COLUMN product_link JSON after recommendation;

ALTER Table history ADD COLUMN user_picture VARCHAR(255) after product_link;
