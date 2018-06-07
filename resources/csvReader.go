package resources

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
)

//name,image_url,width,height,info,date,style,location,dimensions,media,author_id,name,born,died,nationality,wikipedia

type CSVPhoto struct {
	Name        string `csv:"name"`
	ImageURL    string `csv:"image_url"`
	Width       string `csv:"width"`
	Height      string `csv:"height"`
	Info        string `csv:"info"`
	Date        string `csv:"date"`
	Style       string `csv:"style"`
	Location    string `csv:"location"`
	Dimensions  string `csv:"dimensions"`
	Media       string `csv:"media"`
	AuthorId    string `csv:"author_id"`
	AuthorName  string `csv:"author_name"`
	Born        string `csv:"born"`
	Died        string `csv:"died"`
	Nationality string `csv:"nationality"`
	Wikipedia   string `csv:"wikipedia"`
}

func ReadCSVFile(name string) ([]*CSVPhoto, error) {
	clientsFile, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer clientsFile.Close()

	csvPhotos := []*CSVPhoto{}
	if gocsv.UnmarshalFile(clientsFile, &csvPhotos); err != nil {
		return nil, err
	}

	// Log
	for _, csvPhoto := range csvPhotos {
		fmt.Println(csvPhoto)
	}

	return csvPhotos, nil
}
