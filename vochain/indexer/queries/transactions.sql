-- name: CreateTxReference :execresult
INSERT INTO tx_references (
	hash, block_height, tx_block_index, tx_type
) VALUES (
	?, ?, ?, ?
);

-- name: GetTxReference :one
SELECT * FROM tx_references
WHERE id = ?
LIMIT 1;

-- name: GetTxReferenceByHash :one
SELECT * FROM tx_references
WHERE hash = ?
LIMIT 1;

-- name: GetLastTxReferences :many
SELECT * FROM tx_references
ORDER BY id DESC
LIMIT ?
OFFSET ?
;

-- name: CountTxReferences :one
SELECT COUNT(*) FROM tx_references;
