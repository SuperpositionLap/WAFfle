# WAFfle
Web Application Firewall for lightweight environments
(Web Application Firewall for fun, lol, eh...)  

## 목표
⚡ 성능: 초당 50,000+ 요청 처리  
🛡️ 보안: OWASP Top 10 완전 대응  
📊 분석: 실시간 위협 탐지 및 통계  
🎛️ 관리: 직관적인 웹 UI로 모든 설정 제어

## 아키텍처

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Client        │──>─│  WAF Proxy      │──>─│  Backend        │
│   Requests      │    │  (Gin Service)  │    │  Services       │
└─────────────────┘    └─────────────────┘    └─────────────────┘
                                │
                                ▼
                       ┌─────────────────┐
                       │  Log Buffer     │
                       │  (Channel)      │
                       └─────────────────┘
                                │
                                ▼
                    ┌─────────────────────────────┐
                    │     PostgreSQL              │
                    │  + TimescaleDB Extension    │
                    │                             │
                    │  ┌─────────────────────┐    │
                    │  │   WAF Logs          │    │
                    │  │   (Hypertable)      │    │
                    │  └─────────────────────┘    │
                    │  ┌─────────────────────┐    │
                    │  │   Rules & Config    │    │
                    │  │   (Regular Tables)  │    │
                    │  └─────────────────────┘    │
                    │  ┌─────────────────────┐    │
                    │  │   Continuous        │    │
                    │  │   Aggregates        │    │
                    │  └─────────────────────┘    │
                    └─────────────────────────────┘
                                │
                                ▼
                       ┌─────────────────┐
                       │  Admin UI       │
                       │  (Sveltekit)    │
                       └─────────────────┘
```

# 할거
### Challenge 1.1: WAF 엔진 핵심 구현
목표: 기본적인 HTTP 요청 필터링 시스템 구축  
세부 과제:  
✅ Gin 미들웨어 체인 설계  
✅ SQL Injection 탐지 엔진 (정규식 + 파서 기반)  
✅ XSS 필터링 시스템  
✅ Rate Limiting (Token Bucket 알고리즘)  
✅ IP 화이트리스트/블랙리스트  
✅ 비동기 로깅 시스템 (고루틴 + 채널)  
핵심 도전 요소:  
  
성능 최적화: 요청당 1ms 미만 추가 지연  
메모리 효율성: 고루틴 풀 및 오브젝트 풀링  
정확도: False Positive < 1%  

### Challenge 1.2: TimescaleDB 스키마 설계  
목표: 고성능 시계열 로깅 시스템 구축  
세부 과제:  
✅ Hypertable 스키마 설계 (waf_logs, rules, config)  
✅ 자동 파티셔닝 설정 (1시간 단위 청크)  
✅ 압축 정책 구현 (24시간 후 90% 압축)
✅ Continuous Aggregates (분/시간/일 단위 통계)  
✅ 데이터 보존 정책 (90일 자동 삭제)  
✅ 성능 인덱스 최적화  
핵심 도전 요소:  
  
인서트 성능: 초당 100만 행 인서트 달성  
쿼리 최적화: 시계열 쿼리 1000배 성능 향상  
저장 효율성: 90% 압축률 달성  

### Challenge 1.3: Go 서비스 아키텍처  
목표: 확장 가능한 마이크로서비스 기반 설계  
세부 과제:  
✅ Clean Architecture 적용 (Domain/UseCase/Interface)  
✅ 의존성 주입 컨테이너 (Wire 또는 Fx)  
✅ gRPC 서비스 간 통신  
✅ 헬스 체크 및 메트릭 수집  
✅ 설정 관리 (Viper + 환경변수)  
✅ 구조화된 로깅 (Zap + 트레이싱)  

### Challenge 2.1: 실시간 대시보드 구현  
목표: WebSocket 기반 실시간 모니터링 시스템  
세부 과제:  
✅ SvelteKit + WebSocket 통합  
✅ 실시간 트래픽 차트 (Chart.js/D3)  
✅ 위협 탐지 알림 시스템  
✅ 실시간 로그 스트리밍  
✅ 성능 메트릭 대시보드  
  
### Challenge 2.2: 관리 인터페이스 개발  
목표: 직관적인 WAF 설정 및 관리 시스템  
세부 과제:  
 WAF 규칙 CRUD 인터페이스  
✅ 드래그 앤 드롭 규칙 우선순위 조정  
✅ 실시간 규칙 테스트 기능  
✅ IP 리스트 관리 (CSV 업로드 지원)  
✅ 사용자 관리 및 권한 시스템  
✅ 설정 백업/복원 기능  
  
### Challenge 2.3: 데이터 시각화  
목표: 복잡한 보안 데이터의 직관적 표현  
세부 과제:  
✅ 공격 패턴 히트맵  
✅ 지리적 위협 분포 지도  
✅ 시계열 트렌드 분석  
✅ 사용자 정의 대시보드  
✅ 보고서 생성 및 PDF 내보내기  
  
### Challenge 3.2: 고급 공격 대응  
목표: 최신 공격 기법에 대한 대응 능력  
세부 과제:  
✅ GraphQL 쿼리 깊이 제한  
✅ 파일 업로드 보안 검사  
✅ API Rate Limiting (사용자별/IP별)  
✅ DDoS 패턴 탐지  
✅ Bot 탐지  