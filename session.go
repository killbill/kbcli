package kbcli
import (
	"net/http"
	"net/url"
	"strings"
	"log"
	"bytes"
	"encoding/json"
	"time"
	"io/ioutil"
)

const KILLBILL_PREFIX string = "/1.0/kb"


type Session struct {
	// client
	Client       *http.Client
	//
	KillbillIp   string
	KillbillPort string
	// Basic Auth
	Userinfo     *url.Userinfo
	// Multi-tenancy
	ApiKey       string
	ApiSecret    string
	// Default headers
	Header       *http.Header
	// TODO
	JsessionId   string
	Log          bool
}

func (s *Session) createUrl(resourcePathOrUrl string) string {
	if strings.HasPrefix(resourcePathOrUrl, "http:") {
		return resourcePathOrUrl
	} else {
		parts := []string{"http://", s.KillbillIp, ":", s.KillbillPort, KILLBILL_PREFIX, resourcePathOrUrl}
		return strings.Join(parts, "")
	}
}


func (s *Session) Post(resourcePathOrUrl string, body interface{}) (*Response, error) {
	request := Request {
		Method:  "POST",
		Url:  s.createUrl(resourcePathOrUrl),
		Body: body,
	}
	return s.Send(&request, nil)
}

func (s *Session) Get(resourcePathOrUrl string, result JsonDeserializer, queryParams *QueryParams) (*Response, error) {
	request := Request {
		Method: "GET",
		Url:   s.createUrl(resourcePathOrUrl),
		QueryParams: queryParams,
	}
	return s.Send(&request, result)
}


func (s *Session) log(args ...interface{}) {
	if s.Log {
		log.Println(args...)
	}
}

func (s *Session) setJsessionId(headers http.Header) {

	const JSESSIONID string = "JSESSIONID"

	// JsessionId has already been set
	if s.JsessionId != "" {
		return
	}

	if value, ok := headers["Set-Cookie"]; ok {
		for _, element := range value {
			if strings.HasPrefix(element, JSESSIONID) {
				if strings.HasPrefix(element, JSESSIONID) {
					jSessionId := element[len("JSESSIONID:"):13]
					sessionsParts := []string{JSESSIONID, jSessionId}
					s.JsessionId = strings.Join(sessionsParts, "=")
				}
			}
		}
	}
}



// Send constructs and sends an HTTP request.
func (s *Session) Send(inputRequest *Request, responseResult JsonDeserializer) (response *Response, err error) {

	// Allocate a Response object and initialize its Result field with the responseResult we expect to see
	response = &Response {
		Result: responseResult,
	}

	// Extract Url
	parsedUrl, err := url.Parse(inputRequest.Url)
	if err != nil {
		s.log("URL", inputRequest.Url)
		s.log(err)
		return
	}

	// Encode query parameters
	if inputRequest.QueryParams != nil {
		vals := parsedUrl.Query()
		for k, v := range *inputRequest.QueryParams {
			vals.Set(k, v)
		}
		parsedUrl.RawQuery = vals.Encode()
	}

	// Add generic headers
	header := http.Header{}
	if s.Header != nil {
		for k, _ := range *s.Header {
			v := s.Header.Get(k)
			header.Set(k, v)
		}
	}

	//
	// Add Session specific headers
	//
	// Kill Bill Tenant  headers from the Session
	if s.ApiKey != "" && s.ApiSecret != "" {
		header.Set("X-Killbill-ApiKey", s.ApiKey)
		header.Set("X-Killbill-ApiSecret", s.ApiSecret)
	}
	// JSESSIONID
	if s.JsessionId != "" {
		header.Set("Cookie", s.JsessionId)
	}

	if inputRequest.Header != nil {
		for k, v := range *inputRequest.Header {
			header.Set(k, v[0])
		}
	}

	// TODO only accept
	if header.Get("Accept") == "" {
		header.Add("Accept", "application/json")
	}


	// Create the http request
	var req *http.Request
	var buf *bytes.Buffer
	inputRequest.Method = strings.ToUpper(inputRequest.Method)
	if inputRequest.Body != nil {
		var b []byte
		b, err = json.Marshal(&inputRequest.Body)
		if err != nil {
			s.log(err)
			return
		}
		buf = bytes.NewBuffer(b)
		if buf != nil {
			req, err = http.NewRequest(inputRequest.Method, parsedUrl.String(), buf)
		} else {
			req, err = http.NewRequest(inputRequest.Method, parsedUrl.String(), nil)
		}
		if err != nil {
			s.log(err)
			return
		}
		// TODO only support application/json
		header.Add("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest(inputRequest.Method, parsedUrl.String(), nil)
		if err != nil {
			s.log(err)
			return
		}

	}
	// Set all headers in the request
	req.Header = header

	// Basic Auth for the request
	var userInfo *url.Userinfo
	if parsedUrl.User != nil {
		userInfo = parsedUrl.User
	}
	if s.Userinfo != nil {
		userInfo = s.Userinfo
	}
	if userInfo != nil {
		pwd, _ := userInfo.Password()
		req.SetBasicAuth(userInfo.Username(), pwd)
	}



	// Debug log request
	s.log("--------------------------------------------------------------------------------")
	s.log("REQUEST")
	s.log("--------------------------------------------------------------------------------")
	s.log(pretty(req))
	s.log("Body:")

	s.log(pretty(inputRequest.Body))


	// TODO we need to capture elapsed time
	response.timestamp = time.Now()

	//
	// Execute the HTTP request
	//
	var client *http.Client
	if s.Client != nil {
		client = s.Client
	} else {
		client = &http.Client{}
	}
	resp, err := client.Do(req)
	if err != nil {
		s.log(err)
		return
	}
	// Don't forget to close the stream on return
	defer resp.Body.Close()

	// Extract and set JSESSIONID  into session
	s.setJsessionId(resp.Header)

	//
	// Set Response
	//
	response.status = resp.StatusCode
	response.httpResponse = resp
	// Read body if status if good
	if resp.StatusCode < 300 {
		response.body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			s.log(err)
			return
		}
	}
	// TODO Only work for json
	// Deserialize json if Result was provided
	if string(response.body) != "" && response.Result != nil {
		err = response.Result.FromJson(response.body)
	}


	// Debug log response
	s.log("--------------------------------------------------------------------------------")
	s.log("RESPONSE")
	s.log("--------------------------------------------------------------------------------")
	s.log("Status: ", response.status)
	s.log("Result: ", response.Result)

	return
}

