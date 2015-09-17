import(
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "math/rand"
    "regexp"
    "time"
)

type StringSubmatchArray [][]string

func (s StringSubmatchArray) Choice() string {
    rand.Seed(time.Now().UnixNano())
    return s[rand.Intn(len(s))][1]
}

func main(){
    url := "https://www.google.com/search?q=potato&tbm=isch"
    resp, err := http.Get(url)
    
    if err == nil && resp.StatusCode == 200 { 

        defer resp.Body.Close()    
    
        body, err := ioutil.ReadAll(resp.Body)
        
        if err == nil{
            
            re := regexp.MustCompile("src=\"(http[^\"]+)\"")
            potatoes := re.FindAllStringSubmatch(string(body), -1)
            fmt.Println(StringSubmatchArray(potatoes).Choice())
            
        } else {
            log.Println(err)
        } 
    } else {
        log.Println(err)
    }
}