package m5db

import(
        DRM "github.com/fbaube/datarepo/rowmodels"
)

// M5_TableDetails configures the three key tables.
var M5_TableDetails = []DRM.TableDetails{
	TableDetailsCNT,
	TableDetailsINB,
	TableDetailsTRF,
}
