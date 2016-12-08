package graphqlGoClient
import "net/http"
import "io/ioutil"
import "bytes"

type graphqlClient struct{
  hostUrl string
  httpClient *http.Client
  header http.Header
}
//returns graphqlClient
func New(url string)*graphqlClient{
  header:=make(http.Header)
  header.Add("Content-Type","application-graphql")
  return &graphqlClient{hostUrl:url,httpClient:&http.Client{},header:header}
}
//used to add headers
func (client *graphqlClient)AddHeader(headerName string, headerValue string){
  client.header.Add(headerName,headerValue)
}
//used for querying
func (client *graphqlClient)Do(query string)([]byte,error)  {
  req,err:=http.NewRequest("POST",client.hostUrl,bytes.NewBuffer([]byte(query)))
  req.Header=client.header
  res,err:=client.httpClient.Do(req)
  if err!=nil{
    return nil,err
  }
  if err!=nil {
    return nil,err
  }
  defer res.Body.Close()
  body,err:=ioutil.ReadAll(res.Body)
  if err!=nil {
    return nil,err
  }
  return body,nil
}
