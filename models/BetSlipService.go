package models


type BetSlipService struct {
	S []struct {
		IP  bool    `json:"ip"`
		Ap  bool    `json:"ap"`
		Bn  string  `json:"bn"`
		Sn  string  `json:"sn"`
		Hn  string  `json:"hn"`
		An  string  `json:"an"`
		As  int     `json:"as"`
		Hs  int     `json:"hs"`
		Eid int     `json:"eid"`
		En  string  `json:"en"`
		Hd  string  `json:"hd"`
		Cpn string  `json:"cpn"`
		Sid int64   `json:"sid"`
		St  float64 `json:"st"`
		Bs  struct {
			Bmin float64 `json:"bmin"`
			Bmax float64 `json:"bmax"`
			Dm   string  `json:"dm"`
			Bpay float64 `json:"bpay"`
		} `json:"bs"`
		Bo    int           `json:"bo"`
		O     float64       `json:"o"`
		Do    string        `json:"do"`
		Eo    float64       `json:"eo"`
		Rsl   int           `json:"rsl"`
		Wid   int           `json:"wid"`
		Tl    string        `json:"tl"`
		Serr  string        `json:"serr"`
		Cv    float64       `json:"cv"`
		Wo    int           `json:"wo"`
		Ishp  bool          `json:"ishp"`
		Isnu  bool          `json:"isnu"`
		Isds  bool          `json:"isds"`
		Ors   string        `json:"ors"`
		Edt   string        `json:"edt"`
		Idan  bool          `json:"idan"`
		Peid  int           `json:"peid"`
		Isew  bool          `json:"isew"`
		Ispl  bool          `json:"ispl"`
		IsEpm bool          `json:"isEpm"`
		Dt    int           `json:"dt"`
		Iscc  bool          `json:"iscc"`
		Ril   []interface{} `json:"ril"`
		Dil   []interface{} `json:"dil"`
	} `json:"s"`
	C     []interface{} `json:"c"`
	Cinfo []struct {
		Cbn string `json:"cbn"`
		Wid int    `json:"wid"`
		Cbs struct {
			Bmin float64 `json:"bmin"`
			Bmax float64 `json:"bmax"`
			Bpay float64 `json:"bpay"`
		} `json:"cbs"`
		Cbsew struct {
			Bmin float64 `json:"bmin"`
			Bmax float64 `json:"bmax"`
			Bpay float64 `json:"bpay"`
		} `json:"cbsew"`
		Cba  float64 `json:"cba"`
		Isew bool    `json:"isew"`
		Ispl bool    `json:"ispl"`
	} `json:"cinfo"`
	Bo    int     `json:"bo"`
	Berr  int     `json:"berr"`
	Cno   int     `json:"cno"`
	Islog bool    `json:"islog"`
	Issa  bool    `json:"issa"`
	Isca  bool    `json:"isca"`
	Cc    string  `json:"cc"`
	Ot    int     `json:"ot"`
	L     string  `json:"l"`
	K     bool    `json:"k"`
	Ko    bool    `json:"ko"`
	Abo   bool    `json:"abo"`
	Sbc   bool    `json:"sbc"`
	Bal   float64 `json:"bal"`
	Usc   int     `json:"usc"`
}