package JenkinsBuild

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jenkinsServer/Model"
	"jenkinsServer/Utill"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func Build(c *gin.Context) {
	debReq := Model.DebBuildRequest{}
	debReq.Branch = c.Query("branch")
	debReq.Business = c.Query("business")

	urlValues := url.Values{}
    urlValues.Add("branch", debReq.Branch)
    urlValues.Add("business", debReq.Business)
	reqUrl := "http://" + Utill.Conf.JenkinsAddress + ":"+ Utill.Conf.JenkinsPort + "/generic-webhook-trigger/invoke?token=" + Utill.Conf.Token;
	fmt.Printf("URL : %s\n", reqUrl)
    
	resp, _ := http.PostForm(reqUrl,urlValues)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))

	debReqJson , _ := json.Marshal(debReq)
	
	c.JSON(http.StatusOK, gin.H{
		"Code" : http.StatusOK,
		"data": string(debReqJson),
		"Msg":"ReqSuccess",
	})

}