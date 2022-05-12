package parserpdf

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"code.sajari.com/docconv"
	"github.com/karmdip-mi/go-fitz"
)

var (
	inPath  = "example_in"
	outPath = "example_out"

	regAlarms    = regexp.MustCompile(`(?ism)([\w| :\/]+?[\n]){2}(\d{2}\/\d{2}\/\d{4}[\n])(\d{2}:\d{2}:\d{2}[\n])(.+?[\n]{2})`)
	regAudit     = regexp.MustCompile(`(?is)(\d{2}:\d{2}:\d{2}\n\d{2}\/\d{2}\/\d{4}\n)(.+?)( \n)`)
	regCutTitle1 = regexp.MustCompile(`(?si)(Batch Record)(.+?)(Reason)`)
	regCutTitle2 = regexp.MustCompile(`(?si)(Batch Record)(.+?)(Remark:)`)

	requireParameter []string = []string{
		"machine_name",
		"batch",
		"machine_number",
		"operator",
		"product",
		"customer",
		"order",
		"startdate",
		"stopdate",
		"opertime",
		"stoptime",
		"faulttime",
		"goodprod",
		"rejects",
		"startloss",
		"totalprod",
		"proudtime",
		"startuptime",
		"target_dosing_weight",
	}
)

func main() {

	inFiles, err := getListFiles(inPath, ".pdf")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	for _, file := range inFiles {

		switch {
		case strings.Contains(file, "Message.pdf"):
			fileData, err := getDataFitz(file)
			if err != nil {
				fmt.Println("ERROR =", err)
			}
			err = tablesAlarms(fileData, file)
			if err != nil {
				fmt.Println("ERROR =", err)
			}

		case strings.Contains(file, "infor.pdf"):
			fileData, err := getDataFitz(file)
			if err != nil {
				fmt.Println("ERROR =", err)
			}
			err = tableAudit(fileData, file)
			if err != nil {
				fmt.Println("ERROR =", err)
			}

		default:
			err = getDataBatch(file)
			if err != nil {
				fmt.Println("ERROR =", err)
			}

		}
	}
}

// getListFiles is get list of files in given folder.
func getListFiles(root, extension string) (list []string, err error) {

	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if extension != "*" && filepath.Ext(path) != extension {
			return nil
		}

		list = append(list, path)

		return nil
	})

	return list, err
}

// getDataBatch is getting datas batch from pdf
// using library code.sajari.com/docconv.
func getDataBatch(file string) error {
	file_out := strings.Replace(strings.Replace(file, inPath, outPath, 1), ".pdf", "_Batch.txt", 1)
	err := createOutDir(file_out)
	if err != nil {
		return err
	}

	fout, err := os.Create(file_out)
	if err != nil {
		return err
	}
	defer fout.Close()

	resBatch, err := docconv.ConvertPath(file)
	if err != nil {
		return err
	}

	var lines []string
	ss := strings.Split(resBatch.Body, "\n")
	for _, s := range ss {
		s = cutExcess(s)
		if s != "" {
			lines = append(lines, s)
		}
	}

	var sb strings.Builder
	for i := 0; i < len(lines); i++ {
		s := strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(lines[i], ":", ""), " ", "_"))
		if id := find(requireParameter, s); id > -1 {
			sb.WriteString(requireParameter[id]) // parameter
			sb.WriteString(" ")
			sb.WriteString(lines[i+1]) // value
			sb.WriteString("\n")
		}
	}
	fout.WriteString(sb.String())

	return nil
}

// cutExcess is clears string from '\r', '\n' and Spaces from edges.
func cutExcess(s string) string {
	s = strings.Replace(s, "\r", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	s = strings.TrimSpace(s)
	return s
}

// find is determines string exists in []string.
func find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return -1
}

// getDataFitz is getting data from pdf
// using library github.com/karmdip-mi/go-fitz.
func getDataFitz(file string) (data string, err error) {
	doc, err := fitz.New(file)
	if err != nil {
		return data, err
	}
	defer doc.Close()

	var sb strings.Builder

	for n := 0; n < doc.NumPage(); n++ {
		txt, err := doc.Text(n)
		if err != nil {
			return data, err
		}
		sb.WriteString(txt)
	}
	sb.WriteString(" \n")

	return sb.String(), nil
}

