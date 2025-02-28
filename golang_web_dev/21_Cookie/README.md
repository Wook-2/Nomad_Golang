# [Go] Golang으로 쿠키

## 쿠키란 무엇인가?

**쿠키**는 `서버`가 `클라이언트`의 웹 브라우저에 저장하는 정보/파일 입니다. 최대 4kb까지 저장이 가능하며 쿠키의 최대 갯수는 브라우저마다 상이합니다. 상태를 유지를 포함해 여러 활용이 가능하며 클라이언트에게 파일을 저장하게 하므로 서버의 부하를 덜어주는 역할도 합니다.  
**쿠키**는 `key:value`쌍을 가지며 두 값 외에도 여러 값을 가집니다. 아래는 go언어에서의 cookie 타입입니다. 몇 개 알아보자면 `Path`는 쿠키를 전송할 요청 경로를, `Domain`은 쿠키를 전송할 도메인을 뜻하며 `Expires`, `MaxAge`는 쿠키의 만료에 대한 요소인데 뒤에서 다뤄보겠습니다. 

```go
type Cookie struct {
    Name  string
    Value string

    Path       string    // optional
    Domain     string    // optional
    Expires    time.Time // optional
    RawExpires string    // for reading cookies only

    // MaxAge=0 means no 'Max-Age' attribute specified.
    // MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
    // MaxAge>0 means Max-Age attribute present and given in seconds
    MaxAge   int
    Secure   bool
    HttpOnly bool
    SameSite SameSite // Go 1.11
    Raw      string
    Unparsed []string // Raw text of unparsed attribute-value pairs
}
```



## 쿠키 쓰기

고언어에서는 쿠키를 `net/http`패키지의 `SetCookie(w ResponseWriter, cookie *Cookie)`함수로 만들 수 있습니다.   
`ResponseWriter`인터페이스와 `Cookie`구조체를 인자로 전해주어 사용합니다.

```go
func createCookie(w http.ResponseWriter, req *http.Request) { 
    http.SetCookie(w, &http.Cookie{
        Name: "name of cookie",
        Value: "value of cookie",
        Path: "/",
    })
}
```

## 쿠키 읽기

이렇게 만든 쿠키를 `request`의 메서드 `request.Cookie(name string)`를 사용해 읽어줍니다.  
반환값으로 쿠키와 에러를 받으며 `name`을 가진 쿠키가 없을경우 `http.ErrNoCookie`에러를 반환해줍니다.

```go
func readCookie(w http.ResponseWriter, req *http.Request) {
    cookie, err := req.Cookie("name of cookie")
    if err == http.ErrNoCookie {
        fmt.Println("there is no cookie named 'name of cookie'")
        return
    }
    fmt.Fprintln(w, "Your cookie: ", cookie)
}
```

`request.Cookies()`를 사용해 여러개의 쿠키를 한번에 읽어올 수도 있습니다.

```go
func readCookie(w http.ResponseWriter, req *http.Request) {
    cookie, err := req.Cookies("name of cookie")
    if err == http.ErrNoCookie {
        fmt.Println("there is no cookie named 'name of cookie'")
        return
    }
    for _, c := range cookie {
        fmt.Fprinln(w,"cookie: ", c)
    }
    fmt.Fprintln(w, "Your cookie: ", cookie)
}
```



## 전체 코드

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", createCookie)
	http.HandleFunc("/read", readCookie)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func createCookie(w http.ResponseWriter, req *http.Request) { 
    http.SetCookie(w, &http.Cookie{
        Name: "name of cookie",
        Value: "value of cookie",
        Path: "/",
    })
}

func readCookie(w http.ResponseWriter, req *http.Request) {
    cookie, err := req.Cookie("name of cookie")
    if err == http.ErrNoCookie {
        fmt.Println("there is no cookie named 'name of cookie'")
        return
    }
    fmt.Fprintln(w, "Your cookie: ", cookie)
}
```

코드를 돌려 `localhost:8080`에 접속을 하면 인덱스 페이지에서 `name`이 `"name of cookie"`, `value`가 `"value of cookie"`인 쿠키를 만들게 됩니다. 만들어진 쿠키는 크롬 브라우저에서 `F12`-`Application`-`Storage`-`Cookies`에서 확인할 수 있습니다. 이렇게 쿠키를 생성해준 뒤 `localhost:8080/read`로 이동해주면 쿠키의 정보를 제대로 읽어올 수 있는것 을 확인할 수 있습니다. 

## 쿠키 삭제

이렇게 만든 쿠키를 삭제하는 방법에는 두가지가 있습니다. 첫 째는 쿠키의 `Expires`값을 설정해 주는 것, 둘 째는 쿠키의 `MaxAge`값을 설정해 주는 것입니다. 이 두가지의 차이점은 `expires`는 쿠키를 언제까지 지속시킬 지에 대한 시각을 받고, `MaxAge`는 그 시점부터 얼마나 지속시킬지에 대한 시간을 받습니다. 허나 `expires`는 **Deprecated**되었다고 하니 `MaxAge`로 쿠키의 만료시간을 정해주는 게 좋아보입니다.    

> golang에서 **쿠키**의 `Expires`, `MaxAge`값을 따로 설정해 주지 않으면 해당 쿠키는 **session cookie**가 되며 `session cookie`는 브라우저를 닫을 때 같이 삭제되는 쿠키입니다. 이와 대립하는 쿠키의 종류는  **persistent cookie**로 항상 유지되는 쿠키를 뜻합니다. 

쿠키의 `MaxAge`값을 `-1`로 설정해주면 쿠키가 삭제되고 자연수 `n`으로 설정해주면 n초동안 지속된 뒤 삭제됩니다.  
`Expires`값을 통해 쿠키의 수명을 정해주고 싶다면 `time`패키지를 사용해주면 되겠다.

```go
func expireCookie(w http.ResponseWriter, req *http.Request) {
    cookie, err := r.Cookie("cookie-name")
    
    if err := nil{
        // 에러처리
    }
    // 1. 쿠키 삭제
    // cookie.MaxAge = -1
    
    // 2. 쿠키를 5초간 지속
    // cookie.MaxAge = 5
    
    // 3. expires 설정, 현재시간으로 부터 1시간 뒤 == 쿠키의 만료시각
    // expiration := time.Now().Add(time.Hour)
    // cookie.Expires = expiration
    
    http.SetCookie(w, cookie)
    http.Redirect(w, req, "/", http.StatusSeeOther)
}
```

