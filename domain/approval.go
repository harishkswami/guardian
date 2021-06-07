package domain

import "time"

const (
	ApprovalStatusPending  = "pending"
	ApprovalStatusSkipped  = "skipped"
	ApprovalStatusApproved = "approved"
	ApprovalStatusRejected = "rejected"
)

type Approval struct {
	ID            uint    `json:"id"`
	Name          string  `json:"name"`
	Index         int     `json:"-"`
	AppealID      uint    `json:"appeal_id"`
	Status        string  `json:"status"`
	Actor         *string `json:"actor"`
	PolicyID      string  `json:"policy_id"`
	PolicyVersion uint    `json:"policy_version"`

	Approvers []string `json:"approvers,omitempty"`
	Appeal    *Appeal  `json:"appeal,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (a *Approval) IsHumanApproval() bool {
	return len(a.Approvers) > 0
}

type ApprovalRepository interface {
	BulkInsert([]*Approval) error
	GetPendingApprovals(user string) ([]*Approval, error)
}

type ApprovalService interface {
	BulkInsert([]*Approval) error
	GetPendingApprovals(user string) ([]*Approval, error)
	AdvanceApproval(appeal *Appeal) error
}