// tableAudit is creates & parses files from table 'Audit Trail Record'.
func tableAudit(data, file string) error {

	file_out := strings.Replace(strings.Replace(file, inPath, outPath, 1), ".pdf", "_Audit.txt", 1)
	err := createOutDir(file_out)
	if err != nil {
		return err
	}

	fout, err := os.Create(file_out)
	if err != nil {
		return err
	}
	defer fout.Close()

	data = strings.ReplaceAll(data, "\n\n", "\n")
	data = regCutTitle1.ReplaceAllString(data, "")
	data = regCutTitle2.ReplaceAllString(data, "")

	records := regAudit.FindAllString(data, -1)

	for _, rec := range records {

		var sb strings.Builder

		rec = strings.TrimSpace(rec)
		items := strings.Split(rec, "\n")

		if len(items) == 7 {
			sb.WriteString(items[4])
			sb.WriteString(" ")
			sb.WriteString(strings.TrimSpace(items[1]) + " " + strings.TrimSpace(items[0])) // dt
			sb.WriteString(" ")
			sb.WriteString(strings.TrimSpace(items[6])) // user
			sb.WriteString(" ")
			sb.WriteString(strings.TrimSpace(items[3])) // number
			sb.WriteString(" ")
			sb.WriteString(strings.TrimSpace(items[5])) // desctript

		} else {
			sb.WriteString(items[6])
			sb.WriteString(" ")
			sb.WriteString(strings.TrimSpace(items[1]) + " " + strings.TrimSpace(items[0])) // dt
			sb.WriteString(" ")
			sb.WriteString(strings.TrimSpace(items[8])) // user
			sb.WriteString(" ")
			sb.WriteString(strings.TrimSpace(items[5])) // number
			sb.WriteString(" ")
			sb.WriteString(strings.TrimSpace(items[7])) // desctript
			sb.WriteString(" ")
			sb.WriteString(strings.TrimSpace(items[3])) // old
			sb.WriteString(" ")
			sb.WriteString(strings.TrimSpace(items[4])) // new
		}
		sb.WriteString("\n")

		fout.WriteString(sb.String())
	}

	return nil
}

// tablesAlarms is creates files from alarm tables.
func tablesAlarms(data, file string) error {
	parts := strings.Split(data, "Alarm List")

	// part 1
	if len(parts[2]) > 0 {
		file_out := strings.Replace(strings.Replace(file, inPath, outPath, 1), ".pdf", "_Alarm1.txt", 1)
		err := createOutDir(file_out)
		if err != nil {
			return err
		}
		err = recordAlarms(parts[2], file_out)
		if err != nil {
			return err
		}
	}

	// part 2
	if len(parts[4]) > 0 {
		file_out := strings.Replace(strings.Replace(file, inPath, outPath, 1), ".pdf", "_Alarm2.txt", 1)
		err := createOutDir(file_out)
		if err != nil {
			return err
		}
		err = recordAlarms(parts[4], file_out)
		if err != nil {
			return err
		}
	}

	return nil
}

// recordAlarms is parses data from tables 'Alarm List'.
func recordAlarms(data, file string) error {
	fout, err := os.Create(file)
	if err != nil {
		return err
	}
	defer fout.Close()

	records := regAlarms.FindAllString(data, -1)

	for _, rec := range records {

		var sb strings.Builder

		items := strings.Split(rec, "\n")

		sb.WriteString(strings.TrimSpace(items[2]) + " " + strings.TrimSpace(items[3])) // DT
		sb.WriteString(" ")
		sb.WriteString(strings.TrimSpace(items[1])) // user
		sb.WriteString(" ")
		sb.WriteString(strings.TrimSpace(items[4])) // number
		sb.WriteString(" ")
		sb.WriteString(strings.TrimSpace(items[0])) // descript

		if len(items) > 5 {
			sb.WriteString(" ")
			sb.WriteString(strings.TrimSpace(items[5])) // old
		}
		if len(items) > 6 {
			sb.WriteString(" ")
			sb.WriteString(strings.TrimSpace(items[6])) // new
		}
		sb.WriteString("\n")

		fout.WriteString(sb.String())
	}

	return nil
}

// createOutDir is creates directory to store parsed files.
func createOutDir(file string) error {
	path := strings.TrimRight(file, filepath.Base(file))
	err := os.MkdirAll(path, 0777)
	if err != nil {
		return err
	}
	return nil
}
