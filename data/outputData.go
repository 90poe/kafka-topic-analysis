package data

import (
	"encoding/csv"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
)

func RenderTable(dataTable [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"_______Event_Time_______", "Interval (Seconds)", "Gyro", "Magnetometer", "Compass", "Accelerometer", "TiltX", "TiltY"})
	for _, v := range dataTable {
		table.Append(v)
	}
	table.SetColMinWidth(0, 50)
	table.SetRowLine(true)

	table.Render()

}

func ToCSVFile(dataTable [][]string, filename string) {
	// to csv
	file, err := os.Create(filename)
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range dataTable {
		err := writer.Write(value)
		if err != nil {
			log.Println(err.Error())
		}
	}
	log.Println("The file has bee successfully created: " + filename)
}
