package audit

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/11DingKing/sanzoujin-practice/internal/domain"
	"github.com/11DingKing/sanzoujin-practice/internal/repository"
	"time"
)

type Service struct{ Repo repository.AuditRepo }

func (s Service) Record(ctx context.Context, org, actor, obj, objID, action, result, requestID string) error {
	auditCtx := repository.AuditWriteContext(ctx)
	prev, e := s.Repo.LastHash(auditCtx, org)
	if e != nil {
		return e
	}
	payload, _ := json.Marshal([]string{org, actor, obj, objID, action, result, requestID, prev, time.Now().UTC().Format(time.RFC3339Nano)})
	sum := sha256.Sum256(payload)
	ev := domain.AuditEvent{ID: hex.EncodeToString(sum[:8]), OrgID: org, ActorID: actor, ObjectType: obj, ObjectID: objID, Action: action, Result: result, RequestID: requestID, PrevHash: prev, Hash: hex.EncodeToString(sum[:]), CreatedAt: time.Now()}
	return s.Repo.Append(auditCtx, ev)
}
