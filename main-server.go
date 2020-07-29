package main
 
import (
	"bufio"
	"fmt"
	"os"
	"net/http"
	"strings"
)
 
func main() {
	fmt.Println("\n\n▁ ▂ ▄ ▅ ▆ ▇ █   🎀  ＡｌｔＳｕｂ - ｐｒｏ 🎀   █ ▇ ▆ ▅ ▄ ▂ ▁\n",
		       "\n\t\tDeveloped By : Jimmi Simon\n\n",
			   	
	"\tlinkedin -	https://www.linkedin.com/in/jimmisimon/\n",	
	"\tSite	 -	http://jimmisimon.in/projects.php\n")
	var urls string
	fmt.Printf("Enter Subdomain list FileName : ")
	fmt.Scanln(&urls)
	
	file, err := os.Open(urls)
 
	if err != nil {
		fmt.Println("Check the FileName / File Not Available")
		os.Exit(0)
	}
 
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
 
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
 
	file.Close()
 	
 	https := "http://"
	for _, eachline := range txtlines {
		url := https+eachline+"/"
		resp, error := http.Get(url)
		if error != nil {
			continue
			
		}else {
			for k,v := range resp.Header{
				if strings.HasPrefix(k,"Server") && strings.HasSuffix(k,"Server"){
					fmt.Println(url,"\t",resp.StatusCode,http.StatusText(resp.StatusCode),"\t",v)
					
				}
				
			}
			
		}
		
	}
}
