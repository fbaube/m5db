package m5db

import(
        DRM "github.com/fbaube/datarepo/rowmodels"
)

// TableDetailsTRF specifies only two foreign keys.
var TableDetailsTRF = DRM.TableDetails{
        TableSummaryTRF, 
	"idx_topicref", // PKname
	ColumnNamesCsv_TRF, 
	ColumnSpecsTRF, // []D.ColumnSpec
	ColumnPtrsFuncTRF,
	&TopicrefRow{},
}

// TableDetails returns the table
// detail info, given any instance.
func (tro *TopicrefRow) TableDetails() DRM.TableDetails {
     return TableDetailsTRF
}

