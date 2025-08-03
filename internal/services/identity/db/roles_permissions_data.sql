-- Insert Basic Permissions for Ecommerce Platform
INSERT INTO permissions (id, name, description) VALUES
-- User Management Permissions
(1, 'user.read', 'View user information'),
(2, 'user.update', 'Update user information'),
(3, 'user.list', 'List all users'),
(4, 'user.delete', 'Delete users'),

-- Product Management Permissions
(5, 'product.create', 'Create new products'),
(6, 'product.read', 'View product information'),
(7, 'product.update', 'Update product information'),
(8, 'product.delete', 'Delete products'),
(9, 'product.list', 'List all products'),

-- Order Management Permissions
(10, 'order.create', 'Create orders'),
(11, 'order.read', 'View order details'),
(12, 'order.update', 'Update order status'),
(13, 'order.list', 'List all orders'),

-- System Administration Permissions
(14, 'system.admin', 'Full system administration access'),
(15, 'role.manage', 'Manage roles and permissions');

-- Insert Roles for Ecommerce Platform
INSERT INTO roles VALUES
(1, 'admin', 'Administrator with full system access'),
(2, 'staff', 'Staff member with limited management access'),
(3, 'merchant', 'Merchant with full product management access'),
(4, 'user', 'Regular customer user');

-- Role Permission Assignments
-- Admin - All permissions
INSERT INTO role_permissions (role_id, permission_id) VALUES
(1, 14); -- System admin permission

-- Staff - Limited management permissions
INSERT INTO role_permissions (role_id, permission_id) VALUES
(2, 1), (2, 2), (2, 3),  -- User read/update/list
(2, 6), (2, 9),  -- Product read/list only
(2, 11), (2, 12), (2, 13);  -- Order read/update/list

-- Merchant - Full product management permissions
INSERT INTO role_permissions (role_id, permission_id) VALUES
(3, 1), (3, 2),  -- User read/update (own profile)
(3, 5), (3, 6), (3, 7), (3, 8), (3, 9),  -- Full product management
(3, 10), (3, 11), (3, 13);  -- Order create/read/list (for their products)

-- User - Basic customer permissions (product read only)
INSERT INTO role_permissions (role_id, permission_id) VALUES
(4, 1), (4, 2),  -- User read/update (own profile)
(4, 6),  -- Product read only
(4, 10), (4, 11);  -- Order create/read (own orders)
