package main
 
import (
	"bufio"
	"fmt"
	"os"
	"net/http"
	"github.com/gookit/color"
	)
	

 
func main() {
	var urls string
	fmt.Println("\n\n▁ ▂ ▄ ▅ ▆ ▇ █   🎀  ＡｌｔＳｕｂ - ｐｒｏ 🎀   █ ▇ ▆ ▅ ▄ ▂ ▁\n",
		       "\n\t\tDeveloped By : Jimmi Simon\n\n",
			   	
	"\tlinkedin -	https://www.linkedin.com/in/jimmisimon/\n",	
	"\tSite	 -	http://jimmisimon.in/projects.php\n")
	color.Error.Println("File Should Not Contain http:// or https://")
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
 	
 	var method string
 	fmt.Printf("Enter the Method you want http/https/both : ")
 	fmt.Scanln(&method)
 	fmt.Println("\n")
 	if method == "http" {
 	
	 	https := "http://"
		for _, eachline := range txtlines {
			url := https+eachline+"/"
			resp, error := http.Get(url)
			if error != nil {
				continue
			}else {
				fmt.Println(url,"\t",resp.StatusCode,http.StatusText(resp.StatusCode))
			}
			
		}
	}else if method == "https" {
		https := "https://"
		for _, eachline := range txtlines {
			url := https+eachline+"/"
			resp, error := http.Get(url)
			if error != nil {
				continue	
			}else {
				fmt.Println(url,"\t",resp.StatusCode,http.StatusText(resp.StatusCode))
			}
		}
	}else if method == "both"{
		
		for _, eachline := range txtlines {
			https := "http://"
			https1 := "https://"
			url := https+eachline+"/"
			resp, error := http.Get(url)
			if error != nil {
				continue
			}else {
				fmt.Println(url,"\t",resp.StatusCode,http.StatusText(resp.StatusCode))
			}
			url1 := https1+eachline+"/"
			resp1, error1 := http.Get(url1)
			if error1 != nil {
				continue
			}else {
				fmt.Println(url1,"\t",resp1.StatusCode,http.StatusText(resp1.StatusCode))
			}
		}
	
	}else{
		fmt.Println("Enter Correct Option")
		os.Exit(0)
	}
	
}
