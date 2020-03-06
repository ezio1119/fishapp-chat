package domain

import "fmt"

func WrapOnRoomRepoErr(err error) error {
	return fmt.Errorf("room repository error: %w", err)
}

func WrapOnChatInterErr(err error) error {
	return fmt.Errorf("chat interactor error: %w", err)
}

func WrapOnChatPresenErr(err error) error {
	return fmt.Errorf("chat presenter error: %w", err)
}
