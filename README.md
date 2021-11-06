# YAAS - yet another authorization server

## Intro

YAAS is yet another authorization server.
The main goal is to provide a fully OAuth 2.0 + OpenID Connect server being also an Identity Provider, i.e. provide storage of users, groups, tenants etc.
I started this project to learn OAuth and OpenID Connect.

## Design decisions

Some self-imposed decisions are

- code should be fully covered by unit tests
- every function should be documented
- should have an administration UI
- use only well established libraries, e.g.
  - gorm for persistence layer
  - net/http for HTTP
  - gorilla for some technical help (CSRF, Cookies)
- OAuth / OpenID Connect related features should be implemented only with Go standard library
