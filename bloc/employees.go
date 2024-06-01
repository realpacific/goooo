package bloc

import (
	"encoding/csv"
	"github.com/samber/lo"
	. "goooo/logging"
	"io"
	"os"
	"strings"
)

type employee struct {
	Name     string `json:"name"`
	Genre    string `json:"genre"`
	Position string `json:"position"`
}

var employees = make(map[string]employee)

func init() {
	file, err := os.Open("employees.csv")
	if err != nil {
		Logger.Fatalf("impossible to open file %s", err)
	}

	r := csv.NewReader(file)
	_, _ = r.Read() // skip first line
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			Logger.Fatal(err)
		}
		employee := employee{
			Name:     record[1],
			Genre:    record[2],
			Position: record[3],
		}
		employees[record[0]] = employee
	}
	Logger.Debug(employees)
}

type EmployeeParams struct {
	ByName *string
	ById   *string
}

func ListEmployees(opts *EmployeeParams) []employee {
	var values []employee = lo.Values[string, employee](employees)
	if opts.ById != nil {
		item, hasValue := employees[*opts.ById]
		if !hasValue {
			return []employee{}
		}
		return []employee{item}
	}
	return lo.Filter[employee](values, func(item employee, index int) bool {
		if opts.ByName != nil {
			return strings.Contains(item.Name, *opts.ByName)
		}
		return true
	})
}
