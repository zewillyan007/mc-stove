package util

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

func UnzipFile(inZipFile string, outDir string) error {

	r, err := zip.OpenReader(inZipFile)

	if err != nil {
		log.Printf("Error opening zip file: %#e", err)
		return err
	}

	defer r.Close()

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(outDir, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.Create(fpath)
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}
	}
	return nil
}

func ZipFiles(outFileName string, files []string) error {

	newZipFile, err := os.Create(outFileName)
	if err != nil {
		log.Printf("Error creating zip file: %#e", err)
		return err
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	for _, fn := range files {
		if err = addFileToZip(zipWriter, fn); err != nil {
			log.Printf("Error adding file to zipfile: %#e", err)
			return err
		}
	}
	return nil
}

func addFileToZip(zipWriter *zip.Writer, fn string) error {

	fileToZip, err := os.Open(fn)
	if err != nil {
		log.Printf("Error opening file to zip: %#e", err)
		return err
	}
	defer fileToZip.Close()

	info, err := fileToZip.Stat()
	if err != nil {
		log.Printf("Error grabbing file stats: %#e", err)
		return err
	}
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		log.Printf("Error grabbing file info header: %#e", err)
		return err
	}

	header.Name = fn
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		log.Printf("Error creating zip header: %#e", err)
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	if err != nil {
		log.Printf("Error copying file to zip file: %#e", err)
		return err
	}
	return nil
}
