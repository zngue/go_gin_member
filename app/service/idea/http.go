package idea

import (
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"sync"
)

type Ideas struct {
	One string
	OneLength int
	Two string
	TwoLength int
	Three string
	ThreeLength int

}

func ( Ideas ) GetIdeaOne(code *Ideas, wg *sync.WaitGroup)   {
	defer wg.Done()
	req ,_:=http.Get("http://www.lookdiv.com/code")
	doc,_:=goquery.NewDocumentFromResponse(req)
	text := doc.Find("textarea").Text()
	code.One=text
	code.OneLength= len([]byte(text))

	return
}
func (Ideas) GetIdeaTwo(code *Ideas, wg *sync.WaitGroup)   {
	defer wg.Done()
	req,_:=http.Get("http://www.yanjie.site/jjj2.txt")
	byte,_:=ioutil.ReadAll(req.Body)
	codeString:=string(byte)
	code.Two=codeString

	return
}
func (i *Ideas) GetAllIdea()  *Ideas  {
	var wg sync.WaitGroup
	var idea Ideas
	wg.Add(2)
	//go GetIdeaTwo(&idea,&wg)
	go i.GetIdeaOne(&idea,&wg)
	go i.GetIdeaTwo(&idea,&wg)
	wg.Wait()
	return &idea
}
