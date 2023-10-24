package buuurst_dev

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Collector struct {
	config *BuuurstDevConfig
}

type collectorLogBody struct {
	ProjectID     int                 `json:"project_id"`
	ServiceKey    string              `json:"service_key"`
	RequestID     string              `json:"request_id"`
	RequestedAt   int64               `json:"requested_at"`
	RequestMethod string              `json:"method"`
	Path          string              `json:"path"`
	Query         map[string][]string `json:"query"`
	Cookie        map[string]string   `json:"cookie"`
	Header        map[string][]string `json:"header"`
	Body          string              `json:"body"`
	Status        int                 `json:"status"`
}

func NewCollector(config *BuuurstDevConfig) *Collector {
	return &Collector{config: config}
}

func (c *Collector) Collect(w *ResponseWriterWithStatus, r *http.Request, body []byte) {
	if !c.config.Enabled {
		return
	}

	if c.config.IsIgnoredPath(r.URL.Path) {
		return
	}

	cb := &collectorLogBody{}
	cb.ProjectID = c.config.ProjectID
	cb.ServiceKey = c.config.ServiceKey
	cb.RequestID = c.getRequestId(r.Header.Get("X-Request-Id"))
	cb.RequestedAt = time.Now().Unix()
	cb.RequestMethod = r.Method
	cb.Path = r.URL.Path
	cb.Query = r.URL.Query()

	cookies := make(map[string]string)
	for _, cookie := range r.Cookies() {
		cookies[cookie.Name] = cookie.Value
	}
	cb.Cookie = cookies

	headers := make(map[string][]string)
	for k, v := range r.Header {
		if c.config.IsCustomHeader(k) {
			headers[k] = v
		}
	}
	cb.Header = headers
	cb.Body = string(body)
	cb.Status = w.Status()

	c.send(cb)
}

func (c *Collector) getRequestId(reqID string) string {
	if reqID == "" {
		return uuid.New().String()
	}

	return reqID
}

func (c *Collector) send(body *collectorLogBody) error {
	data, err := json.Marshal(body)
	if err != nil {
		return err
	}

	_, err = http.Post(c.config.CollectorURL, "application/json", bytes.NewReader(data))
	if err != nil {
		return err
	}

	return nil
}
