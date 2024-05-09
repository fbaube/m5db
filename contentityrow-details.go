package m5db

import(
	DRM "github.com/fbaube/datarepo/rowmodels"
)

// Implement interface RowModeler

// TableDetailsCNT specifies 11 DB columns,
// incl primary key (assumed) and one foreign key, "inbatch".
var TableDetailsCNT = DRM.TableDetails{
        TableSummaryCNT, 
	"idx_contentity", // IDName
	"IDX_inbatch, RelFP, AbsFP, Descr, T_Cre, T_Imp, T_Edt, " +
		"RawMT, Mimtp, MType, Contt", // ColumnNames
	// One foreign key: "inbatch"
	ColumnSpecsCNT, // []D.ColumnSpecs
	ColumnPtrsFuncCNT,
	// ColumnPtrsMthdCNT,
	&ContentityRow{}, // RowModeler(new(ContentityRow)),
}

// TableDetails returns the table
// detail info, given any instance. 
func (cro *ContentityRow) TableDetails() DRM.TableDetails {
     return TableDetailsCNT
}

