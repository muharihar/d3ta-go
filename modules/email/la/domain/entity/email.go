package entity

// Email type
type Email struct {
	ID uint64 `json:"ID" gorm:"primary_key;column:id"`

	UUID       string        `json:"uuid" gorm:"column:uuid;size:255;unique;not null" sql:"index"`
	TemplateID uint64        `json:"templateID" gorm:"column:template_id;" sql:"index"`
	Template   EmailTemplate `json:"user" gorm:"ForeignKey:TemplateID;AssociationForeignKey:ID;"`

	From     string `json:"from" gorm:"column:from;size:255"`
	FromName string `json:"fromName" gorm:"column:from_name;size:255"`
	To       string `json:"to" gorm:"column:to;size:255"`
	ToName   string `json:"toName" gorm:"column:to_name;size:255"`
	CC       string `json:"cc" gorm:"column:cc;size:500"`
	BCC      string `json:"bcc" gorm:"column:bcc;size:500"`
	Subject  string `json:"subject" gorm:"column:subject;size:500"`
	Body     []byte `json:"body" gorm:"column:body;"`
	Status   string `json:"status" gorm:"column:status;size:100"`

	SentBy string `json:"sentBy" gorm:"column:sent_by;size:255"`

	BaseEntity
}

// TableName get real database table name
func (t Email) TableName() string {
	return "eml_emails"
}
