/*
HAPROXY DATAPLANE API CLIENT (V2)
AUTHOR: Omar Aouini | Allan Nava
mail: aouini.omar93@gmail.com
Date: 19/04/2022
Update: 21/10/2022
MIT License

Copyright (c) 2022 OmarAouini

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package haproxy

type HaproxyErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (h *HaproxyErrorResponse) Error() string {
	return h.Message
}

type HaproxyInfo struct {
	Haproxy struct {
		Pid         int    `json:"pid"`
		Processes   int    `json:"processes"`
		ReleaseDate string `json:"release_date"`
		Uptime      int    `json:"uptime"`
		Version     string `json:"version"`
	} `json:"haproxy"`
}

type HaproxySites struct {
	Version int `json:"_version"`
	Data    []struct {
		Farms []struct {
			Balance struct {
				Algorithm string `json:"algorithm"`
			} `json:"balance"`
			Mode    string `json:"mode"`
			Name    string `json:"name"`
			Servers []struct {
				Address        string      `json:"address"`
				Check          string      `json:"check"`
				Inter          int         `json:"inter"`
				Maxconn        int         `json:"maxconn"`
				Name           string      `json:"name"`
				Port           int         `json:"port"`
				ProxyV2Options interface{} `json:"proxy-v2-options"`
				Weight         int         `json:"weight"`
			} `json:"servers"`
			UseAs string `json:"use_as"`
		} `json:"farms"`
		Name    string `json:"name"`
		Service struct {
			Listeners []struct {
				Name    string `json:"name"`
				Address string `json:"address"`
				Port    int    `json:"port"`
			} `json:"listeners"`
			Mode string `json:"mode"`
		} `json:"service"`
	} `json:"data"`
}

type HaproxyStats []struct {
	Error      string `json:"error"`
	RuntimeAPI string `json:"runtimeAPI"`
	Stats      []struct {
		BackendName string `json:"backend_name"`
		Name        string `json:"name"`
		Stats       struct {
			Bin         int    `json:"bin"`
			Bout        int    `json:"bout"`
			CompByp     int    `json:"comp_byp"`
			CompIn      int    `json:"comp_in"`
			CompOut     int    `json:"comp_out"`
			CompRsp     int    `json:"comp_rsp"`
			ConnRate    int    `json:"conn_rate"`
			ConnRateMax int    `json:"conn_rate_max"`
			ConnTot     int    `json:"conn_tot"`
			Dcon        int    `json:"dcon"`
			Dreq        int    `json:"dreq"`
			Dresp       int    `json:"dresp"`
			Dses        int    `json:"dses"`
			Ereq        int    `json:"ereq"`
			Hrsp1Xx     int    `json:"hrsp_1xx"`
			Hrsp2Xx     int    `json:"hrsp_2xx"`
			Hrsp3Xx     int    `json:"hrsp_3xx"`
			Hrsp4Xx     int    `json:"hrsp_4xx"`
			Hrsp5Xx     int    `json:"hrsp_5xx"`
			HrspOther   int    `json:"hrsp_other"`
			Iid         int    `json:"iid"`
			Intercepted int    `json:"intercepted"`
			Mode        string `json:"mode"`
			Pid         int    `json:"pid"`
			Rate        int    `json:"rate"`
			RateLim     int    `json:"rate_lim"`
			RateMax     int    `json:"rate_max"`
			ReqRate     int    `json:"req_rate"`
			ReqRateMax  int    `json:"req_rate_max"`
			ReqTotal    int    `json:"req_total"`
			Scur        int    `json:"scur"`
			Slim        int    `json:"slim"`
			Smax        int    `json:"smax"`
			Status      string `json:"status"`
			Stot        int    `json:"stot"`
		} `json:"stats"`
		Type string `json:"type"`
	} `json:"stats"`
}
type HaproxyReloads []struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type HaproxyTransactions []struct {
	Version int    `json:"_version"`
	ID      string `json:"id"`
	Status  string `json:"status"`
}

type HaproxyTransaction struct {
	Version int    `json:"_version"`
	ID      string `json:"id"`
	Status  string `json:"status"`
}

type HaproxyConfigurationGlobal struct {
	Version int `json:"_version"`
	Data    struct {
		CPUMaps []struct {
			CPUSet  string `json:"cpu_set"`
			Process string `json:"process"`
		} `json:"cpu_maps"`
		RuntimeApis []struct {
			Address string `json:"address"`
			Level   string `json:"level"`
			Mode    string `json:"mode"`
		} `json:"runtime_apis"`
		Daemon                string `json:"daemon"`
		MasterWorker          bool   `json:"master-worker"`
		Maxconn               int    `json:"maxconn"`
		Nbproc                int    `json:"nbproc"`
		SslDefaultBindCiphers string `json:"ssl_default_bind_ciphers"`
		SslDefaultBindOptions string `json:"ssl_default_bind_options"`
		StatsTimeout          int    `json:"stats_timeout"`
		TuneSslDefaultDhParam int    `json:"tune_ssl_default_dh_param"`
	} `json:"data"`
}

type HaproxyConfigurationDefaults struct {
	Version int `json:"_version"`
	Data    struct {
		ErrorFiles []struct {
			Code int    `json:"code"`
			File string `json:"file"`
		} `json:"error_files"`
		Balance struct {
			Algorithm string        `json:"algorithm,omitempty"`
			Arguments []interface{} `json:"arguments,omitempty"`
		} `json:"balance"`
		ClientTimeout  int `json:"client_timeout"`
		ConnectTimeout int `json:"connect_timeout"`
		DefaultServer  struct {
		} `json:"default_server"`
		Dontlognull          string `json:"dontlognull"`
		HTTPKeepAliveTimeout int    `json:"http_keep_alive_timeout"`
		Httplog              bool   `json:"httplog"`
		Mode                 string `json:"mode"`
		QueueTimeout         int    `json:"queue_timeout"`
		Redispatch           struct {
			Enabled string `json:"enabled"`
		} `json:"redispatch"`
		ServerTimeout int `json:"server_timeout"`
	} `json:"data"`
}

type HaproxyBackends struct {
	Version int `json:"_version"`
	Data    []struct {
		Balance struct {
			Algorithm string        `json:"algorithm,omitempty"`
			Arguments []interface{} `json:"arguments,omitempty"`
		} `json:"balance,omitempty"`
		Mode       string `json:"mode"`
		Name       string `json:"name"`
		Forwardfor struct {
			Enabled string `json:"enabled"`
		} `json:"forwardfor,omitempty"`
		Httpchk struct {
			Method string `json:"method"`
			URI    string `json:"uri"`
		} `json:"httpchk,omitempty"`
		HTTPConnectionMode string `json:"http_connection_mode,omitempty"`
		StickTable         struct {
			Expire int    `json:"expire"`
			Size   int    `json:"size"`
			Store  string `json:"store"`
			Type   string `json:"type"`
		} `json:"stick_table,omitempty"`
		AdvCheck string `json:"adv_check,omitempty"`
		HashType struct {
			Method string `json:"method"`
		} `json:"hash_type,omitempty"`
		HTTPKeepAliveTimeout int `json:"http_keep_alive_timeout,omitempty"`
		HTTPRequestTimeout   int `json:"http_request_timeout,omitempty"`
		ServerTimeout        int `json:"server_timeout,omitempty"`
	} `json:"data"`
}

type HaproxyFrontends struct {
	Version int `json:"_version"`
	Data    []struct {
		DefaultBackend string `json:"default_backend,omitempty"`
		Mode           string `json:"mode,omitempty"`
		Name           string `json:"name"`
		Tcplog         bool   `json:"tcplog,omitempty"`
		Address 	   string `json:"address,omitempty"`
		Port 	       string `json:"port,omitempty"`
		Check.         string `json:"check,omitempty"`
		Forwardfor     struct {
			Enabled string `json:"enabled"`
		} `json:"forwardfor,omitempty"`
	} `json:"data"`
}

type HaproxyBackendSwitchingRules struct {
	Version int `json:"_version"`
	Data    []struct {
		Cond     string `json:"cond"`
		CondTest string `json:"cond_test"`
		Index    int    `json:"index"`
		Name     string `json:"name"`
	} `json:"data"`
}

type HaproxyServers struct {
	Version int `json:"_version"`
	Data    []struct {
		Address string `json:"address"`
		Check   string `json:"check"`
		Name    string `json:"name"`
		Port    int    `json:"port"`
		Weight  int    `json:"weight"`
	} `json:"data"`
}

type HaproxyAcls struct {
	Version int `json:"_version"`
	Data    []struct {
		AclName   string `json:"acl_name"`
		Criterion string `json:"criterion"`
		Index     int    `json:"index"`
		Value     string `json:"value"`
	} `json:"data"`
}

type HaproxyServerSwitchingRules struct {
	Version int `json:"_version"`
	Data    []struct {
		Cond         string `json:"cond"`
		CondTest     string `json:"cond_test"`
		Index        int    `json:"index"`
		TargetServer string `json:"target_server"`
		ReturnHdrs   string `json:"return_hdrs"`
		PathFmt      string `json:"path_fmt"`
		Type         string `json:"type"`
	} `json:"data"`
}

type HaproxyHttpRequestRules struct {
	Version int `json:"_version"`
	Data    []struct {
		Cond       string  `json:"cond"`
		CondTest   string  `json:"cond_test"`
		HdrFormat  string  `json:"hdr_format"`
		HdrName    string  `json:"hdr_name"`
		Index      int     `json:"index"`
		Type       string  `json:"type"`
		ReturnHdrs *string `json:"return_hdrs"`
		PathFmt    *string `json:"path_fmt"`
	} `json:"data"`
}

type HaproxyAddBackend struct {
	AdvCheck string `json:"adv_check"`
	Balance  struct {
		Algorithm string `json:"algorithm"`
	} `json:"balance"`
	Forwardfor struct {
		Enabled string `json:"enabled"`
	} `json:"forwardfor"`
	HttpchkParams struct {
		Method  string `json:"method"`
		URI     string `json:"uri"`
		Version string `json:"version"`
	} `json:"httpchk_params"`
	Mode string `json:"mode"`
	Name string `json:"name"`
}

type HaproxyAddFrontend struct {
	DefaultBackend     string `json:"default_backend"`
	HTTPConnectionMode string `json:"http_connection_mode"`
	Maxconn            int    `json:"maxconn"`
	Mode               string `json:"mode"`
	Name               string `json:"name"`
}

type HaproxyAddAcl struct {
	AclName   string `json:"acl_name"`
	Criterion string `json:"criterion"`
	Index     int    `json:"index"`
	Value     string `json:"value"`
}

type HaproxyAddServer struct {
	Address string `json:"address"`
	Check   string `json:"check"`
	Name    string `json:"name"`
	Port    int    `json:"port"`
	Weight  int    `json:"weight"`
}

type HaproxyAddHttpRequestRule struct {
	Cond      string `json:"cond"`
	CondTest  string `json:"cond_test"`
	HdrFormat string `json:"hdr_format"`
	HdrName   string `json:"hdr_name"`
	Index     int    `json:"index"`
	Type      string `json:"type"`
}

type HaproxyAddBackendSwitchingRule struct {
	Cond     string `json:"cond"`
	CondTest string `json:"cond_test"`
	Index    int    `json:"index"`
	Name     string `json:"name"`
}

type HaproxyCommitTransaction struct {
	Version int    `json:"_version"`
	ID      string `json:"id"`
	Status  string `json:"status"`
}

type DuplicateCount struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}
type HaproxyDuplicateDefinitionsResult struct {
	Acls      []DuplicateCount `json:"acls"`
	Backends  []DuplicateCount `json:"backends"`
	Frontends []DuplicateCount `json:"frontends"`
	Servers   []DuplicateCount `json:"servers"`
}
