// Go in action
// @jeffotoni
// 2019-01-16

package main

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/gdamore/tcell"
	_ "github.com/lib/pq"
	"github.com/rivo/tview"
)

const (
	batchSize  = 80         // The number of rows loaded per batch.
	finderPage = "*finder*" // The name of the Finder page.
)

var (
	app         *tview.Application // The tview application.
	pages       *tview.Pages       // The application pages.
	finderFocus tview.Primitive    // The primitive in the Finder that last had focus.
)

// Main entry point.
func main() {
	// Get connect string from the command line.
	if len(os.Args) < 2 {
		fmt.Println(`Please provide a PostgreSQL connect string.`)
		fmt.Println("See https://godoc.org/github.com/lib/pq for details.")
		return
	}

	// Start the application.
	app = tview.NewApplication()
	finder(os.Args[1])
	if err := app.Run(); err != nil {
		fmt.Printf("Error running application: %s\n", err)
	}
}

// Sets up a "Finder" used to navigate the databases, tables, and columns.
func finder(connString string) {
	// Create the basic objects.
	databases := tview.NewList().ShowSecondaryText(false)
	databases.SetBorder(true).SetTitle("Databases")
	columns := tview.NewTable().SetBorders(true)
	columns.SetBorder(true).SetTitle("Columns")
	tables := tview.NewList()
	tables.ShowSecondaryText(false).
		SetDoneFunc(func() {
			tables.Clear()
			columns.Clear()
			app.SetFocus(databases)
		})
	tables.SetBorder(true).SetTitle("Tables")

	// Create the layout.
	flex := tview.NewFlex().
		AddItem(databases, 0, 1, true).
		AddItem(tables, 0, 1, false).
		AddItem(columns, 0, 3, false)

	// We keep one connection pool per database.
	dbMutex := sync.Mutex{}
	dbs := make(map[string]*sql.DB)
	getDB := func(database string) *sql.DB {
		// Connect to a new database.
		dbMutex.Lock()
		defer dbMutex.Unlock()
		if db, ok := dbs[database]; ok {
			return db
		}
		var newConnString string
		if u, err := url.Parse(connString); err == nil && u.Scheme == "postgres" {
			u.Path = "/" + database
			newConnString = u.String()
		} else {
			dbstring := regexp.MustCompile(`(dbname=)'?[^ ]*'?`)
			matches := dbstring.FindStringSubmatch(connString)
			if matches == nil {
				newConnString = "dbname='" + database + "' " + connString
			} else {
				newConnString = dbstring.ReplaceAllString(connString, "$1'"+database+"'")
			}
		}
		db, err := sql.Open("postgres", newConnString)
		if err != nil {
			panic(err)
		}
		dbs[database] = db
		return db
	}

	// fmt.Println(connString)
	// Get a list of all databases.
	generalDB, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}
	defer generalDB.Close() // We really close the DB because we only use it for this one query.
	rows, err := generalDB.Query("select datname from pg_database where datistemplate = false")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var dbName string
		if err := rows.Scan(&dbName); err != nil {
			panic(err)
		}
		databases.AddItem(dbName, "", 0, func() {
			// A database was selected. Show all of its tables.
			columns.Clear()
			tables.Clear()
			db := getDB(dbName)
			t, err := db.Query("select table_name from information_schema.tables where table_schema = 'public'")
			if err != nil {
				panic(err)
			}
			defer t.Close()
			for t.Next() {
				var tableName string
				if err := t.Scan(&tableName); err != nil {
					panic(err)
				}
				tables.AddItem(tableName, "", 0, nil)
			}
			if err := t.Err(); err != nil {
				panic(err)
			}
			app.SetFocus(tables)

			// When the user navigates to a table, show its columns.
			tables.SetChangedFunc(func(i int, tableName string, t string, s rune) {
				// A table was selected. Show its columns.
				columns.Clear()
				c, err := db.Query(`
            select
              c.column_name,
              c.is_nullable,
              c.data_type,
              c.character_maximum_length,
              c.numeric_precision,
              c.numeric_scale,
              c.ordinal_position,
              tc.constraint_type pkey
            from
              information_schema.columns c
            left join
              information_schema.constraint_column_usage as ccu
            on
              c.table_schema = ccu.table_schema
              and c.table_name = ccu.table_name
              and c.column_name = ccu.column_name
            left join
              information_schema.table_constraints as tc
            on
              ccu.constraint_schema = tc.constraint_schema
              and ccu.constraint_name = tc.constraint_name
            where
              c.table_schema = 'public'
              and c.table_name = $1
          `, tableName)
				if err != nil {
					panic(err)
				}
				defer c.Close()
				columns.SetCell(0, 0, &tview.TableCell{Text: "Name", Align: tview.AlignCenter, Color: tcell.ColorYellow}).
					SetCell(0, 1, &tview.TableCell{Text: "Type", Align: tview.AlignCenter, Color: tcell.ColorYellow}).
					SetCell(0, 2, &tview.TableCell{Text: "Size", Align: tview.AlignCenter, Color: tcell.ColorYellow}).
					SetCell(0, 3, &tview.TableCell{Text: "Null", Align: tview.AlignCenter, Color: tcell.ColorYellow}).
					SetCell(0, 4, &tview.TableCell{Text: "Constraint", Align: tview.AlignCenter, Color: tcell.ColorYellow})
				for c.Next() {
					var (
						columnName, isNullable, dataType     string
						constraintType                       sql.NullString
						size, numericPrecision, numericScale sql.NullInt64
						ordinalPosition                      int
					)
					if err := c.Scan(&columnName,
						&isNullable,
						&dataType,
						&size,
						&numericPrecision,
						&numericScale,
						&ordinalPosition,
						&constraintType,
					); err != nil {
						panic(err)
					}
					sizeText := ""
					if size.Valid {
						sizeText = strconv.Itoa(int(size.Int64))
					} else if numericPrecision.Valid {
						sizeText = strconv.Itoa(int(numericPrecision.Int64))
						if numericScale.Valid {
							sizeText += "," + strconv.Itoa(int(numericScale.Int64))
						}
					}
					color := tcell.ColorWhite
					if constraintType.Valid {
						color = map[string]tcell.Color{
							"CHECK":       tcell.ColorGreen,
							"FOREIGN KEY": tcell.ColorDarkMagenta,
							"PRIMARY KEY": tcell.ColorRed,
							"UNIQUE":      tcell.ColorDarkCyan,
						}[constraintType.String]
					}
					columns.SetCell(ordinalPosition, 0, &tview.TableCell{Text: columnName, Color: color}).
						SetCell(ordinalPosition, 1, &tview.TableCell{Text: dataType, Color: color}).
						SetCell(ordinalPosition, 2, &tview.TableCell{Text: sizeText, Align: tview.AlignRight, Color: color}).
						SetCell(ordinalPosition, 3, &tview.TableCell{Text: isNullable, Align: tview.AlignRight, Color: color}).
						SetCell(ordinalPosition, 4, &tview.TableCell{Text: constraintType.String, Align: tview.AlignLeft, Color: color})
				}
				if err := c.Err(); err != nil {
					panic(err)
				}
			})
			tables.SetCurrentItem(0) // Trigger the initial selection.

			// When the user selects a table, show its content.
			tables.SetSelectedFunc(func(i int, tableName string, t string, s rune) {
				content(db, dbName, tableName)
			})
		})
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}

	// Set up the pages and show the Finder.
	pages = tview.NewPages().
		AddPage(finderPage, flex, true, true)
	app.SetRoot(pages, true)
}

