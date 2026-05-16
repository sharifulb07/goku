-- Create resources table
CREATE TABLE IF NOT EXISTS resources (
    id              SERIAL PRIMARY KEY,
    name            VARCHAR(255) NOT NULL,
    role            VARCHAR(150),
    stack           TEXT,
    email           VARCHAR(255),
    status          VARCHAR(50) DEFAULT 'active',
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at      TIMESTAMP WITH TIME ZONE
);

CREATE INDEX IF NOT EXISTS idx_resources_name ON resources(name);
CREATE INDEX IF NOT EXISTS idx_resources_email ON resources(email);

SELECT 'Table resources created successfully!' as message;