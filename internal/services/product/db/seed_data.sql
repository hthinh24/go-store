-- Common seed data for e-commerce platform
-- Categories, Brands, Product Attributes, and Product Options

-- =============================================
-- CATEGORIES
-- =============================================
INSERT INTO category (name, description, parent_category_id) VALUES
-- Top Level Categories
('Electronics', 'Electronic devices and gadgets', NULL),
('Clothing & Fashion', 'Apparel, shoes, and fashion accessories', NULL),
('Home & Garden', 'Home improvement, furniture, and garden supplies', NULL),
('Sports & Outdoors', 'Sports equipment and outdoor gear', NULL),
('Books & Media', 'Books, movies, music, and games', NULL),
('Health & Beauty', 'Health products, cosmetics, and personal care', NULL),
('Automotive', 'Car parts, accessories, and tools', NULL),
('Toys & Games', 'Toys, board games, and educational products', NULL),

-- Electronics Subcategories
('Computers & Laptops', 'Desktop computers, laptops, and accessories', 1),
('Mobile Phones', 'Smartphones and mobile accessories', 1),
('Audio & Video', 'Headphones, speakers, TVs, and audio equipment', 1),
('Gaming', 'Gaming consoles, games, and accessories', 1),
('Cameras & Photography', 'Digital cameras, lenses, and photography equipment', 1),
('Smart Home', 'IoT devices, smart speakers, and home automation', 1),

-- Clothing Subcategories
('Men''s Clothing', 'Men''s apparel and accessories', 2),
('Women''s Clothing', 'Women''s apparel and accessories', 2),
('Children''s Clothing', 'Kids and baby clothing', 2),
('Shoes', 'Footwear for all ages', 2),
('Accessories', 'Bags, jewelry, watches, and fashion accessories', 2),

-- Home & Garden Subcategories
('Furniture', 'Home and office furniture', 3),
('Kitchen & Dining', 'Cookware, appliances, and dining accessories', 3),
('Home Decor', 'Decorative items and home styling', 3),
('Garden & Outdoor', 'Gardening tools, plants, and outdoor furniture', 3),

-- Sports Subcategories
('Fitness & Exercise', 'Gym equipment and fitness accessories', 4),
('Outdoor Recreation', 'Camping, hiking, and outdoor activities', 4),
('Team Sports', 'Equipment for team sports', 4),
('Water Sports', 'Swimming, surfing, and water activities', 4);

-- =============================================
-- BRANDS
-- =============================================
INSERT INTO brand (name, description, created_by, updated_by) VALUES
-- Electronics Brands
('Apple', 'Premium consumer electronics and software', 'system', 'system'),
('Samsung', 'South Korean multinational electronics company', 'system', 'system'),
('Sony', 'Japanese electronics and entertainment company', 'system', 'system'),
('Microsoft', 'Technology corporation and software developer', 'system', 'system'),
('Google', 'Technology company specializing in internet services', 'system', 'system'),
('Dell', 'Computer technology company', 'system', 'system'),
('HP', 'Hewlett-Packard technology company', 'system', 'system'),
('Lenovo', 'Chinese multinational technology company', 'system', 'system'),
('ASUS', 'Taiwanese multinational computer hardware company', 'system', 'system'),
('LG', 'South Korean electronics company', 'system', 'system'),

-- Fashion Brands
('Nike', 'Athletic footwear and apparel', 'system', 'system'),
('Adidas', 'German multinational sportswear corporation', 'system', 'system'),
('Zara', 'Spanish fast fashion retailer', 'system', 'system'),
('H&M', 'Swedish multinational clothing retailer', 'system', 'system'),
('Uniqlo', 'Japanese casual wear designer and retailer', 'system', 'system'),
('Levi''s', 'American clothing company known for denim', 'system', 'system'),

-- Home & Kitchen Brands
('IKEA', 'Swedish furniture retailer', 'system', 'system'),
('KitchenAid', 'Kitchen appliance brand', 'system', 'system'),
('Dyson', 'British technology company known for vacuums', 'system', 'system'),

-- Generic/Store Brands
('Generic', 'Generic or unbranded products', 'system', 'system'),
('Store Brand', 'Private label store brand', 'system', 'system');

