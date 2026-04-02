// Package secretsource registers all the internal secret sources when imported
package secretsource

import (
	_ "github.com/daudcanugerah/k6/internal/secretsource/file" // import them for init
	_ "github.com/daudcanugerah/k6/internal/secretsource/mock" // import them for init
	_ "github.com/daudcanugerah/k6/internal/secretsource/url"  // import them for init
)
