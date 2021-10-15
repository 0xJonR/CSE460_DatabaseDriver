package main

//imports
import(
	"fmt"

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
func main() {
	fmt.Println("test")
	fmt.Println(buildLink(IMDB_title_akas))
}