-- =============================================
-- PRODUCT ATTRIBUTES
-- =============================================
INSERT INTO product_attribute (name, created_by, updated_by) VALUES
-- Electronics Attributes
('Processor', 'system', 'system'),                    -- 1
('RAM Memory', 'system', 'system'),                   -- 2
('Storage Capacity', 'system', 'system'),             -- 3
('Graphics Card', 'system', 'system'),                -- 4
('Screen Size', 'system', 'system'),                  -- 5
('Operating System', 'system', 'system'),             -- 6
('Battery Life', 'system', 'system'),                 -- 7
('Camera Resolution', 'system', 'system'),            -- 8
('Connectivity', 'system', 'system'),                 -- 9
('Display Type', 'system', 'system'),                 -- 10

-- Clothing Attributes
('Size', 'system', 'system'),                         -- 11
('Material', 'system', 'system'),                     -- 12
('Fit Type', 'system', 'system'),                     -- 13
('Sleeve Length', 'system', 'system'),                -- 14
('Neckline', 'system', 'system'),                     -- 15
('Pattern', 'system', 'system'),                      -- 16

-- General Attributes
('Color', 'system', 'system'),                        -- 17
('Weight', 'system', 'system'),                       -- 18
('Dimensions', 'system', 'system'),                   -- 19
('Power Source', 'system', 'system'),                 -- 20
('Warranty Period', 'system', 'system'),              -- 21
('Country of Origin', 'system', 'system'),            -- 22
('Energy Rating', 'system', 'system'),                -- 23
('Water Resistance', 'system', 'system'),             -- 24
('Capacity', 'system', 'system');                     -- 25

-- =============================================
-- PRODUCT ATTRIBUTE VALUES
-- =============================================
INSERT INTO product_attribute_value (value, product_attribute_id) VALUES
-- Processor Values (1)
('Intel Core i3', 1), ('Intel Core i5', 1), ('Intel Core i7', 1), ('Intel Core i9', 1),
('AMD Ryzen 3', 1), ('AMD Ryzen 5', 1), ('AMD Ryzen 7', 1), ('AMD Ryzen 9', 1),
('Apple M1', 1), ('Apple M2', 1), ('Apple M3', 1),

-- RAM Memory Values (2)
('4GB', 2), ('8GB', 2), ('16GB', 2), ('32GB', 2), ('64GB', 2),
('4GB DDR4', 2), ('8GB DDR4', 2), ('16GB DDR4', 2), ('32GB DDR4', 2),
('8GB DDR5', 2), ('16GB DDR5', 2), ('32GB DDR5', 2), ('64GB DDR5', 2),

-- Storage Capacity Values (3)
('128GB SSD', 3), ('256GB SSD', 3), ('512GB SSD', 3), ('1TB SSD', 3), ('2TB SSD', 3),
('500GB HDD', 3), ('1TB HDD', 3), ('2TB HDD', 3), ('4TB HDD', 3),
('128GB', 3), ('256GB', 3), ('512GB', 3), ('1TB', 3),

-- Graphics Card Values (4)
('Integrated Graphics', 4), ('NVIDIA GTX 1650', 4), ('NVIDIA RTX 3060', 4), ('NVIDIA RTX 4070', 4), ('NVIDIA RTX 4080', 4),
('AMD Radeon RX 6600', 4), ('AMD Radeon RX 7800', 4), ('Apple GPU', 4),

-- Screen Size Values (5)
('13.3 inch', 5), ('14 inch', 5), ('15.6 inch', 5), ('17.3 inch', 5),
('24 inch', 5), ('27 inch', 5), ('32 inch', 5),
('6.1 inch', 5), ('6.7 inch', 5), ('6.9 inch', 5),

-- Operating System Values (6)
('Windows 11 Home', 6), ('Windows 11 Pro', 6), ('macOS', 6), ('Linux Ubuntu', 6),
('Android 14', 6), ('iOS 17', 6), ('Chrome OS', 6),

-- Size Values (11)
('XS', 11), ('S', 11), ('M', 11), ('L', 11), ('XL', 11), ('XXL', 11), ('XXXL', 11),
('28', 11), ('30', 11), ('32', 11), ('34', 11), ('36', 11), ('38', 11), ('40', 11), ('42', 11),
('6', 11), ('7', 11), ('8', 11), ('9', 11), ('10', 11), ('11', 11), ('12', 11),

-- Material Values (12)
('Cotton', 12), ('Polyester', 12), ('Wool', 12), ('Silk', 12), ('Linen', 12), ('Denim', 12),
('Leather', 12), ('Synthetic', 12), ('Cotton Blend', 12), ('Bamboo', 12),

