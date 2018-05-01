package servicehandlers

import (
  "net/http"
  "simplesurveygo/dao"
)
type ServeyHandler struct {
}
func (p ServeyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := methodRouter(p, w, r)
	response.(SrvcRes).RenderResponse(w)
}

func (p ServeyHandler) Get(r *http.Request) SrvcRes {
	dao.GetActiveServey()
}

func (p ServeyHandler) Put(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}

/* func (p ServeyHandler) Post(r *http.Request) SrvcRes {
		sur := dao.Survey {
			Title: "tv survey",
	        Description: "all tvs in city",
			Status: "avtive",
	        Questions: dao.Ques{
			    	Question: "how many tvs in city",
				    Options: ["color","lcd","led"], 
			},
		}
	dao.AddServey(sur) 
} */