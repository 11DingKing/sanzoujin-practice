package domain

import "time"

func CanManageProject(role Role) bool { return role == RoleAdmin || role == RoleCoordinator }
func CanReview(role Role) bool {
	return role == RoleAdmin || role == RoleCoordinator || role == RoleMentor
}
func CanAttend(role Role) bool {
	return role == RoleStudent || role == RoleMentor || role == RoleCoordinator
}
func WithinWindow(now, start, end time.Time) bool { return !now.Before(start) && now.Before(end) }
func NormalizeScore(score int) (int, error) {
	if score < 1 || score > 5 {
		return 0, ErrInvalid
	}
	return score, nil
}
