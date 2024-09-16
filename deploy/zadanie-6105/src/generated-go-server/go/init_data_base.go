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
	`

	_, err := pg.Pool.Exec(ctx, initSQL)
	if err != nil {
		log.Printf("Error initializing database schema: %v\n", err)
		return err
	}

	log.Println("Database initialized successfully.")
	return nil
}
