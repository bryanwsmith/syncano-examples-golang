import(
    "fmt"
    "io/ioutil"
    "net/http"
    "math/rand"
    "regexp"
    "time"
)


func main(){
    url := "https://www.google.com/search?q=potato&tbm=isch"
    
    resp, _ := http.Get(url)
    defer resp.Body.Close()    
    
    body, _ := ioutil.ReadAll(resp.Body)
    
    re := regexp.MustCompile("src=\"(http[^\"]+)\"")
    matches := re.FindAllStringSubmatch(string(body), -1)
    
    potatoes := make([]string, len(matches))
    
    for index, match := range matches {
        potatoes[index] = match[1]
    }
    
    //seed with nanoseconds to get make sure unique random number returned
    rand.Seed(time.Now().UnixNano())
    
    //get random image url and print to stdout
    fmt.Println(potatoes[rand.Intn(len(potatoes))])
}