# error-tree

에러 트리는 에러를 상속 트리처럼 다루게 도와주는 라이브러리입니다.

## 설치

`go get github.com/snowmerak/error-tree`

## 사용법

### 엔드 에러 선언

```go
Timeout := etree.New("timeout", nil)
Badrequest := etree.New("bad request", nil)
Notfound := etree.New("not found", nil)
Facebook := etree.New("facebook", nil)
Google := etree.New("google", nil)
```

각각 `Timeout`, `Badrequest`, `Notfound`, `Facebook`, `Google` 에러를 생성합니다.  
모두 `error` 인터페이스를 만족하기에 `Error()` 메서드를 실행할 수 있습니다.

### 노드 에러 선언

```go
FacebookTimeout := etree.New("facebook timeout", Timeout, Facebook)
FacebookBadrequest := etree.New("facebook bad request", Badrequest, Facebook)
FacebookNotfound := etree.New("facebook not found", Notfound, Facebook)
GoogleTimeout := etree.New("google timeout", Timeout, Google)
GoogleBadrequest := etree.New("google bad request", Badrequest, Google)
GoogleNotfound := etree.New("google not found", Notfound, Google)
```

각각 페이스북과 구글에서 타임아웃, 배드리퀘스트, 낫파운드가 발생했을 경우 제공할 에러입니다.  
부모 에러로 각기 다른 `*etree.Node`를 받아서 기록합니다.

### 에러 확인

```go
fmt.Println(
	etree.Cover(FacebookTimeout, Timeout), " ",
	etree.Cover(FacebookBadrequest, Facebook), " ",
	etree.Cover(FacebookNotfound, Facebook), " ",
	etree.Cover(GoogleTimeout, Timeout), " ",
	etree.Cover(GoogleBadrequest, Google), " ",
	etree.Cover(GoogleNotfound, Notfound), " ",
	etree.Cover(GoogleNotfound, Google), " ",
	etree.Cover(GoogleBadrequest, Facebook), " ",
	etree.Cover(FacebookTimeout, Google), " ",
)
```

`Cover` 함수를 통해 노드 에러가 특정 에러를 상속 받고 있는지 확인합니다.

```bash
true   true   true   true   true   true   true   false   false  
```

결과는 위와같이 일치하는 대로 7개의 `true`, 2개의 `false`를 출력합니다.
