package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("coudn't delete users: %w", err)
	}
	fmt.Println("database reset successfully!")
	return nil
}
