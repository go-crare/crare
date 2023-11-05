package layout

import (
	"gopkg.in/crare"
)

// LocaleFunc is the function used to fetch the locale of the recipient.
// Returned locale will be remembered and linked to the corresponding context.
type LocaleFunc func(crare.Recipient) string

// Middleware builds a crare middleware to make localization work.
//
// Usage:
//
//	b.Use(lt.Middleware("en", func(r crare.Recipient) string {
//		loc, _ := db.UserLocale(r.Recipient())
//		return loc
//	}))
func (lt *Layout) Middleware(defaultLocale string, localeFunc ...LocaleFunc) crare.HandlerFunc {
	var f LocaleFunc
	if len(localeFunc) > 0 {
		f = localeFunc[0]
	}

	return func(c *crare.Context) error {
		locale := defaultLocale
		if f != nil {
			if l := f(c.Sender()); l != "" {
				locale = l
			}
		}

		lt.SetLocale(c, locale)

		defer func() {
			lt.mu.Lock()
			delete(lt.ctxs, c)
			lt.mu.Unlock()
		}()

		return c.Next()
	}
}

// Middleware wraps ordinary layout middleware with your default locale.
func (dlt *DefaultLayout) Middleware() crare.HandlerFunc {
	return dlt.lt.Middleware(dlt.locale)
}
