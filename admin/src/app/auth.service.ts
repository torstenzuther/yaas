import { Injectable } from '@angular/core';
import { authCodeFlowConfig } from "./authConfig";
import { OAuthService } from "angular-oauth2-oidc";

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  constructor(private readonly oauthService: OAuthService) {
  }

  init(): void {
    this.oauthService.configure(authCodeFlowConfig);
    this.oauthService.loadDiscoveryDocumentAndTryLogin();
  }
}
