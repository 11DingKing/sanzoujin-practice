package service

import (
	"context"
	"github.com/11DingKing/sanzoujin-practice/internal/domain"
	"github.com/11DingKing/sanzoujin-practice/internal/repository"
	"sort"
	"strings"
	"time"
)

type MatchingService struct {
	Enrollments repository.EnrollmentRepo
	Groups      repository.GroupRepo
}
type MatchCandidate struct {
	Enrollment domain.Enrollment
	Score      int
	Reasons    []string
}

func (m MatchingService) Rank(ctx context.Context, projectID string, required []string) ([]MatchCandidate, error) {
	items, err := m.Enrollments.ListByProject(ctx, projectID)
	if err != nil {
		return nil, err
	}
	out := make([]MatchCandidate, 0, len(items))
	for _, e := range items {
		if e.Status != domain.EnrollmentAuthorized {
			continue
		}
		score, reasons := qualificationScore(e.Qualifications, required)
		out = append(out, MatchCandidate{Enrollment: e, Score: score, Reasons: reasons})
	}
	sort.SliceStable(out, func(i, j int) bool {
		if out[i].Score == out[j].Score {
			return out[i].Enrollment.CreatedAt.Before(out[j].Enrollment.CreatedAt)
		}
		return out[i].Score > out[j].Score
	})
	return out, nil
}
func qualificationScore(have, required []string) (int, []string) {
	set := map[string]bool{}
	for _, v := range have {
		set[strings.ToLower(strings.TrimSpace(v))] = true
	}
	score := 0
	reasons := []string{}
	for _, v := range required {
		key := strings.ToLower(strings.TrimSpace(v))
		if set[key] {
			score++
			reasons = append(reasons, "具备"+v)
		} else {
			reasons = append(reasons, "缺少"+v)
		}
	}
	return score, reasons
}
func eligibleAt(now, start, end time.Time, skills, required []string) bool {
	if !domain.WithinWindow(now, start, end) {
		return false
	}
	score, _ := qualificationScore(skills, required)
	return score == len(required)
}
