package resolver

// Channel represents the Channel GraphQL object type
type Channel struct {
	id   string
	name string
}

// ID ...
func (c *Channel) ID() string {
	return c.id
}

// Name ...
func (c *Channel) Name() string {
	return c.name
}
