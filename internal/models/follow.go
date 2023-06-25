package models

import "context"

type Follower struct {
	followers map[string][]string
}

func NewFollow() Follower {
	return Follower{
		followers: map[string][]string{
			"nikolay": []string{"anna", "maria", "bob"},
			"john":    []string{"anna"},
		},
	}
}

func (f Follower) Followers(ctx context.Context, user string) ([]string, error) {
	return f.followers[user], nil
}
