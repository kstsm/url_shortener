package queries

const (
	CreateLink = `
WITH inserted AS (INSERT INTO links (original, shortened)
    VALUES ($1, $2)
    ON CONFLICT DO NOTHING
    RETURNING id, original, shortened)
SELECT id, original, shortened
FROM inserted
UNION ALL
SELECT id, original, shortened
FROM links
WHERE original = $1;`

	CreateTgLinks = `
INSERT INTO tg_links (link_id, tg_chat_id, title)
VALUES ($1, $2, $3);`
)
