# Go의 기본

## 레이아웃

### Go 디렉터리
#### /cmd
메인 실행 어플리케이션

#### /internal
개인적인 애플리케이션과 라이브러리 코드.
다른 사람들이 어플리케이션이나 라이브러리에서 임포트하기 원하지 않는 코드들.
Go 의 compiler 자체에서 패턴이 강제됨.
root 뿐 아닌 모든 자식레벨의 디렉터리에 적용될 수 있음.

#### /pkg
외부 어플리케이션에서 사용되어도 괜찮은 코드 모음.

#### /vendor
어플리케이션 종속성 

### 서비스 app 디렉터리

#### /api
OpenAPI/Swagger spec, json schema 파일, 프로토콜 정의 파일


### 웹 app 디렉터리
#### /api
정적파일

### 공통 app 디렉터리
#### /configs
설정 파일 템플릿 , 기본 설정.

#### /init
시스템 int(systemd, upstart, sysv) 설정들

#### /scripts
빌드, 설치, 분석, 기타 작업을 위한 스크립트들.

#### /build
패키징과 CI  
클라우드 컨테이너 설정과 스크립트를 넣으세요


#### /deployments
Iaas, Paas 시스템과 컨테이너 오케스트레이션 배포설정.