-- Color Values (17)
('Black', 17), ('White', 17), ('Gray', 17), ('Silver', 17), ('Gold', 17),
('Red', 17), ('Blue', 17), ('Green', 17), ('Yellow', 17), ('Orange', 17),
('Purple', 17), ('Pink', 17), ('Brown', 17), ('Navy', 17), ('Beige', 17),
('Rose Gold', 17), ('Space Gray', 17), ('Midnight', 17);

-- =============================================
-- PRODUCT OPTIONS
-- =============================================
INSERT INTO product_option (name, created_by, updated_by) VALUES
('Color', 'system', 'system'),                        -- 1
('Size', 'system', 'system'),                         -- 2
('Storage', 'system', 'system'),                      -- 3
('Memory', 'system', 'system'),                       -- 4
('Style', 'system', 'system'),                        -- 5
('Configuration', 'system', 'system'),                -- 6
('Edition', 'system', 'system'),                      -- 7
('Package', 'system', 'system'),                      -- 8
('Capacity', 'system', 'system'),                     -- 9
('Version', 'system', 'system');                      -- 10

-- =============================================
-- PRODUCT OPTION VALUES
-- =============================================
INSERT INTO product_option_value (value, product_option_id) VALUES
-- Color Options (1)
('Black', 1), ('White', 1), ('Gray', 1), ('Silver', 1), ('Gold', 1),
('Red', 1), ('Blue', 1), ('Green', 1), ('Navy', 1), ('Rose Gold', 1),
('Space Gray', 1), ('Midnight', 1), ('Purple', 1), ('Pink', 1),

-- Size Options (2)
('Small', 2), ('Medium', 2), ('Large', 2), ('Extra Large', 2),
('XS', 2), ('S', 2), ('M', 2), ('L', 2), ('XL', 2), ('XXL', 2),
('One Size', 2),

-- Storage Options (3)
('64GB', 3), ('128GB', 3), ('256GB', 3), ('512GB', 3), ('1TB', 3), ('2TB', 3),

-- Memory Options (4)
('8GB', 4), ('16GB', 4), ('32GB', 4), ('64GB', 4),

-- Style Options (5)
('Classic', 5), ('Modern', 5), ('Vintage', 5), ('Sport', 5), ('Casual', 5), ('Formal', 5),

-- Configuration Options (6)
('Basic', 6), ('Standard', 6), ('Premium', 6), ('Pro', 6), ('Ultimate', 6),

-- Edition Options (7)
('Standard Edition', 7), ('Special Edition', 7), ('Limited Edition', 7), ('Anniversary Edition', 7),

-- Package Options (8)
('Basic Package', 8), ('Standard Package', 8), ('Premium Package', 8), ('Bundle', 8),

-- Capacity Options (9)
('Small', 9), ('Medium', 9), ('Large', 9), ('Extra Large', 9),

-- Version Options (10)
('Version 1.0', 10), ('Version 2.0', 10), ('Latest Version', 10), ('Beta Version', 10);

-- =============================================
-- CATEGORY-ATTRIBUTE ASSOCIATIONS
-- =============================================
INSERT INTO product_attribute_category (product_attribute_id, category_id) VALUES
-- Electronics (1) - Computer attributes
(1, 1), (2, 1), (3, 1), (4, 1), (5, 1), (6, 1), (7, 1), (8, 1), (9, 1), (10, 1), (17, 1), (18, 1), (19, 1), (20, 1), (21, 1),

-- Computers & Laptops (9)
(1, 9), (2, 9), (3, 9), (4, 9), (5, 9), (6, 9), (7, 9), (17, 9), (18, 9), (21, 9),

-- Mobile Phones (10)
(3, 10), (5, 10), (6, 10), (7, 10), (8, 10), (9, 10), (17, 10), (18, 10), (21, 10), (24, 10),

-- Clothing & Fashion (2)
(11, 2), (12, 2), (13, 2), (14, 2), (15, 2), (16, 2), (17, 2), (22, 2),

-- Men's Clothing (15), Women's Clothing (16), Children's Clothing (17)
(11, 15), (12, 15), (13, 15), (14, 15), (15, 15), (16, 15), (17, 15),
(11, 16), (12, 16), (13, 16), (14, 16), (15, 16), (16, 16), (17, 16),
(11, 17), (12, 17), (17, 17),

-- Shoes (18)
(11, 18), (12, 18), (17, 18), (22, 18),

-- General attributes for all categories
(17, 3), (18, 3), (19, 3), (21, 3), (22, 3), -- Home & Garden
(17, 4), (18, 4), (19, 4), (21, 4), (22, 4); -- Sports & Outdoors
