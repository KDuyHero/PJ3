package domain

import "context"

type CommonRepository interface {
	Transaction(context.Context, func(context.Context) error) error
}
