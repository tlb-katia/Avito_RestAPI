package openapi

import (
	"context"
	"log"
)

func InitDataBase(ctx context.Context, pg *Postgres) error {
	initSQL := `
	DO $$
	BEGIN
    	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'tender_status') THEN
        	CREATE TYPE tender_status AS ENUM ('Created', 'Published', 'Closed');
    	END IF;
    
    	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'tender_service_type') THEN
        	CREATE TYPE tender_service_type AS ENUM ('Construction', 'Delivery', 'Manufacture');
    	END IF;
	END $$;

	CREATE TABLE IF NOT EXISTS tenders (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		name VARCHAR(255) NOT NULL,
		description TEXT,
		service_type tender_service_type NOT NULL,
		status VARCHAR(20),
		organization_id UUID REFERENCES organization(id) ON DELETE CASCADE,
		creator_username VARCHAR(50) NOT NULL,
		version INT NOT NULL DEFAULT 1,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS tender_versions (
		id SERIAL PRIMARY KEY,
		tender_id UUID REFERENCES tenders(id) ON DELETE CASCADE,
		name VARCHAR(255) NOT NULL,
		description TEXT,
		service_type tender_service_type NOT NULL,
		status VARCHAR(20),
		organization_id UUID REFERENCES organization(id) ON DELETE CASCADE,
		creator_username VARCHAR(50) NOT NULL,
		version INT NOT NULL DEFAULT 1,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS bids (
		bid_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		name VARCHAR(100) NOT NULL,
		description TEXT,
		status VARCHAR(20),
		tender_id UUID REFERENCES tenders(id) ON DELETE CASCADE,
		author_type VARCHAR(20) NOT NULL,
		author_id UUID NOT NULL,
		version INT DEFAULT 1 CHECK (version >= 1),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS bids_versions (
    	bid_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    	name VARCHAR(100) NOT NULL,
    	description TEXT,
    	status VARCHAR(20),
    	tender_id UUID NOT NULL NOT NULL,
    	author_type VARCHAR(20) NOT NULL,
    	author_id UUID NOT NULL,
    	version INT DEFAULT 1 CHECK (version >= 1),
    	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS bid_feedback (
    	feedback_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    	bid_id UUID REFERENCES bids(bid_id) ON DELETE CASCADE,
    	feedback TEXT NOT NULL,
    	username VARCHAR(50) REFERENCES employee(username) ON DELETE CASCADE,
    	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS bid_decisions (
    	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    	bid_id UUID REFERENCES bids(bid_id) ON DELETE CASCADE,
    	decision VARCHAR(20) NOT NULL,
    	decided_by UUID REFERENCES employee(id) ON DELETE SET NULL,
    	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	INSERT INTO organization (id, name, description, type) VALUES
   		(uuid_generate_v4(), 'Organization 1', 'This is the first organization.', 'LLC'),
    	(uuid_generate_v4(), 'Organization 2', 'This is the second organization.', 'IE'),
    	(uuid_generate_v4(), 'Organization 3', 'This is the third organization.', 'JSC');

	INSERT INTO employee (id, username, first_name, last_name) VALUES
    	(uuid_generate_v4(), 'user1', 'John', 'Doe'),
    	(uuid_generate_v4(), 'user2', 'Jane', 'Smith'),
    	(uuid_generate_v4(), 'user3', 'Alice', 'Johnson');

	INSERT INTO organization_responsible (id, organization_id, user_id) VALUES
     	(uuid_generate_v4(), (SELECT id FROM organization WHERE name = 'Organization 1'), (SELECT id FROM employee WHERE username = 'user1')),
     	(uuid_generate_v4(), (SELECT id FROM organization WHERE name = 'Organization 2'), (SELECT id FROM employee WHERE username = 'user2')),
     	(uuid_generate_v4(), (SELECT id FROM organization WHERE name = 'Organization 3'), (SELECT id FROM employee WHERE username = 'user3'));

	INSERT INTO tenders (id, name, description, service_type, status, organization_id, creator_username) VALUES
    	(uuid_generate_v4(), 'Tender 1', 'Description for Tender 1', 'Construction', 'Created', (SELECT id FROM organization WHERE name = 'Organization 1'), 'user1'),
    	(uuid_generate_v4(), 'Tender 2', 'Description for Tender 2', 'Delivery', 'Published', (SELECT id FROM organization WHERE name = 'Organization 2'), 'user2'),
    	(uuid_generate_v4(), 'Tender 3', 'Description for Tender 3', 'Manufacture', 'Closed', (SELECT id FROM organization WHERE name = 'Organization 3'), 'user3');

	INSERT INTO tender_versions (tender_id, name, description, service_type, status, organization_id, creator_username, version) VALUES
    	((SELECT id FROM tenders WHERE name = 'Tender 1' LIMIT 1), 'Tender 1 Version 1', 'Version 1 of Tender 1', 'Construction', 'Created', (SELECT id FROM organization WHERE name = 'Organization 1'), 'user1', 1),
    	((SELECT id FROM tenders WHERE name = 'Tender 2' LIMIT 1), 'Tender 2 Version 1', 'Version 1 of Tender 2', 'Delivery', 'Published', (SELECT id FROM organization WHERE name = 'Organization 2'), 'user2', 1),
    	((SELECT id FROM tenders WHERE name = 'Tender 3' LIMIT 1), 'Tender 3 Version 1', 'Version 1 of Tender 3', 'Manufacture', 'Closed', (SELECT id FROM organization WHERE name = 'Organization 3'), 'user3', 1);

	INSERT INTO bids (bid_id, name, description, status, tender_id, author_type, author_id)
	VALUES
    	(uuid_generate_v4(), 'Bid 1', 'Description for Bid 1', 'Created',
    	(SELECT id FROM tenders WHERE name = 'Tender 1' LIMIT 1),
    	'User',
    	(SELECT id FROM employee WHERE username = 'user1' LIMIT 1)),

    	(uuid_generate_v4(), 'Bid 2', 'Description for Bid 2', 'Published',
    	(SELECT id FROM tenders WHERE name = 'Tender 2' LIMIT 1),
    	'User',
    	(SELECT id FROM employee WHERE username = 'user2' LIMIT 1)),

    	(uuid_generate_v4(), 'Bid 3', 'Description for Bid 3', 'Closed',
    	(SELECT id FROM tenders WHERE name = 'Tender 3' LIMIT 1),
    	'User',
    	(SELECT id FROM employee WHERE username = 'user3' LIMIT 1));
	`

	_, err := pg.Pool.Exec(ctx, initSQL)
	if err != nil {
		log.Printf("Error initializing database schema: %v\n", err)
		return err
	}

	log.Println("Database initialized successfully.")
	return nil
}
