package middleware

import "gopkg.in/crare.v1"

// RestrictConfig defines config for Restrict middleware.
type RestrictConfig struct {
	// Chats is a list of chats that are going to be affected
	// by either In or Out function.
	Chats []int64
}

// Restrict returns a middleware that handles a list of provided
// chats with the logic defined by In and Out functions.
// If the chat is found in the Chats field, In function will be called,
// otherwise Out function will be called.
func Restrict(v RestrictConfig) crare.HandlerFunc {
	return func(c crare.Context) error {
		for _, chat := range v.Chats {
			if chat == c.Sender().ID {
				return c.Next()
			}
		}
		return nil
	}
}

// Blacklist returns a middleware that skips the update for users
// specified in the chats field.
func Blacklist(chats ...int64) crare.HandlerFunc {
	return func(ctx crare.Context) error {
		return Restrict(RestrictConfig{
			Chats: chats,
		})(ctx)
	}
}

// Whitelist returns a middleware that skips the update for users
// NOT specified in the chats field.
func Whitelist(chats ...int64) crare.HandlerFunc {
	return func(ctx crare.Context) error {
		return Restrict(RestrictConfig{
			Chats: chats,
		})(ctx)
	}
}
