package models

type CentralService struct {
	Uvd struct {
		Login  bool   `json:"login"`
		Ov     int    `json:"ov"`
		Sb     int    `json:"sb"`
		UrView string `json:"urView"`
		Efloat bool   `json:"efloat"`
		Showls bool   `json:"showls"`
		Nol    int    `json:"nol"`
		Iarf   bool   `json:"iarf"`
		Pds    int    `json:"pds"`
	} `json:"uvd"`

	Lpd struct {
		Sm struct {
			Fe struct {
				Progms []interface{} `json:"progms"`
			} `json:"fe"`
			Smd []struct {
				Sid  int    `json:"sid"`
				Sn   string `json:"sn"`
				Sen  string `json:"sen"`
				Mt   int    `json:"mt"`
				Son  int    `json:"son"`
				Tc   int    `json:"tc"`
				Tmrc int    `json:"tmrc"`
				Ipc  int    `json:"ipc"`
				Ec   int    `json:"ec"`
				Psc  int    `json:"psc"`
				Orc  int    `json:"orc"`
				Eec  int    `json:"eec"`
				Puc  []struct {
					Ces []interface{} `json:"ces"`
					Cid int           `json:"cid"`
					Cn  string        `json:"cn"`
					Cen string        `json:"cen"`
				} `json:"puc"`
				Pc int `json:"pc"`
			} `json:"smd"`
			V  int           `json:"v"`
			Fv int           `json:"fv"`
			Mc []interface{} `json:"mc"`
		} `json:"sm"`
		Ips struct {
			Ismd []struct {
				Sid  int    `json:"sid"`
				Sn   string `json:"sn"`
				Sen  string `json:"sen"`
				Mt   int    `json:"mt"`
				Son  int    `json:"son"`
				Tc   int    `json:"tc"`
				Tmrc int    `json:"tmrc"`
				Ipc  int    `json:"ipc"`
				Ec   int    `json:"ec"`
				Psc  int    `json:"psc"`
				Orc  int    `json:"orc"`
				Eec  int    `json:"eec"`
				Puc  []struct {
					Ces []struct {
						Eid  int    `json:"eid"`
						En   string `json:"en"`
						Gt   string `json:"gt"`
						Tid  string `json:"tid"`
						Ior  bool   `json:"ior"`
						Ht   string `json:"ht"`
						At   string `json:"at"`
						Ipt  string `json:"ipt"`
						Etts string `json:"etts"`
						Hs   struct {
							V string `json:"v"`
						} `json:"hs"`
						As struct {
							V string `json:"v"`
						} `json:"as"`
						Il   bool   `json:"il"`
						Mts  int    `json:"mts"`
						Isbg bool   `json:"isbg"`
						Lts  string `json:"lts"`
						Pt   int    `json:"pt"`
						Isb  bool   `json:"isb"`
					} `json:"ces"`
					Cid int    `json:"cid"`
					Cn  string `json:"cn"`
				} `json:"puc"`
				Pc int `json:"pc"`
			} `json:"ismd"`
			V int `json:"v"`
		} `json:"ips"`
		Ssm struct {
			Ssmd []struct {
				Sid  int    `json:"sid"`
				Sn   string `json:"sn"`
				Sen  string `json:"sen"`
				Mt   int    `json:"mt"`
				Son  int    `json:"son"`
				Tc   int    `json:"tc"`
				Tmrc int    `json:"tmrc"`
				Ipc  int    `json:"ipc"`
				Ec   int    `json:"ec"`
				Psc  int    `json:"psc"`
				Orc  int    `json:"orc"`
				Eec  int    `json:"eec"`
				Puc  []struct {
					Ces []struct {
						Eid  int      `json:"eid"`
						Gt   string   `json:"gt"`
						Tid  string   `json:"tid"`
						Ior  bool     `json:"ior"`
						Ht   string   `json:"ht"`
						At   string   `json:"at"`
						Est  string   `json:"est"`
						Hs   struct{} `json:"hs"`
						As   struct{} `json:"as"`
						Il   bool     `json:"il"`
						Mts  int      `json:"mts"`
						Isbg bool     `json:"isbg"`
						Pt   int      `json:"pt"`
						Isb  bool     `json:"isb"`
					} `json:"ces"`
					Cid int    `json:"cid"`
					Cn  string `json:"cn"`
					Cen string `json:"cen"`
				} `json:"puc"`
				Pc int `json:"pc"`
			} `json:"ssmd"`
			V int `json:"v"`
		} `json:"ssm"`
		Psm struct {
			Psmd []struct {
				Sid  int    `json:"sid"`
				Sn   string `json:"sn"`
				Sen  string `json:"sen"`
				Mt   int    `json:"mt"`
				Son  int    `json:"son"`
				Tc   int    `json:"tc"`
				Tmrc int    `json:"tmrc"`
				Ipc  int    `json:"ipc"`
				Ec   int    `json:"ec"`
				Psc  int    `json:"psc"`
				Orc  int    `json:"orc"`
				Eec  int    `json:"eec"`
				Puc  []struct {
					Ces []struct {
						Eid  int      `json:"eid"`
						En   string   `json:"en"`
						Gt   string   `json:"gt"`
						Tid  string   `json:"tid"`
						Ior  bool     `json:"ior"`
						Ht   string   `json:"ht"`
						At   string   `json:"at"`
						Esd  string   `json:"esd"`
						Est  string   `json:"est"`
						Hs   struct{} `json:"hs"`
						As   struct{} `json:"as"`
						Il   bool     `json:"il"`
						Mts  int      `json:"mts"`
						Isbg bool     `json:"isbg"`
						Pt   int      `json:"pt"`
						Isb  bool     `json:"isb"`
					} `json:"ces"`
					Cid int    `json:"cid"`
					Cn  string `json:"cn"`
				} `json:"puc"`
				Pc int `json:"pc"`
			} `json:"psmd"`
			V int `json:"v"`
		} `json:"psm"`
		Tipc int `json:"tipc"`
		Tssc int `json:"tssc"`
		Tpc  int `json:"tpc"`
	} `json:"lpd"`
	Lpc struct {
		Sm  int `json:"sm"`
		Smv int `json:"smv"`
		Imv int `json:"imv"`
		Ssv int `json:"ssv"`
	} `json:"lpc"`
	Mpc struct {
		Pv int `json:"pv"`
	} `json:"mpc"`
	Mod struct {
		M struct {
			Spid int `json:"spid"`
			Sid  int `json:"sid"`
			C    struct {
				Hasor  bool `json:"hasor"`
				Hasepm bool `json:"hasepm"`
				K      int  `json:"k"`
				E      []struct {
					Mts int    `json:"mts"`
					K   int    `json:"k"`
					Pk  int    `json:"pk"`
					Edt string `json:"edt"`
					Egn string `json:"egn"`
					IP  bool   `json:"ip"`
					NO  []struct {
						N         string     `json:"n"`
						Mn        string     `json:"mn"`
						F         string     `json:"f"`
						O         [][]string `json:"o"`
						Mo        int        `json:"mo"`
						Expnd     bool       `json:"expnd"`
						CanParlay bool       `json:"canParlay"`
					} `json:"n-o"`
					O struct {
						Ah       []string   `json:"ah"`
						Ou       []string   `json:"ou"`
						OneX2    []string   `json:"1x2"`
						Oe       []string   `json:"oe"`
						Ah1St    []string   `json:"ah1st"`
						Ou1St    []string   `json:"ou1st"`
						OneX21St []string   `json:"1x21st"`
						Oe1St    []string   `json:"oe1st"`
						Tg       []string   `json:"tg"`
						Cs       []string   `json:"cs"`
						Hf       []string   `json:"hf"`
						Tg1St    []string   `json:"tg1st"`
						Cs1St    []string   `json:"cs1st"`
						Tts1St   []string   `json:"tts1st"`
						Ttslast  []string   `json:"ttslast"`
						Sco1St   []string   `json:"sco1st"`
						Scolast  []string   `json:"scolast"`
						Bts      [][]string `json:"bts"`
						Eps      struct {
							N         string     `json:"n"`
							Mn        string     `json:"mn"`
							F         string     `json:"f"`
							O         [][]string `json:"o"`
							Mo        int        `json:"mo"`
							Expnd     bool       `json:"expnd"`
							CanParlay bool       `json:"canParlay"`
						} `json:"eps"`
					} `json:"o"`
					Hide bool   `json:"hide"`
					G    string `json:"g"`
					Cel  []struct {
						Mts int           `json:"mts"`
						K   int           `json:"k"`
						Pk  int           `json:"pk"`
						Edt string        `json:"edt"`
						Egn string        `json:"egn"`
						IP  bool          `json:"ip"`
						NO  []interface{} `json:"n-o"`
						O   struct {
							Ah []string `json:"ah"`
						} `json:"o"`
						Hide bool          `json:"hide"`
						G    string        `json:"g"`
						Cel  []interface{} `json:"cel"`
						Cei  struct {
							Ctid int    `json:"ctid"`
							N    string `json:"n"`
						} `json:"cei"`
						I    []string `json:"i"`
						Pvdr string   `json:"pvdr"`
						Ibs  bool     `json:"ibs"`
						Ibsc bool     `json:"ibsc"`
						Ihe  bool     `json:"ihe"`
						Heid int      `json:"heid"`
						Iscc bool     `json:"iscc"`
					} `json:"cel"`
					Cei struct {
						Ctid int    `json:"ctid"`
						N    string `json:"n"`
					} `json:"cei"`
					I    []string `json:"i"`
					Pvdr string   `json:"pvdr"`
					Ibs  bool     `json:"ibs"`
					Ibsc bool     `json:"ibsc"`
					Ihe  bool     `json:"ihe"`
					Heid int      `json:"heid"`
					Iscc bool     `json:"iscc"`
				} `json:"e"`
				N  string `json:"n"`
				Ec int    `json:"ec"`
			} `json:"c"`
			K     int    `json:"k"`
			Img   string `json:"img"`
			T     int    `json:"t"`
			Dv    int    `json:"dv"`
			Tt    int    `json:"tt"`
			Idm   bool   `json:"idm"`
			Msg   string `json:"msg"`
			Rd    string `json:"rd"`
			IP    bool   `json:"ip"`
			IsEpm bool   `json:"isEpm"`
		} `json:"m"`
		S []struct {
			Spid int    `json:"spid"`
			Sid  int    `json:"sid"`
			Sn   string `json:"sn,omitempty"`
			C    struct {
				Hasor  bool `json:"hasor"`
				Hasepm bool `json:"hasepm"`
				K      int  `json:"k"`
				E      []struct {
					Mts int    `json:"mts"`
					K   int    `json:"k"`
					Pk  int    `json:"pk"`
					Edt string `json:"edt"`
					Egn string `json:"egn"`
					IP  bool   `json:"ip"`
					NO  []struct {
						N         string     `json:"n"`
						Mn        string     `json:"mn"`
						O         [][]string `json:"o"`
						Mo        int        `json:"mo"`
						Expnd     bool       `json:"expnd"`
						CanParlay bool       `json:"canParlay"`
					} `json:"n-o"`
					O struct {
					} `json:"o"`
					Hide bool          `json:"hide"`
					G    string        `json:"g"`
					Cel  []interface{} `json:"cel"`
					Cei  struct {
						Ctid int    `json:"ctid"`
						N    string `json:"n"`
					} `json:"cei"`
					I    []string `json:"i"`
					Pvdr string   `json:"pvdr"`
					Ibs  bool     `json:"ibs"`
					Ibsc bool     `json:"ibsc"`
					Ihe  bool     `json:"ihe"`
					Heid int      `json:"heid"`
					Iscc bool     `json:"iscc"`
				} `json:"e"`
				N  string `json:"n"`
				Ec int    `json:"ec"`
			} `json:"c"`
			K     int    `json:"k"`
			Img   string `json:"img"`
			T     int    `json:"t"`
			Dv    int    `json:"dv"`
			Tt    int    `json:"tt"`
			Idm   bool   `json:"idm"`
			Msg   string `json:"msg"`
			Rd    string `json:"rd"`
			IP    bool   `json:"ip"`
			IsEpm bool   `json:"isEpm"`
		} `json:"s"`
		Hip []struct {
			Ec int    `json:"ec"`
			Rd string `json:"rd"`
			V  int    `json:"v"`
			C  []struct {
				Hasor  bool `json:"hasor"`
				Hasepm bool `json:"hasepm"`
				K      int  `json:"k"`
				E      []struct {
					Mts int    `json:"mts"`
					K   int    `json:"k"`
					Pk  int    `json:"pk"`
					Edt string `json:"edt"`
					Egn string `json:"egn"`
					IP  bool   `json:"ip"`
					NO  []struct {
						N         string     `json:"n"`
						Mn        string     `json:"mn"`
						F         string     `json:"f,omitempty"`
						O         [][]string `json:"o"`
						Mo        int        `json:"mo"`
						Expnd     bool       `json:"expnd"`
						CanParlay bool       `json:"canParlay"`
					} `json:"n-o"`
					O struct {
						Ah       []string   `json:"ah"`
						Ou       []string   `json:"ou"`
						OneX2    []string   `json:"1x2"`
						Oe       []string   `json:"oe"`
						Ah1St    []string   `json:"ah1st"`
						Ou1St    []string   `json:"ou1st"`
						OneX21St []string   `json:"1x21st"`
						Oe1St    []string   `json:"oe1st"`
						Tg       []string   `json:"tg"`
						Cs       []string   `json:"cs"`
						Hf       []string   `json:"hf"`
						Tg1St    []string   `json:"tg1st"`
						Cs1St    []string   `json:"cs1st"`
						Bts      [][]string `json:"bts"`
					} `json:"o"`
					Hide bool   `json:"hide"`
					G    string `json:"g"`
					Cel  []struct {
						Mts int    `json:"mts"`
						K   int    `json:"k"`
						Pk  int    `json:"pk"`
						Edt string `json:"edt"`
						Egn string `json:"egn"`
						IP  bool   `json:"ip"`
						NO  []struct {
							N         string     `json:"n"`
							Mn        string     `json:"mn"`
							O         [][]string `json:"o"`
							Mo        int        `json:"mo"`
							Expnd     bool       `json:"expnd"`
							CanParlay bool       `json:"canParlay"`
						} `json:"n-o"`
						O struct {
							Ou    []string `json:"ou"`
							Ou1St []string `json:"ou1st"`
						} `json:"o"`
						Hide bool          `json:"hide"`
						G    string        `json:"g"`
						Cel  []interface{} `json:"cel"`
						Cei  struct {
							Ctid int    `json:"ctid"`
							N    string `json:"n"`
						} `json:"cei"`
						I    []string `json:"i"`
						Pvdr string   `json:"pvdr"`
						Ibs  bool     `json:"ibs"`
						Ibsc bool     `json:"ibsc"`
						Ihe  bool     `json:"ihe"`
						Heid int      `json:"heid"`
						Iscc bool     `json:"iscc"`
					} `json:"cel"`
					Cei struct {
						Ctid int    `json:"ctid"`
						N    string `json:"n"`
					} `json:"cei"`
					I    []string `json:"i"`
					Pvdr string   `json:"pvdr"`
					Ibs  bool     `json:"ibs"`
					Ibsc bool     `json:"ibsc"`
					Ihe  bool     `json:"ihe"`
					Heid int      `json:"heid"`
					Iscc bool     `json:"iscc"`
				} `json:"e"`
				N  string `json:"n"`
				Ec int    `json:"ec"`
			} `json:"c"`
			K  int    `json:"k"`
			N  string `json:"n"`
			En string `json:"en"`
		} `json:"hip"`
		Ipec int `json:"ipec"`
		Hnph []struct {
			Ec  int    `json:"ec"`
			Tec int    `json:"tec,omitempty"`
			Rd  string `json:"rd"`
			V   int    `json:"v"`
			C   []struct {
				Hasor  bool `json:"hasor"`
				Hasepm bool `json:"hasepm"`
				K      int  `json:"k"`
				E      []struct {
					Mts int    `json:"mts"`
					K   int    `json:"k"`
					Pk  int    `json:"pk"`
					Edt string `json:"edt"`
					Egn string `json:"egn"`
					IP  bool   `json:"ip"`
					NO  []struct {
						N         string     `json:"n"`
						Mn        string     `json:"mn"`
						F         string     `json:"f"`
						O         [][]string `json:"o"`
						Mo        int        `json:"mo"`
						Expnd     bool       `json:"expnd"`
						CanParlay bool       `json:"canParlay"`
					} `json:"n-o"`
					O struct {
						Ah       []string   `json:"ah"`
						Ou       []string   `json:"ou"`
						OneX2    []string   `json:"1x2"`
						Oe       []string   `json:"oe"`
						Ah1St    []string   `json:"ah1st"`
						Ou1St    []string   `json:"ou1st"`
						OneX21St []string   `json:"1x21st"`
						Oe1St    []string   `json:"oe1st"`
						Tg       []string   `json:"tg"`
						Cs       []string   `json:"cs"`
						Hf       []string   `json:"hf"`
						Tg1St    []string   `json:"tg1st"`
						Cs1St    []string   `json:"cs1st"`
						Tts1St   []string   `json:"tts1st"`
						Ttslast  []string   `json:"ttslast"`
						Sco1St   []string   `json:"sco1st"`
						Scolast  []string   `json:"scolast"`
						Bts      [][]string `json:"bts"`
						Eps      struct {
							N         string     `json:"n"`
							Mn        string     `json:"mn"`
							F         string     `json:"f"`
							O         [][]string `json:"o"`
							Mo        int        `json:"mo"`
							Expnd     bool       `json:"expnd"`
							CanParlay bool       `json:"canParlay"`
						} `json:"eps"`
					} `json:"o"`
					Hide bool   `json:"hide"`
					G    string `json:"g"`
					Cel  []struct {
						Mts int           `json:"mts"`
						K   int           `json:"k"`
						Pk  int           `json:"pk"`
						Edt string        `json:"edt"`
						Egn string        `json:"egn"`
						IP  bool          `json:"ip"`
						NO  []interface{} `json:"n-o"`
						O   struct {
							Ah []string `json:"ah"`
						} `json:"o"`
						Hide bool          `json:"hide"`
						G    string        `json:"g"`
						Cel  []interface{} `json:"cel"`
						Cei  struct {
							Ctid int    `json:"ctid"`
							N    string `json:"n"`
						} `json:"cei"`
						I    []string `json:"i"`
						Pvdr string   `json:"pvdr"`
						Ibs  bool     `json:"ibs"`
						Ibsc bool     `json:"ibsc"`
						Ihe  bool     `json:"ihe"`
						Heid int      `json:"heid"`
						Iscc bool     `json:"iscc"`
					} `json:"cel"`
					Cei struct {
						Ctid int    `json:"ctid"`
						N    string `json:"n"`
					} `json:"cei"`
					I    []string `json:"i"`
					Pvdr string   `json:"pvdr"`
					Ibs  bool     `json:"ibs"`
					Ibsc bool     `json:"ibsc"`
					Ihe  bool     `json:"ihe"`
					Heid int      `json:"heid"`
					Iscc bool     `json:"iscc"`
				} `json:"e"`
				N  string `json:"n"`
				Ec int    `json:"ec"`
			} `json:"c"`
			K  int    `json:"k"`
			N  string `json:"n"`
			En string `json:"en"`
		} `json:"hnph"`
		Hnps []struct {
			Ec int    `json:"ec"`
			Rd string `json:"rd"`
			V  int    `json:"v"`
			C  []struct {
				Hasor  bool `json:"hasor"`
				Hasepm bool `json:"hasepm"`
				K      int  `json:"k"`
				E      []struct {
					Mts int    `json:"mts"`
					K   int    `json:"k"`
					Pk  int    `json:"pk"`
					Edt string `json:"edt"`
					Egn string `json:"egn"`
					IP  bool   `json:"ip"`
					NO  []struct {
						N         string     `json:"n"`
						Mn        string     `json:"mn"`
						F         string     `json:"f"`
						O         [][]string `json:"o"`
						Mo        int        `json:"mo"`
						Expnd     bool       `json:"expnd"`
						CanParlay bool       `json:"canParlay"`
					} `json:"n-o"`
					O struct {
						Ah       []string   `json:"ah"`
						Ou       []string   `json:"ou"`
						OneX2    []string   `json:"1x2"`
						Oe       []string   `json:"oe"`
						Ah1St    []string   `json:"ah1st"`
						Ou1St    []string   `json:"ou1st"`
						OneX21St []string   `json:"1x21st"`
						Oe1St    []string   `json:"oe1st"`
						Tg       []string   `json:"tg"`
						Cs       []string   `json:"cs"`
						Hf       []string   `json:"hf"`
						Tg1St    []string   `json:"tg1st"`
						Cs1St    []string   `json:"cs1st"`
						Tts1St   []string   `json:"tts1st"`
						Ttslast  []string   `json:"ttslast"`
						Bts      [][]string `json:"bts"`
					} `json:"o"`
					Hide bool   `json:"hide"`
					G    string `json:"g"`
					Cel  []struct {
						Mts int           `json:"mts"`
						K   int           `json:"k"`
						Pk  int           `json:"pk"`
						Edt string        `json:"edt"`
						Egn string        `json:"egn"`
						IP  bool          `json:"ip"`
						NO  []interface{} `json:"n-o"`
						O   struct {
							Ou    []string `json:"ou"`
							Oe    []string `json:"oe"`
							Ou1St []string `json:"ou1st"`
							Oe1St []string `json:"oe1st"`
						} `json:"o"`
						Hide bool          `json:"hide"`
						G    string        `json:"g"`
						Cel  []interface{} `json:"cel"`
						Cei  struct {
							Ctid int    `json:"ctid"`
							N    string `json:"n"`
						} `json:"cei"`
						I    []string `json:"i"`
						Pvdr string   `json:"pvdr"`
						Ibs  bool     `json:"ibs"`
						Ibsc bool     `json:"ibsc"`
						Ihe  bool     `json:"ihe"`
						Heid int      `json:"heid"`
						Iscc bool     `json:"iscc"`
					} `json:"cel"`
					Cei struct {
						Ctid int    `json:"ctid"`
						N    string `json:"n"`
					} `json:"cei"`
					I    []string `json:"i"`
					Pvdr string   `json:"pvdr"`
					Ibs  bool     `json:"ibs"`
					Ibsc bool     `json:"ibsc"`
					Ihe  bool     `json:"ihe"`
					Heid int      `json:"heid"`
					Iscc bool     `json:"iscc"`
				} `json:"e"`
				N  string `json:"n"`
				Ec int    `json:"ec"`
			} `json:"c"`
			K  int    `json:"k"`
			N  string `json:"n"`
			En string `json:"en"`
		} `json:"hnps"`
		T int `json:"t"`
		R int `json:"r"`
	} `json:"mod"`
	Selobj struct {
		Pid  int    `json:"pid"`
		Isp  bool   `json:"isp"`
		Spt  int    `json:"spt"`
		Sptn string `json:"sptn"`
		Evt  int    `json:"evt"`
		Cids string `json:"cids"`
		Gid  string `json:"gid"`
		Dp   int    `json:"dp"`
		FavT int    `json:"favT"`
		Btp  string `json:"btp"`
		Uibt string `json:"uibt"`
		Edt  string `json:"edt"`
		Isfd bool   `json:"isfd"`
		IP   bool   `json:"ip"`
		Ipo  bool   `json:"ipo"`
		Ifl  bool   `json:"ifl"`
	} `json:"selobj"`
}
