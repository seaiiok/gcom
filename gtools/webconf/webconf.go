package webconf

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

type tpl struct {
	svrClose  chan string
	tplConf   map[string]string
	onGetConf func(map[string]string)
}

func New() *tpl {
	ts := new(tpl)
	ts.svrClose = make(chan string, 0)
	go ts.configServerStart()
	return ts
}

func (t *tpl) configServerStart() error {

	mux := http.NewServeMux()
	mux.HandleFunc("/get", t.webGetConfig)
	mux.HandleFunc("/set", t.webSetConfig)

	svr := &http.Server{
		Addr:    "localhost:497",
		Handler: mux,
	}
	go func() {
		for {
			select {
			case <-t.svrClose:
				err := svr.Close()
				if err != nil {
					fmt.Println("Close:", err)
				}
				err = svr.Shutdown(nil)
				if err != nil {
					fmt.Println("Shutdown:", err)
				}

			case <-time.After(20 * time.Second):

				err := svr.Close()
				if err != nil {
					fmt.Println("Close:", err)
				}
				err = svr.Shutdown(nil)
				if err != nil {
					fmt.Println("Shutdown:", err)
				}

			}
		}

	}()
	err := svr.ListenAndServe()
	return err

}

func (t *tpl) webGetConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("isAppConfig", "yes")
	tp, err := template.ParseFiles("./view/index.html")
	if err != nil {
		w.Write([]byte("web config error:" + err.Error()))
		return
	}
	t.tplConf = make(map[string]string)

	//获取现在的配置
	for i := 0; i < 10; i++ {
		t.tplConf["键"+strconv.Itoa(i)] = "值" + strconv.Itoa(i)
	}

	t.tplConf["Tip"] = "请谨慎修改配置,然后在此输入确认语句!"

	err = tp.Execute(w, t.tplConf)
	if err != nil {
		w.Write([]byte("web config error:" + err.Error()))
		return
	}
}

func (t *tpl) webSetConfig(w http.ResponseWriter, r *http.Request) {
	t.tplConf = make(map[string]string)

	err := r.ParseForm()
	if err != nil {
		w.Write([]byte("配置保存失败:" + err.Error()))
		t.svrClose <- "set fail"
		return
	}
	if r.Method == "POST" {
		for k, v := range r.PostForm {
			if len(v) == 1 {
				t.tplConf[k] = v[0]
			} else {
				w.Write([]byte("配置保存失败:配置项可能存在异常!"))
				t.svrClose <- "set fail"
				return
			}
		}
	}

	w.Write([]byte("配置保存成功!"))
	if _, v := t.tplConf["Tip"]; v == true {
		delete(t.tplConf, "Tip")
	}
	go t.onGetConf(t.tplConf)
	t.svrClose <- "set succ"

}

func (t *tpl) GetConfig(cf func(map[string]string)) {
	t.onGetConf = cf
}
