package enums

import "fmt"

func (a *analytics) handleEmailBounce(em email) error {
	if ok := em.recipient.updateStatus(em.status); ok != nil {
		return fmt.Errorf("error updating user status: %w", ok)
	} else if ok := a.track(em.status); ok != nil {
		return fmt.Errorf("error tracknig user bounce: %w", ok)
	}
	return nil
}
