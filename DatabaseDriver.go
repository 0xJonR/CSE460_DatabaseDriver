/*
CSE460 Database Driver Script
execute with 'go run .'
Written by Jon Rodriguez
October 2021
-------------------------------
TODO:
1) Add Unzip functionality to .gz files after downloading
2) Add SQL Driver code to create database and UPDATE with new data
3) Add ability to create minified version of database
*/
package main

//imports
import(
	"fmt"

	"net/http"

	"os"

	"io"
)

//constants: 
const IMDB_SRC_LINK = "https://datasets.imdbws.com/"
const IMDB_name_basics = "name.basics.tsv.gz"
const IMDB_title_akas = "title.akas.tsv.gz"
const IMDB_title_basics = "title.basics.tsv.gz"
const IMDB_title_crew = "title.crew.tsv.gz"
const IMDB_title_episode = "title.episode.tsv.gz"
const IMDB_title_principals = "title.principals.tsv.gz"
const IMDB_title_ratings = "title.ratings.tsv.gz"

//structs

//functions

func buildLink(which string) string {
	l := IMDB_SRC_LINK + which
	return l
}

/* grab makes an http GET request to the link
and returns the byte representation of the data.
*/
func grab(fp string, link string, finished chan bool) (error) {
	// create file handle
	out, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer out.Close()
	//GET data
	fmt.Println("Downloading: ", fp)
	resp, err2 := http.Get(link)
	if err2 != nil {
		fmt.Println("bad GET")
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}
	//write the body to file
	_, err3 := io.Copy(out, resp.Body)
	if err3 != nil {
		return err3
	}
	defer func(){
		fmt.Printf("\nFinished downloading to file: %s", fp)
	}()
	finished <- true
	return nil
}
func main() {
	imdb_fns := [7]string{IMDB_name_basics, IMDB_title_akas, IMDB_title_basics, IMDB_title_crew, IMDB_title_episode, IMDB_title_principals, IMDB_title_ratings}
	finished := make(chan bool)
	//fmt.Println(imdb_fns)
	for i := 0; i<7; i++ {
		fn := imdb_fns[i]
		link := buildLink(fn)
		fmt.Println(fn)
		go grab(fn, link, finished) //goroutine for speed :D concurrency ftw 
	}
	for g :=0; g<7; g++ {
		<-finished
	}
	fmt.Println("\nFinsihed.")
}
