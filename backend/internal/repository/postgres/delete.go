package postgres

import (
	"context"
	"fmt"
)

// Delete removes a user by their ID.
func (r *UserRepository) Delete(ctx context.Context, id uint64) error {
	query := `
		UPDATE users
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL;
	`

	_, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}
