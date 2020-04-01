package employment

import (
	"time"

	"github.com/fgituser/management-client-psychologist.services/psychologist/internal/users"
)

//Employment ...
type Employment struct {
	Client *users.Client
	Date   time.Time
}
