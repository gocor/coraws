package coraws

import (
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var sharedSession *session.Session
var sharedSessionOnce sync.Once

// NewAwsSession ...
func NewAwsSession(awsConfig *aws.Config) (*session.Session, error) {
	var err error
	sharedSessionOnce.Do(func() {
		sharedSession, err = session.NewSession(awsConfig)
	})
	return sharedSession, err
}
