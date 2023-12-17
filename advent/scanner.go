package advent

import "bytes"

// Scanner function for scanning puzzles where input is large sections of text divided by a full blank line.
func ScanSections(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, []byte{'\n', '\n'}); i >= 0 {
		return i + 2, data[0:i], nil
	}
	if atEOF {
		return len(data), data, nil
	}

	// Request more data.
	return 0, nil, nil
}
