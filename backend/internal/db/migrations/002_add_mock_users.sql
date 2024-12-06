-- Up migration
INSERT INTO users (
    user_name,
    first_name,
    last_name,
    email,
    user_status,
    department
) VALUES 
    ('johndoe', 'John', 'Doe', 'john.doe@example.com', 'A', 'IT'),
    ('janesmith', 'Jane', 'Smith', 'jane.smith@example.com', 'I', 'HR'),
    ('bobwilson', 'Bob', 'Wilson', 'bob.wilson@example.com', 'T', 'Sales');

-- Down migration
-- DELETE FROM users WHERE user_name IN ('johndoe', 'janesmith', 'bobwilson'); 