// Shows the contents of the given table.
func content(db *sql.DB, dbName, tableName string) {
	finderFocus = app.GetFocus()

	// If this page already exists, just show it.
	if pages.HasPage(dbName + "." + tableName) {
		pages.SwitchToPage(dbName + "." + tableName)
		return
	}

	// We display the data in a table embedded in a frame.
	table := tview.NewTable().
		SetFixed(1, 0).
		//SetSeparator(tview.GraphicsVertBar).
		SetBordersColor(tcell.ColorYellow)
	frame := tview.NewFrame(table).
		SetBorders(0, 0, 0, 0, 0, 0)
	frame.SetBorder(true).
		SetTitle(fmt.Sprintf(`Contents of table "%s"`, tableName))

	// How many rows does this table have?
	var rowCount int
	err := db.QueryRow(fmt.Sprintf("select count(*) from %s", tableName)).Scan(&rowCount)
	if err != nil {
		panic(err)
	}

	// Load a batch of rows.
	loadRows := func(offset int) {
		rows, err := db.Query(fmt.Sprintf("select * from %s limit $1 offset $2", tableName), batchSize, offset)
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		// The first row in the table is the list of column names.
		columnNames, err := rows.Columns()
		if err != nil {
			panic(err)
		}
		for index, name := range columnNames {
			table.SetCell(0, index, &tview.TableCell{Text: name, Align: tview.AlignCenter, Color: tcell.ColorYellow})
		}

		// Read the rows.
		columns := make([]interface{}, len(columnNames))
		columnPointers := make([]interface{}, len(columns))
		for index := range columnPointers {
			columnPointers[index] = &columns[index]
		}
		for rows.Next() {
			// Read the columns.
			err := rows.Scan(columnPointers...)
			if err != nil {
				panic(err)
			}

			// Transfer them to the table.
			row := table.GetRowCount()
			for index, column := range columns {
				switch value := column.(type) {
				case int64:
					table.SetCell(row, index, &tview.TableCell{Text: strconv.Itoa(int(value)), Align: tview.AlignRight, Color: tcell.ColorDarkCyan})
				case float64:
					table.SetCell(row, index, &tview.TableCell{Text: strconv.FormatFloat(value, 'f', 2, 64), Align: tview.AlignRight, Color: tcell.ColorDarkCyan})
				case string:
					table.SetCellSimple(row, index, value)
				case time.Time:
					t := value.Format("2006-01-02")
					table.SetCell(row, index, &tview.TableCell{Text: t, Align: tview.AlignRight, Color: tcell.ColorDarkMagenta})
				case []uint8:
					str := make([]byte, len(value))
					for index, num := range value {
						str[index] = byte(num)
					}
					table.SetCell(row, index, &tview.TableCell{Text: string(str), Align: tview.AlignRight, Color: tcell.ColorGreen})
				case nil:
					table.SetCell(row, index, &tview.TableCell{Text: "NULL", Align: tview.AlignCenter, Color: tcell.ColorRed})
				default:
					// We've encountered a type that we don't know yet.
					t := reflect.TypeOf(value)
					str := "?nil?"
					if t != nil {
						str = "?" + t.String() + "?"
					}
					table.SetCellSimple(row, index, str)
				}
			}
		}
		if err := rows.Err(); err != nil {
			panic(err)
		}

		// Show how much we've loaded.
		frame.Clear()
		loadMore := ""
		if table.GetRowCount()-1 < rowCount {
			loadMore = " - press Enter to load more"
		}
		loadMore = fmt.Sprintf("Loaded %d of %d rows%s", table.GetRowCount()-1, rowCount, loadMore)
		frame.AddText(loadMore, false, tview.AlignCenter, tcell.ColorYellow)
	}

	// Load the first batch of rows.
	loadRows(0)

	// Handle key presses.
	table.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEscape:
			// Go back to Finder.
			pages.SwitchToPage(finderPage)
			if finderFocus != nil {
				app.SetFocus(finderFocus)
			}
		case tcell.KeyEnter:
			// Load the next batch of rows.
			loadRows(table.GetRowCount() - 1)
			table.ScrollToEnd()
		}
	})

	// Add a new page and show it.
	pages.AddPage(dbName+"."+tableName, frame, true, true)
}
