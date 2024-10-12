package excel_file_reader

type ExcelFileReaderFunc interface {
	CheckFilePaths() error

	ReadExcelFilesAsStdout() error
}

type ExcelFileReader struct {
	ExcelFileReaderFunc ExcelFileReaderFunc

	filePaths          []string
	isIgnoreRead       bool
	isIgnoreRowsRead   bool
	isIgnoreColumnRead bool
	isIgnoreCsvWrite   bool
}

var (
	defaultOptionExcelFileReader = setDefaultOptionExcelFileReader(nil)
)

func setDefaultOptionExcelFileReader(filePaths []string) *ExcelFileReader {
	return &ExcelFileReader{
		filePaths: filePaths,
	}
}

type ExcelFileReaderOption func(*ExcelFileReader)

// NewExcelFileReader
//
//	use as
//
//	changeExcelFileReader := NewExcelFileReader(
//	    With(""),
//	)
func NewExcelFileReader(filePaths []string, opts ...ExcelFileReaderOption) (opt *ExcelFileReader) {
	opt = setDefaultOptionExcelFileReader(filePaths)
	for _, o := range opts {
		o(opt)
	}
	defaultOptionExcelFileReader = setDefaultOptionExcelFileReader(filePaths)
	return
}

func WithIgnoreRead(isIgnoreRead bool) ExcelFileReaderOption {
	return func(o *ExcelFileReader) {
		o.isIgnoreRead = isIgnoreRead
	}
}

func WithIgnoreRowsRead(isIgnoreRowsRead bool) ExcelFileReaderOption {
	return func(o *ExcelFileReader) {
		o.isIgnoreRowsRead = isIgnoreRowsRead
	}
}

func WithIgnoreColumnRead(isIgnoreColumnRead bool) ExcelFileReaderOption {
	return func(o *ExcelFileReader) {
		o.isIgnoreColumnRead = isIgnoreColumnRead
	}
}

func WithIgnoreCsvWrite(isIgnoreCsvWrite bool) ExcelFileReaderOption {
	return func(o *ExcelFileReader) {
		o.isIgnoreCsvWrite = isIgnoreCsvWrite
	}
}
