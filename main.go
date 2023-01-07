package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jadefox10200/goprint"
)

func main() {

	filePath := flag.String("f", "", "[REQUIERD] File path to PDF")

	defaultPrinterName, _ := goprint.GetDefaultPrinterName()
	printerName := flag.String("p", defaultPrinterName, "Specifies custom printer name. if NOT set default prrinter will be used")

	flag.Parse()

	if len([]rune(*filePath)) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	//open the printer
	printerHandle, err := goprint.GoOpenPrinter(*printerName)
	if err != nil {
		log.Fatalln("Failed to open printer")
	}
	defer goprint.GoClosePrinter(printerHandle)

	//Send to printer:
	err = goprint.GoPrint(printerHandle, *filePath)
	if err != nil {
		log.Fatalln("during the func sendToPrinter, there was an error")
	}

}
