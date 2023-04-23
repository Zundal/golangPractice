# golang 연습하기

```shell
# gin 설치하기
go get -u github.com/gin-gonic/gin
```

## 참고자료

1. [gin 확인하기](https://gin-gonic.com/ko-kr/docs/quickstart/)
2. [golang 시작하기](https://wooiljeong.github.io/go/go-tutorial-01/)

## go lang에 대한 설명

1. go.mod 파일에 의해 정의된 모듈이 패키지를 제공하는 모듈을 추적
   1. go.mod 파일은 Go 언어에서 사용하는 모듈 시스템의 핵심 파일 중 하나입니다. 
   2. 이 파일은 프로젝트의 의존성 관리를 위한 패키지 매니저인 Go Modules가 사용합니다. 
   2. go.mod 파일은 프로젝트의 루트 디렉토리에 위치하며, 프로젝트에서 사용하는 모든 패키지와 해당 패키지의 버전 정보를 명시합니다. 
   3. 이 파일을 사용하여 필요한 패키지를 다운로드하고 프로젝트에서 사용할 수 있습니다. 
   3. Go Modules는 의존성 관리를 자동화하고 프로젝트를 더 쉽게 빌드하고 배포할 수 있도록 지원합니다. 
   4. 이를 통해 Go 언어로 개발하는 프로젝트에서 패키지 버전 관리와 의존성 관리를 보다 간편하게 처리할 수 있습니다.

2. 동시성 지원: 고루틴(Goroutine)이라는 경량 스레드를 통해 동시성을 지원하며, 멀티코어 CPU를 효율적으로 활용할 수 있습니다.

## 추가조사해야할 자료

1. go 모듈 사이클

## go 언어 학습 사이클

1. 변수와 상수
2. 조건문과 반복문
3. 함수
4. 배열, 슬라이스, 맵
5. 포인터와 참조
6. 구조체와 인터페이스
7. 고루틴과 채널
8. defer문
9. 패키지와 모듈
10. 에러 처리

## 객체와 구조체의 차이

- 객체(Object)와 구조체(Struct)는 프로그래밍에서 다른 의미를 갖는 용어입니다.
- 객체는 속성(attribute)과 메서드(method)로 구성된 개체(object)를 의미합니다. 
- 객체 지향 프로그래밍에서는 객체가 중요한 개념으로, 객체의 상호작용을 통해 프로그램을 구성합니다.
- 구조체는 여러 필드(field)를 묶어서 하나의 타입(type)으로 정의하는데 사용되는 구성 요소입니다. 
- Go 언어에서는 구조체를 사용하여 사용자 정의 타입을 정의합니다. 
- 구조체는 필드의 집합으로 이루어져 있으며, 각 필드는 이름과 타입을 가지고 있습니다. 
- 구조체를 사용하면 논리적으로 연관된 데이터를 하나의 객체로 그룹화할 수 있으며, 이를 통해 코드의 가독성과 유지보수성을 높일 수 있습니다.
- 즉, 객체는 개념적인 개체를 표현하는 데 사용되는 개념으로, 구조체는 데이터를 구성하는 필드의 집합을 표현하는 데 사용되는 개념입니다. 
- 객체는 구조체와 메서드를 묶어서 사용하는 개념으로, 구조체는 필드의 집합으로 사용자 정의 타입을 정의하는 개념입니다.