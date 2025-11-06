-- Seed locations
INSERT INTO locations (name, area, description, capacity) VALUES
('Smart Solution', '1st Floor', 'Located on the 1st Floor', 8),
('Technology', '1st Floor', 'Located on the 1st Floor', 4),
('Integrity', '2nd Floor', 'Located on the 2nd Floor', 6),
('Innovation', '3rd Floor', 'Located on the 3rd Floor', 8),
('Loyalty', '3rd Floor', 'Located on the 3rd Floor', 6),
('Quality', '3rd Floor', 'Located on the 3rd Floor', 6),
('Team Work (Open Area)', '3rd Floor', 'Located on the 3rd Floor', 50),
('Excellent', '4th Floor', 'Located on the 4th Floor', 4),
('Open Communication', '4th Floor', 'Located on the 4th Floor', 8),
('General', 'Outside Floor', 'General area', 0)
ON CONFLICT DO NOTHING;

-- Seed admin user (password: admin123)
INSERT INTO users (name, email, password_hash, role) VALUES
('Admin User', 'admin@company.com', '$2a$10$ZKHDlvhleXejXrxGeqDlJOjacCdFeWPYsbfoRMFAGtHI3pwGUCX/S', 'admin')
ON CONFLICT (email) DO NOTHING;

-- Seed sample assets (do one by one to avoid PostgreSQL ON CONFLICT issues)
-- Seed sample assets (do one by one to avoid PostgreSQL ON CONFLICT issues)
INSERT INTO assets (unique_id, name, comment, detail, qty, brand, type, status, category, location_label) VALUES
('AST001', 'Dell Latitude 5420', 'Company laptop for developers', 'Intel i7, 16GB RAM, 512GB SSD', 10, 'Dell', 'it', 'available', 'Laptop', 'Technology');
INSERT INTO assets (unique_id, name, comment, detail, qty, brand, type, status, category, location_label) VALUES
('AST002', 'Dell UltraSharp U2720Q', '4K monitor for design team', '27-inch 4K USB-C Monitor', 15, 'Dell', 'it', 'available', 'Monitor', 'Smart Solution');
-- CORRECTED LINE: Changed 'in_use' to 'available'
INSERT INTO assets (unique_id, name, comment, detail, qty, brand, type, status, category, location_label) VALUES
('AST003', 'Logitech MX Master 3', 'Wireless mouse for senior developers', 'Wireless Bluetooth Mouse', 5, 'Logitech', 'it', 'available', 'Mouse', 'Surabaya Office - Lab 1');
INSERT INTO assets (unique_id, name, comment, detail, qty, brand, type, status, category, location_label) VALUES
('AST004', 'ThinkPad X1 Carbon Gen 11', 'Ultra-light business laptop', 'Intel i7-1365U, 16GB LPDDR5, 1TB SSD, 14" 2.8K', 8, 'Lenovo', 'it', 'available', 'Laptop', 'Integrity');
INSERT INTO assets (unique_id, name, comment, detail, qty, brand, type, status, category, location_label) VALUES
('AST005', 'Ergonomic Office Chair', 'Comfortable chair for long working hours', 'High-back chair with lumbar support', 20, 'Herman Miller', 'non_it', 'available', 'Furniture', 'Quality');
INSERT INTO assets (unique_id, name, comment, detail, qty, brand, type, status, category, location_label) VALUES
('AST006', 'Standing Desk Converter', 'Adjustable height desk converter', '32 inch wide, height adjustable', 12, 'Vari', 'non_it', 'booked', 'Furniture', 'Innovation');
INSERT INTO assets (unique_id, name, comment, detail, qty, brand, type, status, category, location_label) VALUES
('AST007', 'Cisco IP Phone 8851', 'VoIP phone for office communication', 'HD audio, 5-inch color display', 25, 'Cisco', 'it', 'available', 'Phone', 'Excellent');
INSERT INTO assets (unique_id, name, comment, detail, qty, brand, type, status, category, location_label) VALUES
('AST008', 'Brother HL-L2350DW', 'Compact monochrome laser printer', '2400 x 600 dpi, 32ppm', 5, 'Brother', 'it', 'broken', 'Printer', 'Team Work (Open Area)');

-- Seed sample tickets
INSERT INTO tickets (asset_id, category, severity, duration, due_date, reporting, comment) VALUES
((SELECT id FROM assets WHERE unique_id = 'AST002'), 'Repair', 'high', 24, NOW() + INTERVAL '24 hours', (SELECT id FROM users WHERE email = 'admin@company.com'), 'Monitor screen is flickering intermittently'),
((SELECT id FROM assets WHERE unique_id = 'AST008'), 'Repair', 'critical', 4, NOW() + INTERVAL '4 hours', (SELECT id FROM users WHERE email = 'admin@company.com'), 'Printer not feeding paper properly, jam issues'),
((SELECT id FROM assets WHERE unique_id = 'AST003'), 'Booking', 'low', 72, NOW() + INTERVAL '72 hours', (SELECT id FROM users WHERE email = 'admin@company.com'), 'Request to use wireless mouse for new developer'),
((SELECT id FROM assets WHERE unique_id = 'AST005'), 'Inquiry', 'medium', 48, NOW() + INTERVAL '48 hours', (SELECT id FROM users WHERE email = 'admin@company.com'), 'Information about chair maintenance and warranty');