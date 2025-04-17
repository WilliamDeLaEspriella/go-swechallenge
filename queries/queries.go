package queries

const GetRatingChangesBySearch = `
	SELECT id, ticker, company, brokerage, action, rating_from, rating_to, target_from, target_to
	FROM rating_changes
	WHERE ticker ILIKE $1
	   OR company ILIKE $1
	   OR brokerage ILIKE $1
	ORDER BY %s %s
	LIMIT $2 OFFSET $3;
`

const GetRatingChanges = `
	SELECT id, ticker, company, brokerage, action, rating_from, rating_to, target_from, target_to
	FROM rating_changes
	ORDER BY %s %s
	LIMIT $1 OFFSET $2;
`

const InsertRatingChange = `INSERT INTO rating_changes (
		    ticker, company, brokerage, action, rating_from, rating_to, target_from, target_to
		) VALUES (
		    $1, $2, $3, $4, $5, $6, $7, $8
		)`

const BestRatingChange = `SELECT
			  ticker,
			  company
			FROM rating_changes
			WHERE target_to > 0 AND target_from > 0
			GROUP BY ticker, company
			ORDER BY AVG(target_to - target_from) DESC, COUNT(*) DESC
			LIMIT 1;`

const CreateTables = `CREATE TABLE IF NOT EXISTS rating_changes (
			id SERIAL NOT NULL PRIMARY KEY,
			ticker VARCHAR(10) NOT NULL,
			company VARCHAR(100) NOT NULL,
			brokerage VARCHAR(100) NOT NULL,
			action VARCHAR(20) NOT NULL,
			rating_from VARCHAR(50),
			rating_to VARCHAR(50),
			target_from DECIMAL(10, 2),
			target_to DECIMAL(10, 2),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);`

const CountRatingChange = "SELECT COUNT(*) FROM rating_changes"
